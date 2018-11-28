package main

import (
	"flag"
	"os"
	"strings"
)

var providerFile = flag.String(
	"file",
	"providers.se",
	"Files with provider directives. This new provider will be appended to it.",
)

func main() {

	name := flag.String("name", "FooBar", "Name of provider to add.")
	project := flag.String("project", "./", "Path to project")

	t := flag.String("type", "string", "Type of variable to provide")

	flag.Parse()

	os.Chdir(*project)

	providers, err := loadFile(*providerFile)

	if err != nil {
		panic(err)
	}

	if !strings.Contains(*t, "*") {
		typeWithAstrk := "*" + *t
		t = &typeWithAstrk
	}

	addProviderDirective(providers, *name, *t)

	saveProviderFunc(*name, *t)

}
