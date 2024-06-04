package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hako/durafmt"
	"github.com/ollama/ollama/api"
)

const system = "You are an AI assistant who can answer questions and follow instructions."

func predict(model string, query string, doc string, w http.ResponseWriter) (string, error) {
	t0 := time.Now()
	f := w.(http.Flusher)

	client, err := api.ClientFromEnvironment()
	if err != nil {
		return "", err
	}

	// By default, GenerateRequest is streaming.
	req := &api.GenerateRequest{
		Model:  model,
		System: system,
		Prompt: query + "\n\n" + doc,
	}
	results := ""
	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		results += resp.Response
		w.Write([]byte(resp.Response))
		f.Flush()
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		return "", err
	}
	elapsed := durafmt.Parse(time.Since(t0)).LimitFirstN(2)
	fmt.Fprintf(w, "```(%s)```", elapsed)
	f.Flush()
	return results, nil
}
