package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set via -ldflags at build time
var Version = "dev"

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
