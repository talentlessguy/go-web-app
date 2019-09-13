package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Downloading go-web-app binary into ~/.local/bin")

	dir, err := os.UserHomeDir()

	if err != nil {
		panic("Failed to get home directory")
	}
	os.Chdir(dir + "/.local/bin")
	out, err := os.Create("gwa")
	defer out.Close()

	if err != nil {
		panic("Failed to create output file")
	}

	fmt.Println("Fetching from GitHub...")

	resp, err := http.Get("https://github.com/talentlessguy/go-web-app/raw/master/bin/gwa")

	if err != nil {
		panic("Failed to fetch binary")
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic("Failed to save gwa binary")
	}

	err = os.Chmod("gwa", 0777)

	if err != nil {
		panic("Couldn't make gwa executable")
	}

	fmt.Println("Done")
}
