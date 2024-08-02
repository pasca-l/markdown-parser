package parser

import (
	"net/url"
	"regexp"
)

type Conversion struct {
	Ticked    bool
	Converter func(string) (string, error)
}
type Conversions map[string]Conversion

func (c Conversions) applyConversion(content string) (string, error) {
	var err error
	for _, v := range c {
		if v.Ticked {
			content, err = v.Converter(content)
			if err != nil {
				return "", err
			}
		}
	}
	return content, nil
}

func checkTicked(tick string) bool {
	switch tick {
	case "on":
		return true
	case "off":
		return false
	default:
		return false
	}
}

func newConversionParam(params url.Values) Conversions {
	return Conversions{
		"className": {
			Ticked:    checkTicked(params.Get("className")),
			Converter: convertToClassName,
		},
	}
}

func convertHtmlContent(html string, params url.Values) (string, error) {
	conversion := newConversionParam(params)
	converted, err := conversion.applyConversion(html)
	if err != nil {
		return html, err
	}
	return converted, nil
}

func convertByReplacement(content string, replacement map[string]string) (string, error) {
	for pattern, replace := range replacement {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		content = re.ReplaceAllString(content, replace)
	}
	return content, nil
}

func convertToClassName(content string) (string, error) {
	replacement := map[string]string{
		`class=(".*?")`: `className=$1`,
	}
	return convertByReplacement(content, replacement)
}
