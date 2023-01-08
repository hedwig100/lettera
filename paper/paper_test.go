package paper_test

import (
	"testing"

	"github.com/hedwig100/lettera/bib"
	"github.com/hedwig100/lettera/paper"
)

func TestToMarkdown(t *testing.T) {
	p := paper.Paper{
		Title: "A go tool",
		Bib:   &bib.Bib{CiteKey: "101"},
		Doi:   "100/100",
		Note:  "great",
	}
	expect :=
		`| title | note | doi | bibkey |
| -- | -- | -- | -- |
| A go tool | great | [doi](https://doi.org/100/100) | 101 |`

	if actual := paper.ToMarkdownTable([]paper.Paper{p}); actual != expect {
		t.Errorf("expect '%s' but actually get '%s'", expect, actual)
	}
}
