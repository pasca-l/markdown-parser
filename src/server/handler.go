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
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	parsed, err := parser.ParseMarkdownToHtml(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
