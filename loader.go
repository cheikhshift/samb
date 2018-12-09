package samb

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/recoye/config"
	"gopkg.in/yaml.v2"
)

// Load opens a project based on the filesystem
// path supplied.
func Load(path string) (*Project, error) {

	// err will be used to check
	// for errors on file load.
	var err error
	var env = &Project{}

	var ext = filepath.Ext(path)

	switch ext {
	case ".yml":
		err = LoadYAML(path, env)
	default:
		err = LoadSE(path, env)
	}

	if err != nil {
		return nil, err
	}

	chDirRootOf(path)

	env.ProcessImports()

	return env, nil
}

// LoadSE loads .se file with SAMB directives.
func LoadSE(path string, p *Project) error {
	var conf = config.New(path)

	return conf.Unmarshal(p)
}

// LoadYAML loads a YAML file with SAMB
// directives.
func LoadYAML(path string, p *Project) error {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, p)
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
