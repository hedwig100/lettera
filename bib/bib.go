package bib

import (
	"errors"
	"os"
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

// writeToFile write bib to file
// the file must be opened with os.O_APPEND
func (b *Bib) writeToFile(f *os.File) error {
	_, err := f.Write([]byte(b.Body + "\n"))
	return err
}
