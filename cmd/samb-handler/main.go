package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {

	name := flag.String("name", "FooBar", "Name of Handler to add.")
	project := flag.String("project", "./", "Path to project")

	v := flag.String("verb", "*", "HTTP verb of handler. Leave blank to generate a route handler.")

	flag.Parse()

	os.Chdir(*project)

	// set path to where the handler source will be saved
	savePath := "./pkg/api/" + *name + "_handler.go"

	var router string

	if *v == "*" {

		router = GenerateRouter(*name)

		for _, v := range verbs {

			router += GenerateHandler(*name, v)

		}

	}

	if *v != "*" {

		router = "package api\n\n\n"

		router += GenerateHandler(*name, *v)

	}

	fmt.Println("Saving file to : ", savePath)

	err := ioutil.WriteFile(savePath, []byte(router), 0700)

	if err != nil {
		panic(err)
	}

	manageImports()

}

func GenerateHandler(name, v string) (result string) {
	return fmt.Sprintf(
		handlerTemplate,
		v + strings.Title(name),
		v,
		name,
		v+strings.Title(name),
	)
}

func GenerateRouter(name string) (result string) {
	fileSkeleton := []string{
		name,
		name,
		name,
		name,
	}

	for _, v := range verbs {

		fileSkeleton = append(fileSkeleton, v+strings.Title(name))
	}

	fileSkeleton = append(fileSkeleton, name)

	result = fmt.Sprintf(
		routeTemplate,
		fileSkeleton[0],
		fileSkeleton[1],
		fileSkeleton[2],
		fileSkeleton[3],
		fileSkeleton[4],

		fileSkeleton[5],
		fileSkeleton[6],
		fileSkeleton[7],
		fileSkeleton[8],
		fileSkeleton[9],
	)

	return

}

func manageImports() {
	cmd := exec.Command("goimports", "-w", "./")

	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
