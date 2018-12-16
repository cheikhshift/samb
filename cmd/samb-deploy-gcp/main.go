package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/adlane/exec"
	"github.com/cheikhshift/gos/core"
	"github.com/cheikhshift/samb"
	"github.com/cheikhshift/samb/tools"
)

func main() {

	filePath := flag.String("file", "server.se", "path to samb (se) file.")
	projectPath := flag.String("project", "./", "filesystem path of samb project")

	clean := flag.Bool("clean", false, "Remove cached dependencies")

	build := flag.Bool("build", false, "Transpile project (Generate source code for project).")

	flag.Parse()

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	file, err := samb.Load(*filePath)

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

	if *build {
		fmt.Println("Building project!")
		tools.BuildProject(file)
	}

	if _, err := os.Stat("./vendor"); os.IsNotExist(err) {
		// ./vendor does not exist
		runGoDep()

		err = os.RemoveAll(filepath.Join("vendor", file.Package))

		if err != nil {
			panic(err)
		}
	}

	deployToAppEngine()
}

func runGoDep() {

	res, err := core.RunCmdSmart("dep init -gopath")

	if err != nil {
		panic(err)
	}

	fmt.Print(res)
}

func deployToAppEngine() {
	ctx := exec.InteractiveExec("gcloud", "app", "deploy", "--quiet")
	r := reader{}
	// giving the process all the time needed.
	ctx.Receive(&r, 45*time.Minute)
}

func cleanDeps() {

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
