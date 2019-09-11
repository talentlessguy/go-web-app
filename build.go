package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

// CompileToWASM - compile go code to wasm with tinygo
func CompileToWASM() {

	cmd := exec.Command("tinygo", "build", "-o", "build/out.wasm", "./src")

	fmt.Println(chalk.Magenta.Color("\nCompiling to WebAssembly...⌛"))

	err := cmd.Run()

	if err != nil {
		log.Fatal("Failed compile to WebAssembly")
	}

	file, err := os.Open("build/out.wasm")

	if err != nil {
		log.Fatal("Failed to open WebAssembly output (.wasm)")
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		log.Fatal("Failed to load stats for WebAssembly output")
	}

	fmt.Print(
		chalk.Green.Color("\nCompiled successfully! ✨\n"),
		chalk.Yellow.Color("\nBundle Size: "),
		stat.Size()/1024,
		"Kb\n\n",
	)
}

// CreateIndexHTML creates index.html from template.html
// by replacing %PORT% and %IS_PRODUCTION%.
// if `port` argument is an empty string, isProduction is true.
func CreateIndexHTML(port string) {
	template, err := ioutil.ReadFile("template.html")

	if err != nil {
		log.Fatal(err)
	}

	indexHTMLContents := strings.Replace(string(template), "%PORT%", port, -1)
	isProduction := "false"

	// Disable WebSockets connection when running 'gwa build'
	if len(port) == 0 {
		isProduction = "true"
	}

	indexHTMLContents = strings.Replace(indexHTMLContents, "%IS_PRODUCTION%", isProduction, -1)

	err = ioutil.WriteFile("index.html", []byte(indexHTMLContents), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

// CompileToWASMCLI - the same as CompileToWASM but for CLI
func CompileToWASMCLI(c *cli.Context) {
	CreateIndexHTML("")
	CompileToWASM()
}
