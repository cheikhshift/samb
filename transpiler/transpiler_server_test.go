package transpiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cheikhshift/samb"
)

func TestTranspile(t *testing.T) {

	for i, p := range testProjects {
		t.Run(string(i), func(t *testing.T) {

			setupTestEnv()
			defer teardownTestEnv()

			makePkgPaths()

			err := Transpile(p)

			if err != nil {
				panic(err)
			}

			for _, path := range testExpectedFiles {

				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Errorf("File %v was not found", path)
				}

			}
		})
	}
}

func TestExportGlobals(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./pkg/globals/", 0700)

	if err != nil {
		panic(err)
	}

	testGlobals := []struct {
		globals        []samb.Global
		expectedOutput string
	}{
		{
			[]samb.Global{
				{Name: "Foo", Return: "string(\"Foo\")"},
			},
			`// Package globals has your applications
// global variables, exported as package identifiers
// GENERATED CODE, DO NOT EDIT!
package globals


			// 
			var Foo  = string("Foo")
		`,
		},
		{
			[]samb.Global{
				{Name: "Foo", Return: "string(\"Foo\")"},
				{Name: "Bax", Comment: "Sample", Return: "string(\"Bax\")"},
			},
			`// Package globals has your applications
// global variables, exported as package identifiers
// GENERATED CODE, DO NOT EDIT!
package globals


			// 
			var Foo  = string("Foo")
		
			// Sample
			var Bax  = string("Bax")
		`,
		},
		{
			[]samb.Global{
				{Name: "Zo", Return: "10"},
				{Name: "Go", Comment: "Go", Return: "string(\"Go\")"},
			},
			`// Package globals has your applications
// global variables, exported as package identifiers
// GENERATED CODE, DO NOT EDIT!
package globals


			// 
			var Zo  = 10
		
			// Go
			var Go  = string("Go")
		`,
		},
	}

	for i, tt := range testGlobals {
		t.Run(string(i), func(t *testing.T) {

			var generatedFile string

			err := ExportGlobals(tt.globals)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./pkg/globals/variables.go")

			generatedFile = string(fileBytes)

			if generatedFile != tt.expectedOutput {
				t.Errorf("Got %v, expected %v", generatedFile, tt.expectedOutput)
			}

		})
	}
}

func TestExportServer(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./cmd/server/", 0700)

	if err != nil {
		panic(err)
	}

	exportStartTests := []struct {
		start          samb.Go
		expectedOutput string
	}{
		{
			samb.Go{[]string{"println('Hello')"}},
			fmt.Sprintf(cmdWrapper, "", "Start", "println('Hello')"),
		},
		{
			samb.Go{[]string{"println('HelloThree')"}},
			fmt.Sprintf(cmdWrapper, "", "Start", "println('HelloThree')"),
		},
		{
			samb.Go{[]string{"fmt.Println(\"Foo\")", "os.Exit(0)"}},
			fmt.Sprintf(cmdWrapper, "", "Start", "fmt.Println(\"Foo\")\nos.Exit(0)"),
		},
		{
			samb.Go{[]string{"virtualPackage.Exec(100)"}},
			fmt.Sprintf(cmdWrapper, "", "Start", "virtualPackage.Exec(100)"),
		},
	}

	for i, tt := range exportStartTests {
		t.Run(string(i), func(t *testing.T) {

			var generatedFile string

			err := ExportServer(
				&samb.Project{
					Server: samb.Server{Start: tt.start},
				},
			)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./cmd/server/launch.go")

			generatedFile = string(fileBytes)

			if generatedFile != tt.expectedOutput {
				t.Errorf("Got %v, expected %v", generatedFile, tt.expectedOutput)
			}

		})
	}
}

func TestExportExitCode(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./cmd/server/", 0700)

	if err != nil {
		panic(err)
	}

	exportExitTests := []struct {
		exit           samb.Go
		expectedOutput string
	}{
		{
			samb.Go{[]string{"println('Bye')"}},
			fmt.Sprintf(cmdWrapper, "", "Stop", "println('Bye')"),
		},
		{
			samb.Go{[]string{"println('HelloThree')"}},
			fmt.Sprintf(cmdWrapper, "", "Stop", "println('HelloThree')"),
		},
		{
			samb.Go{[]string{"fmt.Println(\"Foo\")", "os.Exit(0)"}},
			fmt.Sprintf(cmdWrapper, "", "Stop", "fmt.Println(\"Foo\")\nos.Exit(0)"),
		},
		{
			samb.Go{[]string{"virtualPackage.Exec(100)"}},
			fmt.Sprintf(cmdWrapper, "", "Stop", "virtualPackage.Exec(100)"),
		},
	}

	for i, tt := range exportExitTests {
		t.Run(string(i), func(t *testing.T) {

			var generatedFile string

			err := ExportServer(
				&samb.Project{
					Server: samb.Server{Shutdown: tt.exit},
				},
			)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./cmd/server/stop.go")

			generatedFile = string(fileBytes)

			if generatedFile != tt.expectedOutput {
				t.Errorf("Got %v, expected %v", generatedFile, tt.expectedOutput)
			}

		})
	}

}

func TestExportMain(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./cmd/server/", 0700)

	if err != nil {
		panic(err)
	}

	exportExitTests := []struct {
		pkg            string
		expectedOutput string
	}{
		{
			"virtual/pkg/v6/sample",
			fmt.Sprintf(mainWrapper, "virtual/pkg/v6/sample"),
		},
		{
			"virtual/v6/sample",
			fmt.Sprintf(mainWrapper, "virtual/v6/sample"),
		},
		{
			"virtual/v9/sample",
			fmt.Sprintf(mainWrapper, "virtual/v9/sample"),
		},
		{
			"github.com/cheikhshift/samb",
			fmt.Sprintf(mainWrapper, "github.com/cheikhshift/samb"),
		},
	}

	for i, tt := range exportExitTests {
		t.Run(string(i), func(t *testing.T) {

			var generatedFile string

			err := ExportMain(
				&samb.Project{
					Package: tt.pkg,
				},
			)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./cmd/server/main.go")

			generatedFile = string(fileBytes)

			if generatedFile != tt.expectedOutput {
				t.Errorf("Got %v, expected %v", generatedFile, tt.expectedOutput)
			}

		})
	}

}

func TestExportConfig(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./cmd/server/", 0700)

	if err != nil {
		panic(err)
	}

	exportExitTests := []struct {
		port           string
		host           string
		expectedOutput string
	}{
		{
			"455",
			"gophersauce.com",
			fmt.Sprintf(configWrapper, "455", "gophersauce.com", ""),
		},
		{
			"755",
			"go.com",
			fmt.Sprintf(configWrapper, "755", "go.com", ""),
		},
		{
			"1455",
			"sample.com",
			fmt.Sprintf(configWrapper, "1455", "sample.com", ""),
		},
		{
			"200",
			"example.com",
			fmt.Sprintf(configWrapper, "200", "example.com", ""),
		},
	}

	for i, tt := range exportExitTests {
		t.Run(string(i), func(t *testing.T) {

			var generatedFile string

			err := ExportConfig(
				samb.Server{
					Port: tt.port, Host: tt.host,
				},
			)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./cmd/server/config.go")

			generatedFile = string(fileBytes)

			if generatedFile != tt.expectedOutput {
				t.Errorf("Got %v, expected %v", generatedFile, tt.expectedOutput)
			}

		})
	}

}
