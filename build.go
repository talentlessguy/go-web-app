package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
	"github.com/ztrue/tracerr"
)

// CompileToWASM - compile go code to wasm with tinygo
func CompileToWASM(mode string) error {

	cmd := exec.Command("tinygo", "build", "-o", "build/out.wasm", "-target", "wasm", "./src")

	if mode == "go" {
		cmd = exec.Command("go", "build", "-o", "build/out.wasm", "./src/*.go")
		cmd.Env = append(cmd.Env, "GOOS=js", "GOARCH=wasm")
	}

	fmt.Println(chalk.Magenta.Color("\nCompiling to WebAssembly...⌛"))

	fmt.Println(">", cmd.Env, cmd.Args)

	err := cmd.Run()

	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	file, err := os.Open("build/out.wasm")

	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		tracerr.PrintSourceColor(err)
	}

	fmt.Print(
		chalk.Green.Color("\nCompiled successfully! ✨\n"),
		chalk.Yellow.Color("\nBundle Size: "),
		stat.Size()/1024,
		"Kb\n\n",
	)

	return err
}

// CompileToWASMCLI - the same as CompileToWASM but for CLI
func CompileToWASMCLI(c *cli.Context) error {
	return CompileToWASM(c.String("mode"))
}
