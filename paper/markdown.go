package paper

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const HEADER = `# Reference List

This file manages a reference paper list. Below is a paper list.

`
const MARKDOWN_TABLE_HEADER = "| title | note | doi | bibkey |\n| -- | -- | -- | -- |\n"

const MARKDOWN_HEADER = HEADER + MARKDOWN_TABLE_HEADER

func WriteToMarkdown(p *Paper, filepath string) error {
	if !strings.HasSuffix(filepath, ".md") {
		return fmt.Errorf("unexpected filepath %s, the filepath should be end with '.md'", filepath)
	}

	exist, err := isExist(filepath)
	if err != nil {
		return err
	}

	if exist {
		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		return appendToMarkdown(p, f)
	} else {
		f, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer f.Close()
		err = initMarkdown(f)
		if err != nil {
			return err
		}
		return appendToMarkdown(p, f)
	}
}

func (p *Paper) toMarkdownRow() string {
	return "| " + p.Title + " | " + p.Note + " | [doi](https://doi.org/" + p.Doi + ") | " + p.Bib.CiteKey + " |"
}

func initMarkdown(w io.Writer) error {
	_, err := w.Write([]byte(MARKDOWN_HEADER))
	return err
}

func appendToMarkdown(p *Paper, w io.Writer) error {
	_, err := w.Write([]byte(p.toMarkdownRow() + "\n"))
	return err
}
