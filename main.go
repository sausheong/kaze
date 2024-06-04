package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hako/durafmt"
	webview "github.com/webview/webview_go"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var exedir string
var doc string

const model = "llama3"

func init() {
	// getting the current executable directory, by default it should the
	// current executable binary, for testing this needs to be set via a
	// environment variable
	var ok bool
	exedir, ok = os.LookupEnv("EXE_DIR")
	if !ok {
		exepath, _ := os.Executable()
		exedir = filepath.Dir(exepath)
	}
	homedir, _ := os.UserHomeDir()
	_, err := os.Stat(filepath.Join(homedir, ".kaze"))
	if err != nil {
		log.Println(".kaze file doesn't exist, will create a empty file now:", err)
		createSettingsEnv()
	}

	err = godotenv.Load(filepath.Join(homedir, ".kaze"))
	if err != nil {
		log.Println("Error loading .kaze file")
	}
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	go webserver()
	w := webview.New(false)
	defer w.Destroy()

	w.SetTitle("kaze")
	w.SetSize(1024, 768, webview.HintNone)
	w.Bind("closeWindow", func() {
		w.Terminate()
	})
	w.Navigate("http://127.0.0.1:" + os.Getenv("PORT"))
	w.Run()

}

func webserver() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir(filepath.Join(exedir, "static")))))
	r.Get("/", index)
	r.Get("/history", history)
	r.Get("/settings", settings)
	r.Get("/document", document)
	r.Get("/document/clear", clear)
	r.Post("/ask", ask)
	r.Post("/document/load", load)
	r.Post("/models/pull", pull)

	err := http.ListenAndServe("127.0.0.1:"+os.Getenv("PORT"), r)
	if err != nil {
		log.Println("cannot start server:", err)
	}
}

// index
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(exedir, "static", "index.html"))
	if err != nil {
		log.Println("Cannot parse template:", err)
	}

	t.Execute(w, nil)
}

func document(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(exedir, "static", "doc.html"))
	if err != nil {
		log.Println("Cannot parse template:", err)
	}
	t.Execute(w, doc)
}

func clear(w http.ResponseWriter, r *http.Request) {
	doc = ""
	homedir, _ := os.UserHomeDir()
	err := os.Remove(filepath.Join(homedir, "kazedoc.gob"))
	if err != nil {
		fmt.Println("Error deleting the file:", err)
	}
}

func history(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(exedir, "static", "history.html"))
	if err != nil {
		log.Println("Cannot parse template:", err)
	}
	env := getEnvSettings()
	data := []byte{}
	if env["HISTORY_FILE"] != "" {
		data, err = os.ReadFile(env["HISTORY_FILE"])
		if err != nil {
			log.Println("Error reading file:", err)
			return
		}
	}
	t.Execute(w, string(data))
}

func settings(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(exedir, "static", "settings.html"))
	if err != nil {
		log.Println("Cannot parse template:", err)
	}
	env := getEnvSettings()
	if env["HISTORY_FILE"] == "" {
		env["HISTORY_FILE"] = getDefaultHistoryFile()
		putEnvSettings(env)
	}
	t.Execute(w, env)
}

func ask(w http.ResponseWriter, r *http.Request) {
	prompt := struct {
		Input string `json:"input"`
	}{}
	// decode JSON from client
	err := json.NewDecoder(r.Body).Decode(&prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("mime-type", "text/event-stream")
	response, err := predict(model, prompt.Input, doc, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeToHistory(prompt.Input, response)
}

func load(w http.ResponseWriter, r *http.Request) {
	log.Println("loading document")
	d := struct {
		File string `json:"file"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println(err)
		return
	}
	split := strings.Split(d.File, ".")
	fmt.Println(split)
	extension := split[len(split)-1]
	if extension != "pdf" {
		log.Println("not pdf:", extension)
		http.Error(w, "not pdf:"+extension, http.StatusInternalServerError)
	}

	content, err := convert(d.File)
	if err != nil {
		log.Println("unable to create a temporary file:", err)
		http.Error(w, "unable to create a temporary file:"+err.Error(), http.StatusInternalServerError)
		return
	}
	doc = clean(content)
	w.Header().Add("mime-type", "text/plain")
	w.Write([]byte(doc))
}

func pull(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("mime-type", "text/event-stream")
	err := pullModel(model, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func pullModel(name string, w http.ResponseWriter) error {
	type PullResponse struct {
		Status    string `json:"status"`
		Digest    string `json:"digest"`
		Total     int64  `json:"total"`
		Completed int64  `json:"completed"`
		Error     string `json:"error"`
	}

	name = strings.ToLower(name)
	reqJson := `{
	"name": "` + name + `"
}`
	r := bytes.NewReader([]byte(reqJson))
	httpResp, err := http.Post("http://localhost:11434/api/pull", "application/json", r)
	if err != nil {
		log.Println("err in calling ollama:", err)
		return err
	}
	decoder := json.NewDecoder(httpResp.Body)
	t0 := time.Now()
	f := w.(http.Flusher)
	for {
		resp := &PullResponse{}
		decoder.Decode(&resp)
		if resp.Status == "success" {
			elapsed := durafmt.Parse(time.Since(t0)).LimitFirstN(2)
			w.Write([]byte("success (" + elapsed.String() + ")<br/>"))
			f.Flush()
			break
		} else {
			if resp.Error != "" {
				w.Write([]byte(resp.Error + "<br/>"))
				f.Flush()
			} else {
				if resp.Status[0:7] == "pulling" {
					percentage := float64(resp.Completed) * 100.0 / float64(resp.Total)
					if percentage < 100.0 && percentage != 0.0 {
						w.Write([]byte(fmt.Sprintf("downloading ... %.1f%%", percentage)))
						f.Flush()
					}
				} else {
					w.Write([]byte(resp.Status + "<br/>"))
					f.Flush()
				}
			}
		}
	}
	return err
}
