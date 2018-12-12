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
	"./pkg/hooks/launch.go",
	"./pkg/hooks/stop.go",
}

func setupHookEnv(){
	err := os.MkdirAll("./pkg/hooks/", 0700)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./cmd/server/", 0700)

	if err != nil {
		panic(err)
	}
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
		{Name: "OO", Type: "string", Return: "string(\"OO\")"},
		{Name: "MM", Type: "string", Return: "string(\"MM\")"},
	},
}

var testProjects = []*samb.Project{
	{
		Provider: []samb.Global{
			{Name: "Foo", Type: "int", Return: "string(\"Foo\")"},
			{Name: "Baz", Type: "string", Return: "string(\"Baz\")"},
			{Name: "Z", Type: "string", Return: "string(\"Z\")"},
			{Name: "Bar", Type: "string", Return: "string(\"Bar\")"},
		},
	},
	{
		Global: []samb.Global{
			{Name: "Foo", Type: "int", Return: "string(\"Foo\")"},
		},
	},
	{
		Package: "sample/test",
	},
	{
		Routes: samb.Routes{
			Route: []samb.Route{
				{Path: "Sample", Method: "POST"},
				{Path: "/hello", Method: "PUT"},
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
		if  basePath := "/Hello"; strings.Contains(r.URL.Path , basePath) && r.Method == "POST"{
		

	}`,
	},
	{
		samb.Route{Method: "GET", Path: "/echo"},
		`
		if  basePath := "/echo"; strings.Contains(r.URL.Path , basePath) && r.Method == "GET"{
		

	}`,
	},
	{
		samb.Route{Method: "PUT", Provide: []string{"Foo"}, Path: "/with_provider"},
		`
		if  basePath := "/with_provider"; strings.Contains(r.URL.Path , basePath) && r.Method == "PUT"{
		
//
var Foo = string("Foo")


	}`,
	},
	{
		samb.Route{Method: "DELETE", Path: "/object/path/res"},
		`
		if  basePath := "/object/path/res"; strings.Contains(r.URL.Path , basePath) && r.Method == "DELETE"{
		

	}`,
	},
	{
		samb.Route{Method: "*", Path: "/*"},
		`
		if  basePath := "/*"; strings.Contains(r.URL.Path , basePath) {
		

	}`,
	},
	{
		samb.Route{Method: "*", Path: "/baz_path"},
		`
		if  basePath := "/baz_path"; strings.Contains(r.URL.Path , basePath) {
		

	}`,
	},
}
