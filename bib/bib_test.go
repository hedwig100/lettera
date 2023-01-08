package bib_test

import (
	"testing"

	"github.com/hedwig100/lettera/bib"
)

func TestBibParse(t *testing.T) {
	body := `@article{citekey1,
	doi = {10/10},
	year = 2023,
	author = {hedwig100},
	title = {Lettera},
	journal = {reference manager}
}`
	b, err := bib.NewBib(body)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if b.Body != body || b.CiteKey != "citekey1" || b.Title != "Lettera" {
		t.Errorf("unexpected bib, expect (%s,citekey1,Lettera),but actually (%s,%s,%s)", body, b, b.CiteKey, b.Title)
	}
}

func TestBibFetch(t *testing.T) {
	doi := "10.1038/nphys1170"
	expect := `@article{Aspelmeyer_2009,
	doi = {10.1038/nphys1170},
	url = {https://doi.org/10.1038%2Fnphys1170},
	year = 2009,
	month = {jan},
	publisher = {Springer Science and Business Media {LLC}},
	volume = {5},
	number = {1},
	pages = {11--12},
	author = {Markus Aspelmeyer},
	title = {Measured measurement},
	journal = {Nature Physics}
}`

	b, err := bib.FetchBib(doi)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if b.Body != expect {
		t.Errorf("unexpected bib, expect %s,but actually %s", expect, b.Body)
	}
}
