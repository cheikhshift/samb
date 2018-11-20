// Package transpiler converts samb struct types into Go code.
package transpiler

import (
	"os"
)

// Setup invokes local function
// MakePkgPaths.
func Setup() {

	MakePkgPaths()

}

// MakePkgPaths creates the folders
// the transpiler will write to. The folders
// are created relative to your program's
// working directory. They are as follows :
// ./pkg/api, ./pkg/globals, ./pkg/method
//, ./cmd/server
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
