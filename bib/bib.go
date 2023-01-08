package bib

import (
	"errors"
	"io"
	"strings"

	"github.com/nickng/bibtex"
)

type Bib struct {
	Body    string
	Title   string
	CiteKey string
}

func parseBibBody(body string) (*bibtex.BibEntry, error) {
	bibtexBody, err := bibtex.Parse(strings.NewReader(body))
	if err != nil {
		return &bibtex.BibEntry{}, err
	}
	if len(bibtexBody.Entries) != 1 {
		return &bibtex.BibEntry{}, errors.New("this bib body has more than one body")
	}
	return bibtexBody.Entries[0], nil
}

func NewBib(body string) (*Bib, error) {
	bibEntry, err := parseBibBody(body)
	if err != nil {
		return &Bib{}, err
	}

	if title, ok := bibEntry.Fields["title"]; ok {
		return &Bib{
			Body:    body,
			Title:   title.String(),
			CiteKey: bibEntry.CiteName,
		}, nil
	} else {
		return &Bib{
			Body:    body,
			CiteKey: bibEntry.CiteName,
		}, nil
	}
}

func (b *Bib) String() string {
	return b.Body
}

// writeln write bib to io.Writer with newline
func (b *Bib) writeln(w io.Writer) error {
	_, err := w.Write([]byte(b.Body + "\n"))
	return err
}
