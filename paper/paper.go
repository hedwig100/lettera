package paper

import (
	"strings"

	"github.com/hedwig100/lettera/bib"
)

type DOI = string

type Paper struct {
	Title string
	Bib   *bib.Bib
	Doi   DOI
	Note  string
}

func (p *Paper) String() string {
	return p.Title
}

func (p *Paper) toMarkdownRow() string {
	return "| " + p.Title + " | " + p.Note + " | [doi](" + p.Doi + ") | " + p.Bib.Key + " |"
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
