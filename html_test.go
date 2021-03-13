package htmlconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveTags(t *testing.T) {
	assert.Equal(t, "Regular string", RemoveTags("Regular string"))
	assert.Equal(t, "Regular string with tags", RemoveTags("Regular string <b>with tags</b>"))
	assert.Equal(t, "Regular string with tags and attributes.", RemoveTags(`Regular string <span class="some-class">with tags</span> and <i>attributes</i>.`))
}

func TestIsHTML(t *testing.T) {
	assert.Equal(t, true, IsHTML("<b>This is HTML</b>"))
	assert.Equal(t, false, IsHTML("This is not HTML"))
}

func TestToText(t *testing.T) {
	assert.Equal(t, "Hello Gordon", ToText("<i>Hello</i> <b>Gordon</b>"))
	assert.Equal(t, "Without Stylesheets", ToText("<style>* {font-weight:bold;}</style> Without Stylesheets"))
	assert.Equal(t, "...", ToText("&hellip;"))
}
