package bib

import (
	"errors"
	"os"
	"strings"

	"github.com/jschaf/bibtex"
)

type Bib struct {
	Body string
	Key  bibtex.CiteKey
}

func parseBibBody(body string) (bibtex.Entry, error) {
	parser := bibtex.New()
	bibAst, err := parser.Parse(strings.NewReader(body))
	if err != nil {
		return bibtex.Entry{}, err
	}
	bibEntry, err := parser.Resolve(bibAst)
	if err != nil {
		return bibtex.Entry{}, err
	}
	if len(bibEntry) != 1 {
		return bibtex.Entry{}, errors.New("this bib body has more than one body")
	}
	return bibEntry[0], nil
}

func NewBib(body string) (*Bib, error) {
	bibEntry, err := parseBibBody(body)
	if err != nil {
		return &Bib{}, err
	}
	return &Bib{
		Body: body,
		Key:  bibEntry.Key,
	}, nil
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
