package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proj",
	Short: "Project scaffolding tool",
	Long:  `proj helps you quickly create new projects with proper structure and boilerplate code.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add commands here
}
