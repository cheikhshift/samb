package samb

import (
	"log"
	"os"
	"strings"

	"github.com/recoye/config"
)

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
