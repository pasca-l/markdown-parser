package server

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/pasca-l/markdown-parser/parser"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "index.html")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if filepath.Ext(fileHeader.Filename) != ".md" {
		http.Error(w, ".md file expected", http.StatusBadRequest)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parsed, err := parser.ParseMarkdownToHtml(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set(
		"Content-Disposition",
		fmt.Sprintf(
			"attachment; filename=%s.html",
			strings.Trim(
				fileHeader.Filename,
				filepath.Ext(fileHeader.Filename),
			),
		),
	)
	w.Write(parsed)
}
