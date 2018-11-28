package transpiler

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cheikhshift/samb"
)

// WrapEndpoint adds conditional Go statements
// around the specified h parameter. Parameter
// h is Go code. parameter r (samb.Route) specifies
// which condition is evaluated.
func WrapEndpoint(path string, r samb.Route, h string) string {
	res := fmt.Sprintf(`
		if  strings.Contains(r.URL.Path , "%s") `, path+r.Path)

	if r.Method != "" &&
		r.Method != "*" {

		res += fmt.Sprintf(`&& r.Method == "%s"`, strings.ToUpper(r.Method))

	}

	return fmt.Sprintf(endpointWrapper, res, h)
}

// GetHandler Generates the Go code executed specified to
// be executed by
// a samb route.
func GetHandler(p *samb.Project, r samb.Route, providers []string) (handler string) {

	gocode := GetCustomCode(r)
	endPointCode := gocode + r.Handler
	var providerInitializer = GetProviderInits(p, providers, endPointCode)

	handler = providerInitializer + endPointCode
	
	if r.Handler != "" {
		handler += "\nreturn"
	}

	return
}

// GetCustomCode returns the Go
// code to be ran, by the specified
// samb route.
func GetCustomCode(r samb.Route) string {
	return strings.Join(r.Go.Do, "\n") + "\n"
}

// WriteRecoveryFuncs generates a file
// with the recovery function called
// on panic.
func WriteRecoveryFuncs(r []string) error {

	for i := range r {
		r[i] += "(w,r,n.(string))"
	}

	recovFile := fmt.Sprintf(recoveryWrapper, strings.Join(r, "\n"))

	err := ioutil.WriteFile("./pkg/api/recover.go", []byte(recovFile), 0700)

	return err

}
