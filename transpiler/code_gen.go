package transpiler

import (
	"fmt"
	"io/ioutil"

	"github.com/cheikhshift/samb"
)

func Transpile(p *samb.Project) error {

	err := ExportGlobals(p.Global)

	if err != nil {
		return err
	}

	err = ExportConfig(p.Server)

	if err != nil {
		return err
	}

	err = ExportRoutes(p)

	if err != nil {
		return err
	}

	return ExportServer(p)
}

func ExportGlobals(globals []samb.Global) error {

	var globalStr string

	for _, g := range globals {

		globalStr += fmt.Sprintf(`
			var %s %s = %s
		`, g.Name, g.Type, g.Return)
	}

	globalStr = fmt.Sprintf(globalWrapper, globalStr)

	return ioutil.WriteFile(
		"./pkg/globals/variables.go",
		[]byte(globalStr),
		0700,
	)
}
