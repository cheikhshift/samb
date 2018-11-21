package transpiler

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/cheikhshift/samb"
)

func TestSetup(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	Setup()

	for _, path := range testExpectedPkgFolders {
		t.Run(path, func(t *testing.T) {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				t.Errorf("Path %v was not found", path)
			}
		})
	}
}

func TestMakePkgPaths(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	makePkgPaths()

	for _, path := range testExpectedPkgFolders {
		t.Run(path, func(t *testing.T) {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				t.Errorf("Path %v was not found", path)
			}
		})
	}
}

func TestProcessRoute(t *testing.T) {

	for _, tt := range routeTests {
		t.Run(tt.route.Path, func(t *testing.T) {

			route, err := ProcessRoute(&samb.Project{Provider: testProject.Provider}, tt.route, "", []string{})

			if err != nil {
				panic(err)
			}

			if route != tt.expectedRoute {
				t.Errorf("got  %v, want %v for route [%v] "+tt.route.Path, route, tt.expectedRoute, tt.route.Method)
			}

		})
	}

}

func TestExportRoutes(t *testing.T) {

	setupTestEnv()
	defer teardownTestEnv()

	err := os.MkdirAll("./pkg/api/", 0700)

	if err != nil {
		panic(err)
	}

	expectedHandlers := []string{
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/Hello") && r.Method == "POST"{
		

	}
}`,
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/echo") && r.Method == "GET"{
		

	}
}`,
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/with_provider") && r.Method == "PUT"{
		
//
var Foo = string("Foo")


	}
}`,
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/object/path/res") && r.Method == "DELETE"{
		

	}
}`,
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/*") {
		

	}
}`,
		`//package api contains your web app's handler definitions.
		// GENERATED CODE, DO NOT EDIT!!
		package api

		import (
		
	)

		// Handles routing of application.
		func Handler(w http.ResponseWriter, r *http.Request) {

			defer catchPanic(w,r)

			
		if  strings.Contains(r.URL.Path , "/baz_path") {
		

	}
}`,
	}

	for i, tt := range routeTests {
		t.Run(tt.route.Path, func(t *testing.T) {

			testProject := &samb.Project{
				Routes: samb.Routes{
					Route: []samb.Route{tt.route},
				},
				Provider: testProject.Provider,
			}

			err := ExportRoutes(testProject)

			if err != nil {
				panic(err)
			}

			fileBytes, err := ioutil.ReadFile("./pkg/api/handler.go")

			generatedFile := string(fileBytes)

			if generatedFile != expectedHandlers[i] {
				println(generatedFile)
				println("///")
				t.Errorf("got  %v, want %v", generatedFile, expectedHandlers[i])
			}

		})
	}

}
