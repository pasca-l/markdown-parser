package parser

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Markdown struct {
	Text string `json:"text"`
}

func newMarkdown(text string) Markdown {
	return Markdown{Text: text}
}

func requestMarkdownRender(md Markdown) (string, error) {
	url := "https://api.github.com/markdown"

	body, err := json.Marshal(md)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if err != nil {
		return "", err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
