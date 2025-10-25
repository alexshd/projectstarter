package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version of the application - update this manually when tagging a release
var Version = "0.3.1"

var rootCmd = &cobra.Command{
	Use:     "proj",
	Short:   "Project scaffolding tool",
	Long:    `proj helps you quickly create new projects with proper structure and boilerplate code.`,
	Version: Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Customize version template to show just the version
	rootCmd.SetVersionTemplate(fmt.Sprintf("proj version %s\n", Version))
}
