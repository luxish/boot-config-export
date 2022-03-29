package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

type OutType string

const (
	TYPE_ENVFILE   OutType = "env"
	TYPE_CONFIGMAP OutType = "cm"
)

type ExternalConfig interface {
	Transform()
}

type ExternalConfigAbstract struct {
	OutputFileName string
	Values         map[string]interface{}
}

func (extconfigAbs ExternalConfigAbstract) outputWriter() io.Writer {
	// Console output
	if extconfigAbs.OutputFileName == "" {
		return os.Stdout
	}

	// File output
	outDir, _ := path.Split(extconfigAbs.OutputFileName)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.Mkdir(outDir, 0755)
	}
	outFile, err := os.Create(extconfigAbs.OutputFileName)
	if err != nil {
		panic(fmt.Sprintf("Can not create file %s: %s", extconfigAbs.OutputFileName, err.Error()))
	}

	return outFile
}

// Environment variables file external configuration
type EnvFileExternalConfig struct {
	ExternalConfigAbstract
}

func (envFileCtx EnvFileExternalConfig) Transform() {
	ctx := TemplateExportContext{TYPE_ENVFILE, envFileCtx.outputWriter()}
	ctx.RunTemplate(SortedEnvVarMap(envFileCtx.Values))
}

// K8s ConfigMap file external configuration
type ConfigMapFileExternalConfig struct {
	ExternalConfigAbstract
}

func (cmFileCtx ConfigMapFileExternalConfig) Transform() {
	ctx := TemplateExportContext{TYPE_CONFIGMAP, cmFileCtx.outputWriter()}
	ctx.RunTemplate(SortedEnvVarMap(cmFileCtx.Values))
}

// Factory function
func CreateExternalConfig(fileName string, outputFileName string, outType string, values map[string]interface{}) ExternalConfig {
	switch outType {
	case "cm":
		return ConfigMapFileExternalConfig{ExternalConfigAbstract{outputFileName, values}}
	default:
		return EnvFileExternalConfig{ExternalConfigAbstract{outputFileName, values}}
	}
}
