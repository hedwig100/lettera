package bib

type Bib struct {
	body string
}

func (b Bib) String() string {
	return b.body
}
