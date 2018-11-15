package samb

// Adds required code to project
func (p *Project) ProcessImports() {

	for _, i := range p.Require {

		file, err := Load(i)

		if err != nil {
			panic(err)
		}

		p.Provider = append(p.Provider, file.Provider...)
		p.Global = append(p.Global, file.Global...)
		p.Packages = append(p.Packages, file.Packages...)
		p.Import = append(p.Import, file.Import...)

		// process file route imports
		p.Templates.Template = append(p.Templates.Template, file.Templates.Template...)
		p.Routes.Route = append(p.Routes.Route, file.Routes.Route...)

	}

}

func (p *Project) ProcessServerImports() {
	for _, i := range p.Server.Require {

		file, err := Load(i)

		if err != nil {
			panic(err)
		}

		p.Templates.Template = append(p.Templates.Template, file.Templates.Template...)
		p.Routes.Route = append(p.Routes.Route, file.Routes.Route...)
		p.Import = append(p.Import, file.Import...)

	}
}
