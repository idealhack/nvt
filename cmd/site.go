package cmd

import (
	"os"

	"github.com/idealhack/nvt/site"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(siteCmd)
}

var siteCmd = &cobra.Command{
	Use:   "site",
	Short: "Takes nvALT notes files and generates a static wiki-like website",
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		site.Check(err)

		site.SetConfig(wd)
		site.ProcessNotes(site.NotesDirectory)
	},
}
