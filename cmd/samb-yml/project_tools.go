package main

import (
	"log"
	"os/exec"

	"github.com/cheikhshift/samb"
	"github.com/cheikhshift/samb/transpiler"
)

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

func newServer(p string) {
	println(p)
}
