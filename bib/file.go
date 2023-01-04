package bib

import (
	"fmt"
	"os"
	"strings"
)

func WriteToFile(b Bib, filepath string) error {
	if !strings.HasSuffix(filepath, ".bib") {
		return fmt.Errorf("unexpected filepath %s, the filepath should be end with '.bib'", filepath)
	}
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return b.writeToFile(f)
}
