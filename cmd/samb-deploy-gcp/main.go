package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {

	projectPath := flag.String("project", "./", "filesystem path of samb project")

	createProject := flag.Bool("new", false, "Specify if command should create a new project")

	flag.Parse()

	err := os.Chdir(*projectPath)

	if err != nil {
		panic(err)
	}

	err = os.Chdir("./cmd/appengine")

	if err != nil {
		panic(err)
	}

	if *createProject {

		newProject()
		return
	}

	deployProject()
}

func newProject() {

	cmd := exec.Command("gcloud", "app", "create", "--quiet")

	res, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	log.Print(string(res))
}

func deployProject() {

	cmd := exec.Command("gcloud", "app", "deploy", "--quiet")

	res, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	log.Print(string(res))
}
