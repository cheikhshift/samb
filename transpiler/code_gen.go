package transpiler

import (
	"fmt"
	"io/ioutil"

	"github.com/cheikhshift/samb"
)

// Transpile generates Go source from
// your samb probject. The files are
// saved relative to your program's
// working directory... ie : ./pkg/api..
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

// ExportGlobals saves an array of samb globals
// to path ./pkg/globals/variables.go as go code.
// if write was not successful an error will be
// returned.
func ExportGlobals(globals []samb.Global) error {

	var globalStr string

	for _, g := range globals {

		globalStr += fmt.Sprintf(`
			// %s
			var %s %s = %s
		`, g.Comment, g.Name, g.Type, g.Return)
	}

	globalStr = fmt.Sprintf(globalWrapper, globalStr)

	return ioutil.WriteFile(
		"./pkg/globals/variables.go",
		[]byte(globalStr),
		0700,
	)
}
