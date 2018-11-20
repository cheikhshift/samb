package transpiler

import (
	"fmt"
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

	return fmt.Sprintf(`%s{
		%s
	}`, res, h)
}

// GetHandler Generates the Go code executed specified to
// be executed by
// a samb route.
func GetHandler(p *samb.Project, r samb.Route, providers []string) (handler string) {

	var providerInitializer = GetProviderInits(p, providers)
	gocode := GetCustomCode(r)

	handler = providerInitializer + gocode + r.Handler

	return
}

// StripSlashes turns ever letter following
// a forward slash to uppercase, then removes slashes from string,
func StripSlashes(i string) string {
	var tlc = fmt.Sprintf(
		"%s", i, strings.Replace(
			strings.Title(strings.Replace(i, "/", " ", -1)), " ", "", -1,
		),
	)

	return tlc
}

// GetCustomCode returns the Go
// code to be ran, by the specified
// samb route.
func GetCustomCode(r samb.Route) string {
	return strings.Join(r.Go.Do, "\n") + "\n"
}
