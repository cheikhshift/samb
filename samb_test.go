package samb

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {

	var flagtests = []struct {
		file       []byte
		parsedPort string
	}{
		{testFileOne, "5555"},
		{testFileTwo, "5556"},
		{testFileThree, "1556"},
		{testFileFour, "556"},
	}

	for i, tt := range flagtests {
		t.Run(string(i), func(t *testing.T) {

			err := ioutil.WriteFile("./test_temp.se", tt.file, 0700)

			if err != nil {
				panic(err)
			}

			project, err := Load("./test_temp.se")

			if err != nil {
				panic(err)
			}

			os.Remove("./test_temp.se")

			if project.Server.Port != tt.parsedPort {
				t.Errorf("got %q, want %q", project.Server.Port, tt.parsedPort)
			}
		})
	}

}

func TestProcessImports(t *testing.T) {

	expectedRouteLength := 2

	err := ioutil.WriteFile("./test_temp.se", testFileFour, 0700)

	if err != nil {
		panic(err)
	}

	testProject := &Project{
		Require: []string{"./test_temp.se"},
	}

	testProject.ProcessImports()

	os.Remove("./test_temp.se")

	if expectedRouteLength != len(testProject.Routes.Route) {
		t.Errorf("got %q, want %q", len(testProject.Routes.Route), expectedRouteLength)
	}

}


func TestMergeWith(t *testing.T){

	var mergeTests = []struct {
		expectedImportLength int
		expectedRouteLength int
		project *Project
	}{
			{
				3, 0 , &Project{
				Import : []string{"Foo", "Baz",""},
				},
			},
			{
				2, 1 , &Project{
				Import : []string{"Foo", "Baz"},
				Routes : Routes{
						Route : []Route{
							Route{Path : "Sample"},
						},
					},
				},
			},
			{
				4, 2 , &Project{
				Import : []string{"Foo", "Baz", "Extra", "TEst"},
				Routes : Routes{
						Route : []Route{
							Route{Path : "Sample"},
							Route{Path : "SampleThree"},
						},
					},
				},
			},
	}

	for i, tt := range mergeTests {
		t.Run(string(i), func(t *testing.T) {

			tempProject := &Project{}

			tempProject.MergeWith(tt.project)

			if len(tempProject.Import) != tt.expectedImportLength  {
				t.Errorf("got %v, want %v for number of imports", len(tempProject.Import), tt.expectedImportLength)
			}

			if len(tempProject.Routes.Route) != tt.expectedRouteLength  {
				t.Errorf("got %v, want %v for number of routes", len(tempProject.Routes.Route), tt.expectedRouteLength)
			}
		})
	}

}

func TestHasProvider(t *testing.T){

	testProject := &Project{
		Provider : []Global{
			Global{Name:"Foo"},
			Global{Name:"Baz"},
			Global{Name:"Go"},
			Global{Name:"Vn"},
			Global{Name:"Bar"},
		},
	}

	// name of providers to look for.
	providerTests := []string{
		"Foo",
		"Baz",
		"Go",
		"Vn",
		"Bar",
	}

	for _, name := range providerTests {
		t.Run(name, func(t *testing.T) {

	
			has := testProject.HasProvider(name)

			if  !has  {
				t.Errorf("got %v, want %v for provider " + name, has, true)
			}

			
		})
	}

}

func TestGetProvider(t *testing.T){

	testProject := &Project{
		Provider : []Global{
			Global{Name:"Foo", Type:"int"},
			Global{Name:"Baz", Type:"string"},
			Global{Name:"Go", Type:"int64"},
			Global{Name:"Vn", Type: "int" },
			Global{Name:"Bar", Type: "string"},
		},
	}

	// name of providers to look for.
	expectedProviderTypes := []struct{
		expectedType string
		name string
	}{
		{"int", "Foo" },
		{"string", "Baz"},
		{"int64", "Go"},
		{"int", "Vn"},
		{"string", "Bar"},
	}

	for _, tt := range expectedProviderTypes {
		t.Run(tt.name, func(t *testing.T) {

	
			p := testProject.GetProvider(tt.name)

			if  p.Type != tt.expectedType  {
				t.Errorf("got type %v, want %v for provider " + p.Name, p.Type , tt.expectedType )
			}

			
		})
	}

}
