package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "go-web-app"
	app.Usage = "Simple CLI for setting up Go WebAssembly frontend app."
	app.Authors = []*cli.Author{{
		Name:  "v1rtl",
		Email: "pilll.PL22@gmail.com",
	}}
	app.Version = "0.0.9"
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize a web app",
			Action: InitWebApp,
		},
		{
			Name:   "build",
			Usage:  "Compile Go to WebAssembly",
			Action: CompileToWASMCLI,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "mode",
					Value:   "tinygo",
					Usage:   "Go compiler to use",
					Aliases: []string{"m"},
				},
			},
		},
		{
			Name:   "dev",
			Usage:  "Run a development server",
			Action: RunDevServer,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "port",
					Value:   "8080",
					Usage:   "Dev Server port",
					Aliases: []string{"p"},
				},
			},
		},
	}
}

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run go-web-app, %v", err)
	}
}
