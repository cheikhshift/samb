package transpiler

import (
	"fmt"
	"io/ioutil"

	"github.com/cheikhshift/samb"

	"strings"
)

func ExportRoutes(p *samb.Project) (err error) {

	var routeDef, def string

	for _, r := range p.Routes.Route {
		def, err = ProcessRoute(p, r, "", p.Routes.Provide)

		if err != nil {
			break
		}

		routeDef += def
	}

	routeDef = fmt.Sprintf(routeWrapper, strings.Join(p.Import, "\n"), routeDef)

	err = ioutil.WriteFile("./pkg/api/handler.go", []byte(routeDef), 0700)

	return
}

func ProcessRoute(p *samb.Project, r samb.Route, path string, providers []string) (def string, err error) {

	providers = append(providers, r.Provide...)

	if len(r.Route) > 0 {

		newRoot := path + r.Path

		var defChild string

		for _, route := range r.Route {
			defChild, err = ProcessRoute(p, route, newRoot, providers)
			if err != nil {
				break
			}

			def += defChild
		}

		if err != nil {
			return
		}

	}

	err = VerifyProviders(p, providers)
	if err != nil {
		return
	}

	h := GetHandler(p, r, providers)

	def = WrapEndpoint(path, r, h+def)

	return
}
