package transpiler

import (
	"fmt"
	"strings"

	"github.com/cheikhshift/samb"
)

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

func GetHandler(p *samb.Project, r samb.Route, providers []string) (handler string) {

	var providerInitializer = GetProviderInits(p, providers)
	gocode := GetCustomCode(r)

	handler = providerInitializer + gocode + r.Handler

	return
}

func StripSlashes(i string) string {
	var tlc = fmt.Sprintf(
		"%s", i, strings.Replace(
			strings.Title(strings.Replace(i, "/", " ", -1)), " ", "", -1,
		),
	)

	return tlc
}

func GetCustomCode(r samb.Route) string {
	return strings.Join(r.Go.Do, "\n") + "\n"
}
