package transpiler

import (
	"fmt"
	"io/ioutil"

	"github.com/cheikhshift/samb"

	"strings"
)

// ExportServer converts your server's
// start directive to a function, executed on
// server start. The file is saved to ./cmd/server/launch.go
func ExportServer(p *samb.Project) error {

	startCode := strings.Join(p.Server.Start.Do, "\n")

	startCode = fmt.Sprintf(cmdWrapper, strings.Join(p.Import, "\n"), "Start", startCode)

	err := ioutil.WriteFile("./cmd/server/launch.go", []byte(startCode), 0700)

	if err != nil {
		return err
	}

	return ExportExitCode(p)
}

// ExportExitCode converts your server's
// shutdown directive to a function, executed on
// server termination. The file is saved to
// ./cmd/server/stop.go
func ExportExitCode(p *samb.Project) error {
	shutdownCode := strings.Join(p.Server.Shutdown.Do, "\n")

	shutdownCode = fmt.Sprintf(cmdWrapper, strings.Join(p.Import, "\n"), "Stop", shutdownCode)

	err := ioutil.WriteFile("./cmd/server/stop.go", []byte(shutdownCode), 0700)

	if err != nil {
		return err
	}

	return ExportMain(p)
}

// ExportMain writes the files needed
// to launch your server. The file is saved
// to ./cmd/server/main.go
func ExportMain(p *samb.Project) error {

	mainCode := fmt.Sprintf(mainWrapper, p.Package)

	err := ioutil.WriteFile("./cmd/server/main.go", []byte(mainCode), 0700)

	if err != nil {
		return err
	}

	return nil
}

// ExportConfig writes the port and host specified
// by your server to ./cmd/server/config.go as Go code.
// The variables written will have the same identifier name
// as their respective directive name.
func ExportConfig(s samb.Server) error {

	serverConfig := fmt.Sprintf(configWrapper, s.Port, s.Host, s.Webroot)

	err := ioutil.WriteFile("./cmd/server/config.go", []byte(serverConfig), 0700)

	if err != nil {
		return err
	}

	return nil

}
