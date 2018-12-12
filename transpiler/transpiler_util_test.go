package transpiler

import (
	"testing"

	"github.com/cheikhshift/samb"
)

func TestVerifyProviders(t *testing.T) {

	providerTests := []struct {
		name string
	}{
		{"Foo"},
		{"Bar"},
		{"Baz"},
		{"Z"},
		{"OO"},
		{"MM"},
	}

	for _, tt := range providerTests {
		t.Run(tt.name, func(t *testing.T) {

			err := VerifyProviders(testProject, []string{tt.name})

			if err != nil {
				t.Errorf("got  %v, want %v for provider "+tt.name, err, nil)
			}

		})
	}

}

func TestGetProviderInits(t *testing.T) {
	providerTests := []struct {
		name          string
		expectedValue string
	}{
		{"Foo", "\n//\nvar Foo = string(\"Foo\")\n"},
		{"Bar", "\n//\nvar Bar = string(\"Bar\")\n"},
		{"Baz", "\n//\nvar Baz = string(\"Baz\")\n"},
		{"Z", "\n//\nvar Z = string(\"Z\")\n"},
	}

	for _, tt := range providerTests {
		t.Run(tt.name, func(t *testing.T) {

			p := GetProviderInits(testProject, []string{tt.name}, "")

			if p != tt.expectedValue {
				t.Errorf("got  %v, want %v for provider "+tt.name, p, tt.expectedValue)
			}

		})
	}
}

func TestGetCustomCode(t *testing.T) {

	codeGetTest := []struct {
		route         samb.Route
		expectedValue string
	}{
		{
			samb.Route{Go: samb.Go{[]string{"println(\"Foo\")", "pkg.Function(12)"}}},
			"println(\"Foo\")\npkg.Function(12)\n",
		},
		{
			samb.Route{Go: samb.Go{[]string{"println(\"Baz\")", "pkg.Fn(1)"}}},
			"println(\"Baz\")\npkg.Fn(1)\n",
		},
		{
			samb.Route{Go: samb.Go{[]string{"i := 0", "return"}}},
			"i := 0\nreturn\n",
		},
	}

	for i, tt := range codeGetTest {
		t.Run(string(i), func(t *testing.T) {

			p := GetCustomCode(tt.route)

			if p != tt.expectedValue {
				t.Errorf("got  %v, want %v", p, tt.expectedValue)
			}

		})
	}
}

func TestWrapEndpoint(t *testing.T) {

	expectedValues := []string{
		`
		if  basePath := "/Hello"; strings.Contains(r.URL.Path , basePath) && r.Method == "POST"{
		
	}`,
		`
		if  basePath := "/echo"; strings.Contains(r.URL.Path , basePath) && r.Method == "GET"{
		
	}`,
		`
		if  basePath := "/with_provider"; strings.Contains(r.URL.Path , basePath) && r.Method == "PUT"{
		
	}`,
		`
		if  basePath := "/object/path/res"; strings.Contains(r.URL.Path , basePath) && r.Method == "DELETE"{
		
	}`,
		`
		if  basePath := "/*"; strings.Contains(r.URL.Path , basePath) {
		
	}`,
		`
		if  basePath := "/baz_path"; strings.Contains(r.URL.Path , basePath) {
		
	}`,
	}

	for i, tt := range routeTests {
		t.Run(string(i), func(t *testing.T) {

			result := WrapEndpoint("", tt.route, "")

			if result != expectedValues[i] {
				println(result)
				println("///")
				t.Errorf("got  %v, want %v", result, expectedValues[i])
			}

		})
	}
}

func TestGetHandler(t *testing.T) {

	blankString := `
`

	expectedValues := []string{
		blankString,
		blankString,
		`
//
var Foo = string("Foo")

`,
		blankString,
		blankString,
		blankString,
	}

	for i, tt := range routeTests {
		t.Run(string(i), func(t *testing.T) {

			result := GetHandler(testProject, tt.route, tt.route.Provide)

			if result != expectedValues[i] {
				println(result)
				println("///")
				t.Errorf("got  %v, want %v", result, expectedValues[i])
			}

		})
	}
}
