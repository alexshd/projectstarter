package generator

import "fmt"

func (g *ViteElmGenerator) packageJsonTemplate(projectName string) string {
	return fmt.Sprintf(`{
  "name": "%s",
  "version": "1.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "test": "elm-test",
    "postinstall": "elm-tooling install"
  },
  "devDependencies": {
    "@tailwindcss/vite": "^4.1.16",
    "elm-tooling": "^1.16.0",
    "tailwindcss": "^4.1.16",
    "vite": "^7.1.12",
    "vite-plugin-elm-watch": "^1.4.3"
  }
}
`, projectName)
}

func (g *ViteElmGenerator) viteConfigTemplate() string {
	return `import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import elmWatch from 'vite-plugin-elm-watch'

export default defineConfig({
  plugins: [
    tailwindcss(),
    elmWatch()
  ]
})
`
}

func (g *ViteElmGenerator) indexHtmlTemplate(projectName string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>%s</title>
</head>
<body>
  <div id="app"></div>
  <script type="module" src="/src/main.js"></script>
</body>
</html>
`, projectName)
}

func (g *ViteElmGenerator) mainJsTemplate() string {
	return `import './style.css'
import { Elm } from './Main.elm'

Elm.Main.init({
  node: document.getElementById('app')
})
`
}

func (g *ViteElmGenerator) styleCssTemplate() string {
	return `@import "tailwindcss";

body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
`
}

func (g *ViteElmGenerator) mainElmTemplate() string {
	return `module Main exposing (main)

import Browser
import Html exposing (Html, div, h1, text, button)
import Html.Attributes exposing (class)
import Html.Events exposing (onClick)


-- MAIN


main : Program () Model Msg
main =
    Browser.sandbox
        { init = init
        , view = view
        , update = update
        }


-- MODEL


type alias Model =
    { count : Int
    }


init : Model
init =
    { count = 0
    }


-- UPDATE


type Msg
    = Increment
    | Decrement


update : Msg -> Model -> Model
update msg model =
    case msg of
        Increment ->
            { model | count = model.count + 1 }

        Decrement ->
            { model | count = model.count - 1 }


-- VIEW


view : Model -> Html Msg
view model =
    div [ class "min-h-screen bg-gray-100 flex items-center justify-center" ]
        [ div [ class "bg-white p-8 rounded-lg shadow-lg" ]
            [ h1 [ class "text-3xl font-bold text-center mb-6 text-gray-800" ]
                [ text "Elm + Vite + Tailwind" ]
            , div [ class "flex items-center justify-center gap-4" ]
                [ button
                    [ onClick Decrement
                    , class "px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
                    ]
                    [ text "-" ]
                , div [ class "text-2xl font-mono w-16 text-center" ]
                    [ text (String.fromInt model.count) ]
                , button
                    [ onClick Increment
                    , class "px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                    ]
                    [ text "+" ]
                ]
            ]
        ]
`
}

func (g *ViteElmGenerator) elmJSONTemplate() string {
	return `{
    "type": "application",
    "source-directories": [
        "src"
    ],
    "elm-version": "0.19.1",
    "dependencies": {
        "direct": {
            "elm/browser": "1.0.2",
            "elm/core": "1.0.5",
            "elm/html": "1.0.0"
        },
        "indirect": {
            "elm/json": "1.1.3",
            "elm/time": "1.0.0",
            "elm/url": "1.0.0",
            "elm/virtual-dom": "1.0.3"
        }
    },
    "test-dependencies": {
        "direct": {},
        "indirect": {}
    }
}
`
}

func (g *ViteElmGenerator) elmToolingJsonTemplate() string {
	return `{
  "tools": {
    "elm": "0.19.1",
    "elm-format": "0.8.7",
    "elm-json": "0.2.13"
  }
}
`
}

func (g *ViteElmGenerator) gitignoreTemplate() string {
	return `# Dependencies
node_modules/
elm-stuff/

# Build output
dist/

# Elm
.elm-spa/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Logs
npm-debug.log*
yarn-debug.log*
yarn-error.log*
`
}

func (g *ViteElmGenerator) readmeTemplate(projectName string) string {
	return fmt.Sprintf(`# %s

Vite + Elm + Tailwind CSS project

## Setup

`+"```bash"+`
npm install
`+"```"+`

## Development

`+"```bash"+`
npm run dev
`+"```"+`

Open http://localhost:5173

## Build

`+"```bash"+`
npm run build
`+"```"+`

## Testing

`+"```bash"+`
npm test
`+"```"+`

## Stack

- [Vite](https://vitejs.dev/) - Build tool
- [Elm](https://elm-lang.org/) - Functional programming language
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [vite-plugin-elm-watch](https://github.com/ChristophP/vite-plugin-elm-watch) - Hot reload for Elm
- [elm-tooling](https://elm-tooling.github.io/elm-tooling-cli/) - Elm tools installer

## License

MIT
`, projectName)
}
