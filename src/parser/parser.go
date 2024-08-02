package parser

import (
	"net/url"
)

func ParseMarkdownToHtml(content []byte, params url.Values) ([]byte, error) {
	md := newMarkdown(string(content))
	html, err := requestMarkdownRender(md)
	if err != nil {
		return []byte{}, err
	}

	html, err = convertHtmlContent(html, params)
	if err != nil {
		return []byte{}, err
	}

	return []byte(html), nil
}
