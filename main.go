package main

import (
	"flag"
	"fmt"
	"time"
)

// Configuration structure for the CLI
type Config struct {
	fileName  string
	outType   string
	outputDir string
}

// Parses the configuration from the program arguments
func ParseConfig() *Config {
	var fileName, outType, outputDir string
	flag.StringVar(&fileName, "f", "", "File to import")
	flag.StringVar(&outType, "t", "env", "Output type: env, cm")
	flag.StringVar(&outputDir, "o", "out", "Directory for the output files")
	flag.Parse()
	return &Config{fileName, outType, outputDir}
}

func processFile(config Config) {
	if !IsYamlFile(config.fileName) {
		fmt.Println("Not a Yaml file " + config.fileName)
		return
	}

	// Read the provided Yaml file
	root := YamlFileToMap(config.fileName)
	// Process the contents
	envMap := TraverseYaml(root)

	extConfig := CreateExternalConfig(config, envMap)
	extConfig.Transform()
}

func main() {
	// In case of panic, print the error and skip the stacktrace.
	defer func() {
		recovery := recover()
		if recovery != nil {
			fmt.Println(recovery)
		}
	}()

	start := time.Now()

	// Parse configuration from CLI.
	config := ParseConfig()

	if config.fileName == "" {
		panic("No file to process. Run with flag -h to check the usage.")
	}

	processFile(*config)

	elapsed := time.Since(start)
	fmt.Printf("Done in %v", elapsed)
}
