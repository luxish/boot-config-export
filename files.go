package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// Parses the YAML file located in the specified path and returns the the contents in a map stucture.
// In this implementation, using an array in the YAML root is not supported
func YamlFileToMap(filePath string) map[interface{}]interface{} {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("File not found: %s", filePath))
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		panic(fmt.Sprintf("Could not unmarshal yaml file. %s", err.Error()))
	}
	return m
}

// Creates an output channel that accepts strings. All the lines will be appended
// to the file contents.
func MapToOutputFile(fileName string) chan string {
	in := make(chan string)
	outFile, err := os.Create(fileName)
	if err != nil {
		panic(fmt.Sprintf("Could not create the output file "+fileName, err.Error()))
	}
	go func() {
		defer outFile.Close()
		for {
			line, hasNext := <-in
			outFile.WriteString(line + "\n")
			if !hasNext {
				return
			}
		}
	}()
	return in
}

func ChangeExtensionFromFileName(name string) string {
	return strings.TrimRight(strings.TrimRight(name, "yaml"), "yml") + "env"
}

func ResolveOutputFile(inFilePath string, outDir string) string {
	return filepath.Join(outDir, ChangeExtensionFromFileName(filepath.Base(inFilePath)))
}