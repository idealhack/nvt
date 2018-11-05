package cmd

import (
	"github.com/idealhack/nvt/title"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(titleCmd)
}

var titleCmd = &cobra.Command{
	Use:   "title [files]",
	Short: "Takes markdown files and add title to plain links",
	Long: `Turn https://example.com to [Example Domain](https://example.com/)
It works best when the links are articles in utf-8 encoding.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			title.ProcessFile(arg)
		}
	},
}
