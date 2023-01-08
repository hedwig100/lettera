package paper

import (
	"fmt"
	"os"
	"strings"
)

const MARKDOWN_HEADER = `# Reference List

This file manages a reference paper list. Below is a paper list.

`

func WriteToMarkdown(p *Paper, filepath string) error {
	if !strings.HasSuffix(filepath, ".md") {
		return fmt.Errorf("unexpected filepath %s, the filepath should be end with '.md'", filepath)
	}

	exist, err := isExist(filepath)
	if err != nil {
		return err
	}

	if exist {
		return appendToMarkdown(p, filepath)
	} else {
		return initMarkdown([]Paper{*p}, filepath)
	}
}

func (p *Paper) toMarkdownRow() string {
	return "| " + p.Title + " | " + p.Note + " | [doi](https://doi.org/" + p.Doi + ") | " + p.Bib.CiteKey + " |"
}

func markdownTableHeader() string {
	return "| title | note | doi | bibkey |\n| -- | -- | -- | -- |"
}

func ToMarkdownTable(papers []Paper) string {
	rows := make([]string, len(papers)+1)
	rows[0] = markdownTableHeader()
	for i, paper := range papers {
		rows[i+1] = paper.toMarkdownRow()
	}
	return strings.Join(rows, "\n")
}

func initMarkdown(papers []Paper, filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(MARKDOWN_HEADER))
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(ToMarkdownTable(papers) + "\n"))
	return err
}

func appendToMarkdown(p *Paper, filepath string) error {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(p.toMarkdownRow() + "\n"))
	return err
}
