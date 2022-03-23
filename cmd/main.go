package main

import (
	"fmt"
	"hello/cmd/app"
	"os"
)

var (
	gitHash   string
	gitTag    string
	buildTime string
	goVersion string
)

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Tag: %s \n", gitTag)
		fmt.Printf("Git Commit hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
		return
	}

	app.RunApp()
}
