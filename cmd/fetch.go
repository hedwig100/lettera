package cmd

import (
	"fmt"

	"github.com/hedwig100/lettera/bib"
	"github.com/hedwig100/lettera/paper"
	"github.com/spf13/cobra"
)

var (
	doi          string
	bibfilePath  string
	markdownPath string
	note         string

	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "fetch bibtex from https://doi.org/",
		Long:  "fetch bibtex from https://doi.org/",
		RunE: func(cmd *cobra.Command, args []string) error {
			b, err := bib.FetchBib(doi)
			if err != nil {
				return err
			}
			err = bib.WriteToFile(b, bibfilePath)
			if err != nil {
				return err
			}
			p := paper.NewPaper(b, doi, note)
			err = paper.WriteToMarkdown(p, markdownPath)
			if err != nil {
				return err
			}
			fmt.Println(b)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&doi, "doi", "d", "", "DOI of the paper you want to fetch")
	fetchCmd.Flags().StringVarP(&bibfilePath, "bibfile", "b", "lettera.bib", "filepath that you want to write the fetched bib data to")
	fetchCmd.Flags().StringVarP(&markdownPath, "markdown", "m", "lettera.md", "markdown filepath that you want to write the fetched paperdata to")
	fetchCmd.Flags().StringVarP(&note, "note", "n", "", "what you want to note about the reference paper")
	fetchCmd.MarkFlagRequired("doi")
}
