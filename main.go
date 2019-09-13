package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "go-web-app"
	app.Usage = "Simple CLI for setting up Go WebAssembly frontend app."
	app.Author = "v1rtl (twitter.com/v1rtl)"
	app.Version = "0.0.7"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "Initialize a web app",
			Action: InitWebApp,
		},
		{
			Name:   "build",
			Usage:  "Compile Go to WebAssembly",
			Action: CompileToWASMCLI,
		},
		{
			Name:   "dev",
			Usage:  "Run a development server",
			Action: RunDevServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "port",
					Value: "8080",
					Usage: "Dev Server port",
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
		log.Fatal("Failed to run go-web-app")
	}
}
