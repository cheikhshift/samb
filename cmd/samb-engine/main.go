package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cheikhshift/samb"
	"github.com/cheikhshift/samb/tools"
)

func main() {

	filePath := flag.String("file", "server.yml", "path to SAMB project entrypoint file.")
	projectPath := flag.String("project", "./", "filesystem path of samb project")

	flag.Parse()

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	file, err := samb.Load(*filePath)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./cmd/appengine", 0700)

	if err != nil {
		panic(err)
	}

	err = writeMainFile(file)

	if err != nil {
		panic(err)
	}

	err = writeConfig()

	if err != nil {
		panic(err)
	}

	tools.ManageImports()

}

func writeMainFile(file *samb.Project) error {

	appengineLauncher := fmt.Sprintf(
		appEngineTemplate,
		file.Package,
		file.Package,
	)

	return ioutil.WriteFile("./cmd/appengine/main.go", []byte(appengineLauncher), 0700)
}

func writeConfig() error {
	return ioutil.WriteFile("./cmd/appengine/app.yaml", appEngineConfigTemplate, 0700)
}
