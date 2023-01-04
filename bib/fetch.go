package bib

import (
	"io"
	"net/http"
)

func FetchBib(doi string) (*Bib, error) {
	client := &http.Client{}

	url := "https://doi.org/" + doi
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &Bib{}, err
	}
	req.Header.Add("Accept", "application/x-bibtex")

	resp, err := client.Do(req)
	if err != nil {
		return &Bib{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Bib{}, err
	}
	return NewBib(string(body))
}
