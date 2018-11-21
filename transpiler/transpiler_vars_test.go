package transpiler

import (
	"os"

	"github.com/cheikhshift/samb"
)

var testExpectedPkgFolders = []string{"./pkg/api", "./pkg/methods", "./pkg/globals", "./cmd/server"}
var testExpectedFiles = []string{
	"./pkg/globals/variables.go",
	"./pkg/api/handler.go",
	"./cmd/server/main.go",
	"./cmd/server/launch.go",
	"./cmd/server/stop.go",
}

func setupTestEnv() {
	err := os.Mkdir("./test_temp", 0700)

	if err != nil {
		panic(err)
	}

	err = os.Chdir("./test_temp")

	if err != nil {
		panic(err)
	}
}

func teardownTestEnv() {

	err := os.Chdir("../")

	if err != nil {
		panic(err)
	}

	err = os.RemoveAll("./test_temp")

	if err != nil {
		panic(err)
	}
}

var testProject = &samb.Project{
	Provider: []samb.Global{
		{Name: "Foo", Type: "int", Return: "string(\"Foo\")"},
		{Name: "Baz", Type: "string", Return: "string(\"Baz\")"},
		{Name: "Z", Type: "string", Return: "string(\"Z\")"},
		{Name: "Bar", Type: "string", Return: "string(\"Bar\")"},
	},
}

var testProjects = []*samb.Project{
	&samb.Project{
		Provider: []samb.Global{
			{Name: "Foo", Type: "int", Return: "string(\"Foo\")"},
			{Name: "Baz", Type: "string", Return: "string(\"Baz\")"},
			{Name: "Z", Type: "string", Return: "string(\"Z\")"},
			{Name: "Bar", Type: "string", Return: "string(\"Bar\")"},
		},
	},
	&samb.Project{
		Global: []samb.Global{
			{Name: "Foo", Type: "int", Return: "string(\"Foo\")"},
		},
	},
	&samb.Project{
		Package: "sample/test",
	},
	&samb.Project{
		Routes: samb.Routes{
			Route: []samb.Route{
				samb.Route{Path: "Sample", Method: "POST"},
				samb.Route{Path: "/hello", Method: "PUT"},
			},
		},
	},
}

var routeTests = []struct {
	route         samb.Route
	expectedRoute string
}{
	{
		samb.Route{Method: "POST", Path: "/Hello"},
		`
		if  strings.Contains(r.URL.Path , "/Hello") && r.Method == "POST"{
		

	}`,
	},
	{
		samb.Route{Method: "GET", Path: "/echo"},
		`
		if  strings.Contains(r.URL.Path , "/echo") && r.Method == "GET"{
		

	}`,
	},
	{
		samb.Route{Method: "PUT", Provide: []string{"Foo"}, Path: "/with_provider"},
		`
		if  strings.Contains(r.URL.Path , "/with_provider") && r.Method == "PUT"{
		
//
var Foo = string("Foo")


	}`,
	},
	{
		samb.Route{Method: "DELETE", Path: "/object/path/res"},
		`
		if  strings.Contains(r.URL.Path , "/object/path/res") && r.Method == "DELETE"{
		

	}`,
	},
	{
		samb.Route{Method: "*", Path: "/*"},
		`
		if  strings.Contains(r.URL.Path , "/*") {
		

	}`,
	},
	{
		samb.Route{Method: "*", Path: "/baz_path"},
		`
		if  strings.Contains(r.URL.Path , "/baz_path") {
		

	}`,
	},
}
