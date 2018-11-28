package transpiler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cheikhshift/samb"
)

// GetProviderInits generates the Go
// source required to initialize a provider variable.
// providers with name's r and w are skipped.
// The project will have the information about the provider,
// and providers is a list of name's of providers to generate
// code for.
func GetProviderInits(p *samb.Project, providers []string, endPointCode string) (res string) {
	for _, providerId := range providers {

		if providerId == "r" || providerId == "w" || 
		(!strings.Contains(endPointCode, providerId) && len(endPointCode) > 2 ) {
			continue
		}

		provider := p.GetProvider(providerId)
		res += fmt.Sprintf("\n//%s\nvar %s = %s\n", provider.Comment, provider.Name, provider.Return)
	}

	return
}

// VerifyProviders checks to see if a project
// has the providers specified by `providers` parameter.
func VerifyProviders(p *samb.Project, providers []string) (err error) {

	for _, provider := range providers {
		if !p.HasProvider(provider) {
			err = errors.New("Provider " + provider + " not found!")
			break
		}
	}

	return
}
