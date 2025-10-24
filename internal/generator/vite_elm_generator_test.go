package generator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestViteElmGenerator_Generate(t *testing.T) {
	t.Run("creates project successfully", func(t *testing.T) {
		tmpDir := t.TempDir()
		projectName := filepath.Join(tmpDir, "test-project")

		gen := NewViteElmGenerator()
		err := gen.Generate(projectName)
		if err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		// Verify directory was created
		if _, err := os.Stat(projectName); os.IsNotExist(err) {
			t.Error("Project directory was not created")
		}
	})

	t.Run("fails when directory already exists", func(t *testing.T) {
		tmpDir := t.TempDir()
		projectName := filepath.Join(tmpDir, "existing-project")

		// Create directory first
		if err := os.MkdirAll(projectName, 0o755); err != nil {
			t.Fatalf("Failed to create test directory: %v", err)
		}

		gen := NewViteElmGenerator()
		err := gen.Generate(projectName)

		if err == nil {
			t.Error("Expected error when directory exists, got nil")
		}

		if !strings.Contains(err.Error(), "already exists") {
			t.Errorf("Expected 'already exists' error, got: %v", err)
		}
	})

	t.Run("creates all required directories", func(t *testing.T) {
		tmpDir := t.TempDir()
		projectName := filepath.Join(tmpDir, "test-dirs")

		gen := NewViteElmGenerator()
		if err := gen.Generate(projectName); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		requiredDirs := []string{
			projectName,
			filepath.Join(projectName, "src"),
			filepath.Join(projectName, "public"),
		}

		for _, dir := range requiredDirs {
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Errorf("Required directory not created: %s", dir)
			}
		}
	})

	t.Run("creates all required files", func(t *testing.T) {
		tmpDir := t.TempDir()
		projectName := filepath.Join(tmpDir, "test-files")

		gen := NewViteElmGenerator()
		if err := gen.Generate(projectName); err != nil {
			t.Fatalf("Generate() failed: %v", err)
		}

		requiredFiles := []string{
			"package.json",
			"vite.config.js",
			"index.html",
			"src/main.js",
			"src/style.css",
			"src/Main.elm",
			"elm.json",
			"elm-tooling.json",
			".gitignore",
			"README.md",
		}

		for _, file := range requiredFiles {
			fullPath := filepath.Join(projectName, file)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				t.Errorf("Required file not created: %s", file)
			}
		}
	})
}

func TestViteElmGenerator_PackageJson(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.packageJsonTemplate("my-app")

	// Verify it's valid JSON
	var pkg map[string]interface{}
	if err := json.Unmarshal([]byte(content), &pkg); err != nil {
		t.Fatalf("packageJsonTemplate() produced invalid JSON: %v", err)
	}

	t.Run("has correct name", func(t *testing.T) {
		if pkg["name"] != "my-app" {
			t.Errorf("Expected name 'my-app', got '%v'", pkg["name"])
		}
	})

	t.Run("has module type", func(t *testing.T) {
		if pkg["type"] != "module" {
			t.Errorf("Expected type 'module', got '%v'", pkg["type"])
		}
	})

	t.Run("has required scripts", func(t *testing.T) {
		scripts, ok := pkg["scripts"].(map[string]interface{})
		if !ok {
			t.Fatal("scripts field is not a map")
		}

		requiredScripts := []string{"dev", "build", "test", "postinstall"}
		for _, script := range requiredScripts {
			if _, exists := scripts[script]; !exists {
				t.Errorf("Missing required script: %s", script)
			}
		}

		// Verify postinstall runs elm-tooling
		if scripts["postinstall"] != "elm-tooling install" {
			t.Errorf("postinstall script incorrect: %v", scripts["postinstall"])
		}
	})

	t.Run("has required dependencies", func(t *testing.T) {
		devDeps, ok := pkg["devDependencies"].(map[string]interface{})
		if !ok {
			t.Fatal("devDependencies field is not a map")
		}

		requiredDeps := []string{
			"@tailwindcss/vite",
			"elm-tooling",
			"tailwindcss",
			"vite",
			"vite-plugin-elm-watch",
		}

		for _, dep := range requiredDeps {
			if _, exists := devDeps[dep]; !exists {
				t.Errorf("Missing required dependency: %s", dep)
			}
		}
	})
}

func TestViteElmGenerator_ViteConfig(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.viteConfigTemplate()

	t.Run("imports required plugins", func(t *testing.T) {
		requiredImports := []string{
			"defineConfig",
			"@tailwindcss/vite",
			"vite-plugin-elm-watch",
		}

		for _, imp := range requiredImports {
			if !strings.Contains(content, imp) {
				t.Errorf("Missing required import: %s", imp)
			}
		}
	})

	t.Run("configures plugins", func(t *testing.T) {
		if !strings.Contains(content, "tailwindcss()") {
			t.Error("Missing tailwindcss plugin configuration")
		}
		if !strings.Contains(content, "elmWatch()") {
			t.Error("Missing elmWatch plugin configuration")
		}
	})
}

func TestViteElmGenerator_StyleCss(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.styleCssTemplate()

	t.Run("imports tailwindcss", func(t *testing.T) {
		if !strings.Contains(content, `@import "tailwindcss"`) {
			t.Error("style.css doesn't import tailwindcss")
		}
	})
}

func TestViteElmGenerator_IndexHtml(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.indexHtmlTemplate("Test App")

	t.Run("has correct title", func(t *testing.T) {
		if !strings.Contains(content, "<title>Test App</title>") {
			t.Error("Index.html doesn't have correct title")
		}
	})

	t.Run("has app mount point", func(t *testing.T) {
		if !strings.Contains(content, `id="app"`) {
			t.Error("Index.html doesn't have app mount point")
		}
	})

	t.Run("loads main.js", func(t *testing.T) {
		if !strings.Contains(content, "/src/main.js") {
			t.Error("Index.html doesn't load main.js")
		}
	})
}

func TestViteElmGenerator_MainJs(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.mainJsTemplate()

	t.Run("imports style.css", func(t *testing.T) {
		if !strings.Contains(content, "import './style.css'") {
			t.Error("main.js doesn't import style.css")
		}
	})

	t.Run("imports Elm Main module", func(t *testing.T) {
		if !strings.Contains(content, "import { Elm } from './Main.elm'") {
			t.Error("main.js doesn't import Main.elm")
		}
	})

	t.Run("initializes Elm app", func(t *testing.T) {
		if !strings.Contains(content, "Elm.Main.init") {
			t.Error("main.js doesn't initialize Elm app")
		}
	})
}

func TestViteElmGenerator_MainElm(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.mainElmTemplate()

	t.Run("declares Main module", func(t *testing.T) {
		if !strings.Contains(content, "module Main exposing (main)") {
			t.Error("Main.elm doesn't declare Main module")
		}
	})

	t.Run("has Browser.sandbox", func(t *testing.T) {
		if !strings.Contains(content, "Browser.sandbox") {
			t.Error("Main.elm doesn't use Browser.sandbox")
		}
	})

	t.Run("has Model type", func(t *testing.T) {
		if !strings.Contains(content, "type alias Model") {
			t.Error("Main.elm doesn't define Model type")
		}
	})

	t.Run("has Msg type", func(t *testing.T) {
		if !strings.Contains(content, "type Msg") {
			t.Error("Main.elm doesn't define Msg type")
		}
	})

	t.Run("uses Tailwind classes", func(t *testing.T) {
		tailwindClasses := []string{"bg-gray-100", "rounded", "hover:bg"}
		for _, class := range tailwindClasses {
			if !strings.Contains(content, class) {
				t.Errorf("Main.elm doesn't use Tailwind class: %s", class)
			}
		}
	})
}

func TestViteElmGenerator_ElmJson(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.elmJSONTemplate()

	// Verify it's valid JSON
	var elmJson map[string]interface{}
	if err := json.Unmarshal([]byte(content), &elmJson); err != nil {
		t.Fatalf("elmJSONTemplate() produced invalid JSON: %v", err)
	}

	t.Run("is application type", func(t *testing.T) {
		if elmJson["type"] != "application" {
			t.Errorf("Expected type 'application', got '%v'", elmJson["type"])
		}
	})

	t.Run("has src directory", func(t *testing.T) {
		dirs, ok := elmJson["source-directories"].([]interface{})
		if !ok || len(dirs) == 0 {
			t.Fatal("source-directories not properly configured")
		}
		if dirs[0] != "src" {
			t.Errorf("Expected 'src' directory, got '%v'", dirs[0])
		}
	})

	t.Run("has browser dependencies", func(t *testing.T) {
		deps := elmJson["dependencies"].(map[string]interface{})
		direct := deps["direct"].(map[string]interface{})

		requiredDeps := []string{"elm/browser", "elm/core", "elm/html"}
		for _, dep := range requiredDeps {
			if _, exists := direct[dep]; !exists {
				t.Errorf("Missing required dependency: %s", dep)
			}
		}
	})
}

func TestViteElmGenerator_ElmToolingJson(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.elmToolingJsonTemplate()

	// Verify it's valid JSON
	var tooling map[string]interface{}
	if err := json.Unmarshal([]byte(content), &tooling); err != nil {
		t.Fatalf("elmToolingJsonTemplate() produced invalid JSON: %v", err)
	}

	t.Run("has required tools", func(t *testing.T) {
		tools, ok := tooling["tools"].(map[string]interface{})
		if !ok {
			t.Fatal("tools field is not a map")
		}

		requiredTools := []string{"elm", "elm-format", "elm-json"}
		for _, tool := range requiredTools {
			if _, exists := tools[tool]; !exists {
				t.Errorf("Missing required tool: %s", tool)
			}
		}
	})
}

func TestViteElmGenerator_Gitignore(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.gitignoreTemplate()

	t.Run("ignores node_modules", func(t *testing.T) {
		if !strings.Contains(content, "node_modules/") {
			t.Error(".gitignore doesn't ignore node_modules")
		}
	})

	t.Run("ignores elm-stuff", func(t *testing.T) {
		if !strings.Contains(content, "elm-stuff/") {
			t.Error(".gitignore doesn't ignore elm-stuff")
		}
	})

	t.Run("ignores dist", func(t *testing.T) {
		if !strings.Contains(content, "dist/") {
			t.Error(".gitignore doesn't ignore dist")
		}
	})
}

func TestViteElmGenerator_Readme(t *testing.T) {
	gen := NewViteElmGenerator()
	content := gen.readmeTemplate("Test Project")

	t.Run("has project name as title", func(t *testing.T) {
		if !strings.Contains(content, "# Test Project") {
			t.Error("README doesn't have project name as title")
		}
	})

	t.Run("has setup instructions", func(t *testing.T) {
		if !strings.Contains(content, "npm install") {
			t.Error("README doesn't have setup instructions")
		}
	})

	t.Run("has development instructions", func(t *testing.T) {
		if !strings.Contains(content, "npm run dev") {
			t.Error("README doesn't have development instructions")
		}
	})

	t.Run("documents the stack", func(t *testing.T) {
		stack := []string{"Vite", "Elm", "Tailwind", "vite-plugin-elm-watch"}
		for _, tech := range stack {
			if !strings.Contains(content, tech) {
				t.Errorf("README doesn't mention %s", tech)
			}
		}
	})
}
