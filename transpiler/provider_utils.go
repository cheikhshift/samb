package transpiler

import (
	"errors"
	"fmt"

	"github.com/cheikhshift/samb"
)

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

func VerifyProviders(p *samb.Project, providers []string) (err error) {

	for _, provider := range providers {
		if !p.HasProvider(provider) {
			err = errors.New("Provider " + provider + " not found!")
			break
		}
	}

	return
}
