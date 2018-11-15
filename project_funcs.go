package samb

import (
	"os"
)

// Adds required code to project
func (p *Project) ProcessImports() {

	if p.Server.Require != nil {
		p.Require = append(p.Require, p.Server.Require...)
	}

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	for _, i := range p.Require {

		os.Chdir(wd)

		file, err := Load(i)

		if err != nil {
			panic(err)
		}

		p.MergeWith(file)

	}

}

func (p *Project) MergeWith(file *Project) {

	p.Provider = append(p.Provider, file.Provider...)
	p.Global = append(p.Global, file.Global...)
	p.Packages = append(p.Packages, file.Packages...)
	p.Import = append(p.Import, file.Import...)

	// process file route imports
	p.Templates.Template = append(p.Templates.Template, file.Templates.Template...)
	p.Routes.Route = append(p.Routes.Route, file.Routes.Route...)

}
