package transpiler

import (
	"os"
)

func Setup() {

	MakePkgPaths()

}

func MakePkgPaths() {
	err := os.MkdirAll("./pkg/api", 0700)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./pkg/globals", 0700)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./pkg/methods", 0700)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("./cmd/server", 0700)

	if err != nil {
		panic(err)
	}
}
