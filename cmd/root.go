package cmd

import "github.com/spf13/cobra"

import (
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nvt",
	Short: "nvALT tools",
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
