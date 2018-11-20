package samb

import (
	"log"
	"os"
	"strings"

	"github.com/recoye/config"
)

// Load, open a project based on the filesystem
// path supplied.
func Load(path string) (*Project, error) {

	conf := config.New(path)

	env := &Project{}

	err := conf.Unmarshal(env)
	if err != nil {
		return nil, err
	}

	chDirRootOf(path)

	env.ProcessImports()

	return env, nil
}

// chDirRootOf will change the working directory to the directory
// of the specified SAMB
// source file path.
func chDirRootOf(file string) {

	parts := strings.Split(file, "/")
	file = strings.Replace(file, parts[len(parts)-1], "", -1)

	if file == "" {
		file = "./"
	}

	log.Println("Changing Directory to ", file)

	err := os.Chdir(file)

	if err != nil {
		panic(err)
	}
}
