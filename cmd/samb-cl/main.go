package main

import (
	"flag"
	"os"

	"github.com/cheikhshift/samb"
	"github.com/cheikhshift/samb/tools"
)

func main() {

	filePath := flag.String("file", "server.se", "path to samb (se) file.")
	projectPath := flag.String("project", "./", "filesystem path of samb project")

	newFile := flag.Bool("new", false, "If true, a new project will be created. First argument after will be your package's Go import path.")

	flag.Parse()

	if *newFile {
		args := flag.Args()

		if len(args) == 0 {
			panic("Please provide a Go package import path for your new project.")
		}
		tools.NewServer(args[0])
	}

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	file, err := samb.Load(*filePath)

	if err != nil {
		panic(err)
	}

	tools.BuildProject(file)

}
