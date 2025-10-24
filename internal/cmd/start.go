package cmd

import (
	"fmt"
	"log/slog"

	"github.com/alexshd/projectstarter/internal/generator"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new project",
	Long:  `Create a new project with proper structure and boilerplate code.`,
}

var startGoCmd = &cobra.Command{
	Use:   "go <project-name>",
	Short: "Create a new Go project",
	Long: `Create a new Go project with:
  - cmd/projectname/main.go (working code)
  - internal/ (ready for packages)
  - README.md
  - LICENSE (MIT)
  - go.mod
  - .gitignore
  - Basic passing test`,
	Example: `  # Create project with short name
  proj start go myapp

  # Create project with full module path
  proj start go github.com/user/myapp`,
	Args: cobra.ExactArgs(1),
	RunE: runStartGo,
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(startGoCmd)
}

func runStartGo(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	slog.Info("Creating Go project", "name", projectName)

	gen := generator.NewGoGenerator()
	if err := gen.Generate(projectName); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	slog.Info("✅ Project created successfully!", "name", projectName)
	slog.Info("Next steps:",
		"cd", projectName,
		"run", "go mod tidy && go run cmd/"+projectName+"/main.go")

	return nil
}
