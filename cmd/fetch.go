package cmd

import (
	"fmt"

	"github.com/hedwig100/lettera/bib"
	"github.com/spf13/cobra"
)

var (
	doi string

	fetchCmd = &cobra.Command{
		Use:   "fetch",
		Short: "fetch bibtex from https://doi.org/",
		Long:  "fetch bibtex from https://doi.org/",
		RunE: func(cmd *cobra.Command, args []string) error {
			bib, err := bib.FetchBib(doi)
			if err != nil {
				return err
			}
			fmt.Println(bib)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&doi, "doi", "d", "", "DOI of the paper you want to fetch")
	fetchCmd.MarkFlagRequired("doi")
}
