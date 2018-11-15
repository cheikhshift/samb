package samb

import "github.com/recoye/config"

func Load(path string) (*Project, error) {

	conf := config.New(path)

	env := &Project{}

	err := conf.Unmarshal(env)
	if err != nil {
		return nil, err
	}

	env.ProcessImports()
	env.ProcessServerImports()

	return env, nil
}
