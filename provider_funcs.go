package samb

func (p *Project) HasProvider(name string) (has bool) {

	for _, pr := range p.Provider {
		if pr.Name == name {
			has = true
			break
		}
	}

	return
}

func (p *Project) GetProvider(name string) (pr Global) {

	for _, pr = range p.Provider {
		if pr.Name == name {
			break
		}
	}

	return
}
