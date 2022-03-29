package main

import (
	"flag"
	"fmt"
)

func main() {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
		}
	}()

	var fileName, outType, outputFileName string
	flag.StringVar(&fileName, "f", "", "(Mandatory) YAML configuration file to trasform.")
	flag.StringVar(&outType, "t", "env", "Output type: env, cm")
	flag.StringVar(&outputFileName, "o", "", "Output file name.")
	flag.Parse()

	if fileName == "" {
		panic("No file to process. Run with flag -h to check the usage.")
	}

	// Read the provided Yaml file
	root := YamlFileToMap(fileName)
	// Process the contents
	envMap := TraverseYaml(root)
	// Run transformation
	extConfigCtx := CreateExternalConfig(fileName, outputFileName, outType, envMap)
	extConfigCtx.Transform()
}
