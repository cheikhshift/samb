package samb

// HasProvider checks to see if the project
// has a provider with the specified name.
func (p *Project) HasProvider(name string) (has bool) {

	if name == "r" || name == "w" {
		return true
	}

	for _, pr := range p.Provider {
		if pr.Name == name {
			has = true
			break
		}
	}

	return
}

// GetProvider gets a provider from a project,
// based on the name passed. It is recommended
// to check for a provider with local function
// HasProvider prior to getting it.
func (p *Project) GetProvider(name string) (pr Global) {

	for _, pr = range p.Provider {
		if pr.Name == name {
			break
		}
	}

	return
}
