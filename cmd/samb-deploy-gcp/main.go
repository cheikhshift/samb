package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cheikhshift/gos/core"
)

func main() {

	projectPath := flag.String("project", "./", "filesystem path of samb project")

	clean := flag.Bool("clean", false, "Remove cached dependencies")

	flag.Parse()

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	err = os.Chdir("./cmd/appengine")

	if err != nil {
		panic(err)
	}

	if *clean {
		cleanDeps()
		return
	}

	if _, err := os.Stat("./vendor"); os.IsNotExist(err) {
		// ./vendor does not exist
		runGoDep()
	}


	res, err := core.RunCmdSmart("gcloud app deploy --quiet")

	if err != nil {
		panic(err)
	}

	fmt.Print(res)
}

func runGoDep() {

	res, err := core.RunCmdSmart("dep init -gopath")

	if err != nil {
		panic(err)
	}

	fmt.Print(res)
}

func cleanDeps(){

	err := os.RemoveAll("./vendor")

	if err != nil {
		panic(err)
	}

	err = os.RemoveAll("./Gopkg.lock")

	if err != nil {
		panic(err)
	}

	err = os.RemoveAll("./Gopkg.toml")

	if err != nil {
		panic(err)
	}
}
