package markdown

import (
	"testing"

	"github.com/mickael-menu/zk/core/note"
	"github.com/mickael-menu/zk/util/opt"
	"github.com/mickael-menu/zk/util/test/assert"
)

func TestParseTitle(t *testing.T) {
	test := func(source string, expectedTitle string) {
		content := parse(t, source)
		assert.Equal(t, content.Title, opt.NewNotEmptyString(expectedTitle))
	}

	test("", "")
	test("#  ", "")
	test("#A title", "")
	test("   # A title", "A title")
	test("# A title", "A title")
	test("#   A  title   ", "A  title")
	test("## A title", "A title")
	test("Paragraph \n\n## A title\nBody", "A title")
	test("# Heading 1\n## Heading 1.a\n# Heading 2", "Heading 1")
	test("## Small Heading\n# Bigger Heading", "Bigger Heading")
	test("# A **title** with [formatting](http://stripped)", "A title with formatting")
}

func TestParseBody(t *testing.T) {
	test := func(source string, expectedBody string) {
		content := parse(t, source)
		assert.Equal(t, content.Body, opt.NewNotEmptyString(expectedBody))
	}

	test("", "")
	test("# A title\n    \n", "")
	test("Paragraph \n\n# A title", "")
	test("Paragraph \n\n# A title\nBody", "Body")

	test(
		`## Small Heading
# Bigger Heading
     
## Smaller Heading
Body
several lines
# Body heading

Paragraph:

* item1
* item2
    
`,
		`## Smaller Heading
Body
several lines
# Body heading

Paragraph:

* item1
* item2`,
	)
}

func TestParseLead(t *testing.T) {
	test := func(source string, expectedLead string) {
		content := parse(t, source)
		assert.Equal(t, content.Lead, opt.NewNotEmptyString(expectedLead))
	}

	test("", "")
	test("# A title\n    \n", "")

	test(
		`Paragraph
# A title`,
		"",
	)

	test(
		`Paragraph 
# A title
Lead`,
		"Lead",
	)

	test(
		`# A title
Lead
multiline

other`,
		"Lead\nmultiline",
	)

	test(
		`# A title

Lead
multiline

## Heading`,
		"Lead\nmultiline",
	)

	test(
		`# A title

## Heading
Lead
multiline

other`,
		`## Heading
Lead
multiline`,
	)

	test(
		`# A title

* item1
* item2

Paragraph`,
		`* item1
* item2`,
	)
}

func parse(t *testing.T, source string) note.Content {
	content, err := NewParser().Parse(source)
	assert.Nil(t, err)
	return content
}
