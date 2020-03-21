package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummary(t *testing.T) {

	content1 := "<p>This is the first paragraph of the content</p><p>This is the second paragraph of the content</p>"
	content2 := "<p>This is the first, but really, really long paragraph of the content.  I'm testing to see that we can still truncate to 200 characters even though the first paragraph says otherwise.  I wonder what else I'll have to say in order to make this truncate correctly?</p><p>Here's the second paragraph</p>"

	assert.Equal(t, "This is the first paragraph of the content", Summary(content1))
	assert.Equal(t, "This is the first, but really, really long paragraph of the content. I'm testing to see that we can still truncate to 200 characters even though the first paragraph says otherwise. I wonder what else ...", Summary(content2))
}
