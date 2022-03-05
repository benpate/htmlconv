package htmlconv

import (
	"regexp"
)

var firstParagraph *regexp.Regexp

func init() {
	firstParagraph = regexp.MustCompile("^(.*?)<(br|/p|/div)>")
}

// Summary returns the first few sentences of content from an HTML document
func Summary(html string) string {

	result := firstParagraph.FindString(html)

	// Remove any extra whitespace now.
	result = CollapseWhitespace(RemoveTags(result))

	// If it's STILL too long, then truncate to 200 characters.
	if len(result) > 200 {
		result = result[:200] + "..."
	}

	return result
}
