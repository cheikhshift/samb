package transpiler

import (
	"errors"
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

func VerifyProviders(p *samb.Project, providers []string) (err error) {

	for _, provider := range providers {
		if !p.HasProvider(provider) {
			err = errors.New("Provider " + provider + " not found!")
			break
		}
	}

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

func GetCusomCode(r samb.Route) string {
	return strings.Join(r.Go.Do, "\n") + "\n"
}

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
	gocode := GetCusomCode(r)

	handler = providerInitializer + gocode + r.Handler

	return
}

func GetProviderInits(p *samb.Project, providers []string) (res string) {
	for _, providerId := range providers {
		if providerId == "r" || providerId == "w" {
			continue
		}

		provider := p.GetProvider(providerId)
		res += fmt.Sprintf("\nvar %s = %s\n", provider.Name, provider.Return)
	}

	return
}
