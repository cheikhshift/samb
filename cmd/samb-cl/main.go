package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"samb"
	"samb/transpiler"
)

func main() {

	filePath := flag.String("file", "server.se", "path to samb (se) file.")
	projectPath := flag.String("project", "./", "filesystem path of samb project")

	newFile := flag.Bool("new", false, "If true, a new project will be created. First argumenet after will be your package's Go import path.")

	flag.Parse()


	if *newFile {
		args := flag.Args()

		if len(args) == 0 {
			panic("Please provide a Go package import path for your new project.")
		}
		newServer(args[0])
	}

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	file, err := samb.Load(*filePath)

	if err != nil {
		panic(err)
	}

	buildProject(file)

}

func buildProject(file *samb.Project) {
	transpiler.Setup()

	err := transpiler.Transpile(file)

	if err != nil {
		panic(err)
	}

	log.Println("Formatting output")

	formatCode()

	log.Println("Adding imports")
	manageImports()
}

func formatCode() {
	cmd := exec.Command("gofmt", "-w", "./")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func manageImports() {
	cmd := exec.Command("goimports", "-w", "./")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
