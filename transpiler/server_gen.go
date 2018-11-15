package transpiler

import (
	"fmt"
	"io/ioutil"

	"github.com/cheikhshift/samb"

	"strings"
)

func ExportServer(p *samb.Project) error {

	startCode := strings.Join(p.Server.Start.Do, "\n")

	startCode = fmt.Sprintf(cmdWrapper, "Start", startCode)

	err := ioutil.WriteFile("./cmd/server/launch.go", []byte(startCode), 0700)

	if err != nil {
		return err
	}

	return ExportExitCode(p)
}

func ExportExitCode(p *samb.Project) error {
	shutdownCode := strings.Join(p.Server.Shutdown.Do, "\n")

	shutdownCode = fmt.Sprintf(cmdWrapper, "Stop", shutdownCode)

	err := ioutil.WriteFile("./cmd/server/stop.go", []byte(shutdownCode), 0700)

	if err != nil {
		return err
	}

	return ExportMain(p)
}

func ExportMain(p *samb.Project) error {

	mainCode := fmt.Sprintf(mainWrapper, p.Package)

	err := ioutil.WriteFile("./cmd/server/main.go", []byte(mainCode), 0700)

	if err != nil {
		return err
	}

	return nil
}

func ExportConfig(s samb.Server) error {

	serverConfig := fmt.Sprintf(configWrapper, s.Port, s.Host, s.Webroot)

	err := ioutil.WriteFile("./cmd/server/config.go", []byte(serverConfig), 0700)

	if err != nil {
		return err
	}

	return nil

}
