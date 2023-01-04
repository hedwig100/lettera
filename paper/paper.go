package paper

import (
	"fmt"
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

func MarkDownTableHeader() string {
	return "| title | note | doi | bibkey | \n| -- | -- | -- | -- |"
}

func ToMarkdownTable(papers []Paper) string {
	cols := make([]string, len(papers))
	for i, paper := range papers {
		fmt.Println(paper.Bib.Key)
		cols[i] = "| " + paper.Title + " | " + paper.Note + " | [doi](" + paper.Doi + ") |" + paper.Bib.Key + " |"
	}
	return strings.Join(cols, "\n")
}
