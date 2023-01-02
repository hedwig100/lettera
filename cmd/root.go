package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Lettera",
	Short: "Lettera is a reference manager",
	Long: ` Lettera is a reference manager.
 You can manage bibtex,pdf of your reference papers`,
}

func Execute() error {
	return rootCmd.Execute()
}
