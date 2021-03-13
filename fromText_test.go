package htmlconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromText(t *testing.T) {

	original1 := `This is a less than "<".
		This is a greater than ">"
		This is a tag that will get displayed <evil tag="true"/><evil>`

	expected1 := `This is a less than &quot;&lt;&quot;.<br> This is a greater than &quot;&gt;&quot;<br> This is a tag that will get displayed &lt;evil tag=&quot;true&quot;/&gt;&lt;evil&gt;`

	assert.Equal(t, expected1, FromText(original1))
}
