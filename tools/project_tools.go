package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/cheikhshift/samb"
	"github.com/cheikhshift/samb/transpiler"
)

func BuildProject(file *samb.Project) {
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

func NewServer(p string) {

	ioutil.WriteFile("./server.se", []byte(fmt.Sprintf(serverTemplate, p)), 0700)
	ioutil.WriteFile("./endpoints.se", []byte(routeTemplate), 0700)
	ioutil.WriteFile("./providers.se", []byte(providerTemplate), 0700)
}
