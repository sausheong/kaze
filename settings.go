package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

var envString = `HISTORY_FILE=""
MODEL="llama-3"
PORT=1234`

func getEnvSettings() map[string]string {
	homedir, _ := os.UserHomeDir()
	env, err := godotenv.Read(filepath.Join(homedir, ".kaze"))
	if err != nil {
		log.Println("Error loading .kaze file")
	}
	return env
}

func putEnvSettings(env map[string]string) error {
	homedir, _ := os.UserHomeDir()
	err := godotenv.Write(env, filepath.Join(homedir, ".kaze"))
	if err != nil {
		log.Println("Error loading .kaze file")
	}
	err = godotenv.Overload(filepath.Join(homedir, ".kaze"))
	if err != nil {
		log.Println("Error loading .kaze file into environment")
	}
	return err
}

func writeToHistory(input string, response string) {
	env := getEnvSettings()
	file, err := os.OpenFile(env["HISTORY_FILE"], os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("failed to open file: %v\n", err)
	}
	defer file.Close()

	response = fmt.Sprintf("## %s\n\n %s\n [%s]\n\n", input, response, time.Now().Format("Jan 2, 2006 15:04:05"))
	_, err = file.WriteString(response)
	if err != nil {
		log.Printf("failed to write to file: %v\n", err)
	}
}

func getDefaultHistoryFile() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, "kaze.md")
}

func createSettingsEnv() error {
	homedir, _ := os.UserHomeDir()
	err := os.WriteFile(filepath.Join(homedir, ".kaze"), []byte(envString), 0644)
	if err != nil {
		log.Println("cannot create .kaze file:", err)
		return err
	}
	return nil
}
