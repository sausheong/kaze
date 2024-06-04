package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func convert(inputpdf string) (string, error) {
	tempdir, err := os.MkdirTemp("", "kaze")
	if err != nil {
		log.Println("unable to create a temporary directory:", err)
		return "", err
	}
	defer os.RemoveAll(tempdir)

	cmd := exec.Command(filepath.Join(exedir, "bin", "pdftotext"), inputpdf, filepath.Join(tempdir, "output.txt"))
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command error: %s\n", err)
		return "", err
	}

	text, err := os.ReadFile(filepath.Join(tempdir, "output.txt"))
	if err != nil {
		log.Printf("cannot read text: %s\n", err)
		return "", err
	}
	content := string(text)
	content = strings.ToValidUTF8(content, "")
	return content, nil
}

func clean(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if r <= 127 {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
