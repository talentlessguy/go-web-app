![go-web-app cover](go-web-app.jpg)

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/talentlessguy/create-go-web-app.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/talentlessguy/create-go-web-app)](https://goreportcard.com/report/github.com/talentlessguy/create-go-web-app)
[![Codacy grade](https://img.shields.io/codacy/grade/c3198e0739ec48bba8902b83a02c3a55.svg?style=flat-square)](https://app.codacy.com/app/komfy/go-web-app)

**Simple CLI for setting up Go WebAssembly frontend app.**

## What's included 📦

- Dev Server
- [TinyGo](https://tinygo.org) for small WebAssembly output, otherwise fallback to `go`
- Git setup
- README file
- Glue files (`index.html` + `wasm_exec.js`)
- Auto-reload

## Requirements ✅

- Go 1.12+
- Browser that supports WebAssembly
- lld (LLVM linker)

## Install 🔄

### Using Go script

```sh
curl -L -o /tmp/install.go http://bit.ly/gwa-setup
go run /tmp/install.go
```

This will install `gwa` into `~/.local/bin` so be sure that `~/.local/bin` is in your `$PATH`.

To add this directory to PATH:

#### Fish

```sh
set -gx PATH $PATH ~/.local/bin
```

#### Bash

```sh
export PATH="$PATH:$HOME/local/bin"
```

### With `go get`

```sh
go get github.com/talentlessguy/go-web-app
```

Then use as `go-web-app` command instead of `gwa`

## Commands 💻

### `gwa init <app name>`

Initialize a project in a picked directory.

#### Project tree

`out.wasm` is generated when building. Other files are automatically added.

```text
├── src
│   └── main.go
├── build
│   └── out.wasm
├── go.mod
├── index.html
├── README.md
└── wasm_exec.js
```

### `gwa dev --port <port>`

Builds the project on first run and launches a development server with specified port.

Default port is **8080**.

After launching a server, you should go to `http://localhost:<port>`

Every time you change a file in `src` dev server automatically compiles and updates the page.

### `gwa build`

Compiles go code to WebAssembly. Compiled `out.wasm` file could be found in `build` folder.

Everything in `src` compiles to `build`, every go file.

After build, binary size is shown

## Stargazers over time

[![Stargazers over time](https://starchart.cc/talentlessguy/go-web-app.svg)](https://starchart.cc/talentlessguy/go-web-app)
