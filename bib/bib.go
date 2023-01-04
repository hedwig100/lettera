package bib

import "os"

type Bib struct {
	body string
}

func (b Bib) String() string {
	return b.body
}

// writeToFile write bib to file
// the file must be opened with os.O_APPEND
func (b *Bib) writeToFile(f *os.File) error {
	_, err := f.Write([]byte(b.body + "\n"))
	return err
}
