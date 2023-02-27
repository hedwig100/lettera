package paper_test

import (
	"bytes"
	"testing"

	"github.com/hedwig100/lettera/paper"
)

const EXPECT_HEADER = `# Reference List

This file manages a reference paper list. Below is a paper list.

| title | note | doi | bibkey |
| -- | -- | -- | -- |
`

func TestInitMarkdown(t *testing.T) {
	b := &bytes.Buffer{}
	if err := paper.InitMarkdown(b); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if actual := b.String(); actual != EXPECT_HEADER {
		t.Errorf("unexpected actual=%s,expect=%s", actual, EXPECT_HEADER)
	}
}
