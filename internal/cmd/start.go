package cmd

import (
	"fmt"
	"log/slog"

	"github.com/alexshd/projectstarter/internal/generator"
	"github.com/fatih/color"
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

var startViteElmCmd = &cobra.Command{
	Use:   "vite-elm <project-name>",
	Short: "Create a new Vite + Elm + Tailwind project",
	Long: `Create a new Vite + Elm + Tailwind CSS project with:
  - Vite build setup
  - Elm with hot reload (vite-plugin-elm-watch)
  - Tailwind CSS with @tailwindcss/vite plugin
  - elm-tooling for tool management
  - Working counter example
  - package.json with dev, build, test scripts`,
	Example: `  # Create new Vite + Elm project
  proj start vite-elm myapp`,
	Args: cobra.ExactArgs(1),
	RunE: runStartViteElm,
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.AddCommand(startGoCmd)
	startCmd.AddCommand(startViteElmCmd)
}

func runStartGo(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	slog.Info("Creating Go project", "name", projectName)

	gen := generator.NewGoGenerator()
	if err := gen.Generate(projectName); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	fmt.Println()
	color.Green("🐹 Go project created successfully!")
	fmt.Println()
	color.Cyan("🚀 Next steps:")
	color.Yellow("   cd %s", projectName)
	color.Yellow("   go mod tidy && go run cmd/%s/main.go", projectName)
	fmt.Println()

	return nil
}

func runStartViteElm(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	slog.Info("Creating Vite + Elm + Tailwind project", "name", projectName)

	gen := generator.NewViteElmGenerator()
	if err := gen.Generate(projectName); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	fmt.Println()
	color.Green("🌳 Elm project created successfully!")
	fmt.Println()
	color.Cyan("⚡ Next steps:")
	color.Yellow("   cd %s", projectName)
	color.Yellow("   npm install && npm run dev")
	fmt.Println()

	return nil
}
