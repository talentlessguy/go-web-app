package main

import (
	"fmt"
	"log"

	fs "github.com/fsnotify/fsnotify"
	rl "github.com/talentlessguy/golang-reload-browser"
	"github.com/ttacon/chalk"
	cli "github.com/urfave/cli/v2"
)

// RunDevServer - launches a server for hosting html page with wasm_exec and fetched wasm binaries
func RunDevServer(c *cli.Context) error {
	port := c.String("port")
	watcher, err := fs.NewWatcher()

	mode := c.String("mode")

	defer watcher.Close()

	if err != nil {
		log.Fatalf("Failed to start file watcher, %v", err)
	}

	CompileToWASM(mode)

	fmt.Println(
		chalk.Bold.TextStyle("\n ▶ Listening on http://localhost:"+port),
		chalk.Yellow.Color("\n\nShutdown server: Ctrl + C"),
	)

	rl.StartReloadServer(":" + port)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fs.Write == fs.Write {
					log.Println("modified file:", event.Name)
					log.Println("Reloading browser.")
					CompileToWASM(mode)
					rl.SendReload()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("src")
	if err != nil {
		log.Fatal(err)
	}
	<-done

	return err
}
