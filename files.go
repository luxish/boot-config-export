package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

type OutType string

const (
	YAML_FILE_EXT          = ".yaml"
	YML_FILE_EXT           = ".yml"
	ENV_FILE_EXT           = ".env"
	TYPE_CONFIGMAP OutType = "cm"
	TYPE_ENVFILE   OutType = "env"
)

var TYPE_TO_TEMPLATE_MAP = map[OutType]string{
	TYPE_ENVFILE:   "template/config.env.tmpl",
	TYPE_CONFIGMAP: "template/configmap.yaml.tmpl",
}

var TYPE_TO_EXT_MAP = map[OutType]string{
	TYPE_ENVFILE:   ENV_FILE_EXT,
	TYPE_CONFIGMAP: YAML_FILE_EXT,
}

type ExportContext struct {
	OutType  OutType
	OutDir   string
	FileName string
}

func (ctx *ExportContext) TemplatePath() string {
	return TYPE_TO_TEMPLATE_MAP[ctx.OutType]
}

func (ctx *ExportContext) OutputFile() string {
	return path.Join(ctx.OutDir, ctx.FileName) + TYPE_TO_EXT_MAP[ctx.OutType]
}

func (ctx *ExportContext) RunTemplate(envValues map[string]interface{}) {
	tmplPath := ctx.TemplatePath()
	outFileName := ctx.OutputFile()

	ensureDirExists(ctx.OutDir)
	outFile, err := os.Create(outFileName)

	if err != nil {
		panic(fmt.Sprintf("Can not create file %s: %s", outFileName, err.Error()))
	}

	tmpl := template.Must(template.New(path.Base(tmplPath)).Funcs(createFuncMapForTemplates()).ParseFiles(tmplPath))
	err = tmpl.Execute(outFile, envValues)
	if err != nil {
		panic("Can not execute " + err.Error())
	}
}

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

func IsYamlFile(filePath string) bool {
	ext := filepath.Ext(filePath)
	return ext == YAML_FILE_EXT || ext == YML_FILE_EXT
}

func FileName(filePath string) string {
	return strings.TrimRight(strings.TrimRight(filepath.Base(filePath), YAML_FILE_EXT), YML_FILE_EXT)
}

func OutTypeFromString(str string) OutType {
	switch str {
	case "cm":
		return TYPE_CONFIGMAP
	case "env":
		return TYPE_ENVFILE
	default:
		return TYPE_ENVFILE
	}
}

func ensureDirExists(outDir string) {
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		os.Mkdir(outDir, 0755)
	}
}

func createFuncMapForTemplates() template.FuncMap {
	return template.FuncMap{
		"quoteIfString": func(arg0 reflect.Value, args ...reflect.Value) reflect.Value {
			kind := reflect.TypeOf(arg0.Interface()).Kind()
			if kind == reflect.String {
				return reflect.ValueOf(fmt.Sprintf("%q", arg0))
			}
			return arg0
		},
		"quote": func(arg0 reflect.Value, args ...reflect.Value) reflect.Value {
			return reflect.ValueOf(fmt.Sprintf("\"%v\"", arg0))
		},
	}
}
