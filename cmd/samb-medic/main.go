package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	name := flag.String("name", "FooBar", "Name of recovery function to add.")
	project := flag.String("project", "./", "Path to project")

	flag.Parse()

	os.Chdir(*project)

	// set path to where the handler source will be saved
	savePath := "./pkg/api/" + *name + "_recover.go"

	recoverySource := fmt.Sprintf(recoveryHandlerTemplate, *name, *name)

	fmt.Println("Saving file to : ", savePath)
	err := ioutil.WriteFile(savePath, []byte(recoverySource), 0700)

	if err != nil {
		panic(err)
	}

	fmt.Println("Add this to your server's recover directive : do ", *name+";")

}
