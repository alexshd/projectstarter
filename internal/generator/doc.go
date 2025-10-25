// Package generator provides project scaffolding generators for different project types.
//
// The generator package contains implementations for creating new projects with
// proper structure, boilerplate code, and configuration files.
//
// # Available Generators
//
// GoGenerator creates Go projects with:
//   - cmd/projectname/main.go with working code
//   - internal/ directory for packages
//   - go.mod with proper module path
//   - README.md, LICENSE, .gitignore
//   - Basic passing test
//
// ViteElmGenerator creates Vite + Elm + Tailwind projects with:
//   - Vite build setup with hot reload
//   - Elm with vite-plugin-elm-watch
//   - Tailwind CSS v4 with @tailwindcss/vite plugin
//   - elm-tooling for tool management
//   - Working counter example with Tailwind styling
//
// # Usage
//
//	// Create a Go project
//	gen := generator.NewGoGenerator()
//	err := gen.Generate("myapp")
//
//	// Create a Vite + Elm project
//	gen := generator.NewViteElmGenerator()
//	err := gen.Generate("my-elm-app")
//
// # Design
//
// Each generator follows a consistent pattern:
//   - NewXGenerator() constructor returns a generator instance
//   - Generate(projectName) creates the project structure
//   - Template methods provide file contents
//   - All generators are thoroughly tested
package generator
