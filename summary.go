package html

// Summary returns the first few sentences of content from an HTML document
func Summary(html string) string {

	text := ToText(html)
	text = CollapseWhitespace(text)

	if len(text) > 200 {
		text = text[:200] + "..."
	}

	return text
}
