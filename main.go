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

func processFile(filePath string, outType OutType, outDir string) {
	if !IsYamlFile(filePath) {
		fmt.Println("Not a Yaml file " + filePath)
		return
	}

	// Read the provided Yaml file
	root := YamlFileToMap(filePath)
	// Process the contents
	envMap := TraverseYaml(root)

	ctx := ExportContext{outType, outDir, FileName(filePath)}
	ctx.RunTemplate(SortedMap(envMap))
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

	processFile(config.fileName, OutTypeFromString(config.outType), config.outputDir)

	elapsed := time.Since(start)
	fmt.Printf("Done in %v", elapsed)
}
