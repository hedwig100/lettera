package paper

import (
	"github.com/hedwig100/lettera/bib"
)

type DOI = string

type Paper struct {
	Title string
	Bib   *bib.Bib
	Doi   DOI
	Note  string
}

func NewPaper(b *bib.Bib, doi DOI, note string) *Paper {
	return &Paper{
		Title: b.Title,
		Bib:   b,
		Doi:   doi,
		Note:  note,
	}
}

func (p *Paper) String() string {
	return p.Title
}
