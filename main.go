package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"
)

// Configuration structure for the CLI
type Config struct {
	fileName  string
	directory string
	outputDir string
}

// Parses the configuration from the program arguments
func ParseConfig() *Config {
	var fileName, dir, outputDir string
	flag.StringVar(&fileName, "f", "", "File to import")
	flag.StringVar(&dir, "d", "", "Directory for input files")
	flag.StringVar(&outputDir, "o", "out", "Directory for the output files")
	flag.Parse()
	return &Config{fileName, dir, outputDir}
}

func processDir(dirPath string, outDir string) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic("Could not read directory " + dirPath)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		processFile(path.Join(dirPath, file.Name()), outDir)
	}
}

func processFile(filePath string, outDir string) {

	if !IsYamlFile(filePath) {
		fmt.Println("Skipped file: " + filePath)
		return
	}

	// Read the provided Yaml file
	root := YamlFileToMap(filePath)

	// Process the contents
	envMap := TraverseYaml(root)

	// Create the file the env file
	envMapKeys := make([]string, 0)
	for k := range envMap {
		envMapKeys = append(envMapKeys, k)
	}
	sort.Strings(envMapKeys)

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.Mkdir(outDir, 0755)
	}
	outFileName := ResolveOutputFile(filePath, outDir)
	out := ToOutputFile(outFileName)
	for _, key := range envMapKeys {
		out <- (key + "=" + envMap[key])
	}
	close(out)
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

	// All configurations exclude each other. In case multiple configurations are specified
	// the application will execute the first one in the order: file, dir.
	if config.fileName != "" {
		processFile(config.fileName, config.outputDir)
	} else if config.directory != "" {
		processDir(config.directory, config.outputDir)
	} else {
		panic("No options found. Run with flag -h to check the usage.")
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in %v", elapsed)
}
