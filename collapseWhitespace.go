package html

import (
	"regexp"
)

var whitespace *regexp.Regexp

func init() {
	whitespace = regexp.MustCompile(`\s+`)
}

// CollapseWhitespace converts all whitespace characters into a single SPACE character
func CollapseWhitespace(text string) string {
	return whitespace.ReplaceAllString(text, " ")
}
