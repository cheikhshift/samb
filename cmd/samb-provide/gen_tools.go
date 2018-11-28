package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func loadFile(path string) (string, error) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	stringData := string(data)

	return stringData, nil

}

func addProviderDirective(providers, name, t string) {

	if strings.Contains(providers, name) {
		log.Println("Provider already declared, skipping addition.")
		os.Exit(1)
	}

	providers += fmt.Sprintf(providerTemplate,
		name,
		t,
		name,
	)

	err := ioutil.WriteFile(*providerFile, []byte(providers), 0700)

	if err != nil {
		panic(err)
	}
}

func saveProviderFunc(name, t string) {

	f := fmt.Sprintf(fnTemplate, name, name, t)
	fileName := fmt.Sprintf(nameTemplate, name)
	writeTo := filepath.Join("pkg", "api", fileName)

	err := ioutil.WriteFile(writeTo, []byte(f), 0700)

	if err != nil {
		panic(err)
	}

	fmt.Println("Saved file to : ", writeTo)
}
