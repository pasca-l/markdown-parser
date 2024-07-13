package parser

func ParseMarkdownToHtml(content []byte) ([]byte, error) {
	md := newMarkdown(string(content))
	html, err := requestMarkdownRender(md)
	if err != nil {
		return []byte{}, err
	}

	return []byte(html), nil
}
