package main

import (
	"embed"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"text/template"

	yaml "gopkg.in/yaml.v2"
)

//go:embed template/config.env.tmpl template/configmap.yaml.tmpl
var resources embed.FS

var TYPE_TO_TEMPLATE_MAP = map[OutType]string{
	TYPE_ENVFILE:   "template/config.env.tmpl",
	TYPE_CONFIGMAP: "template/configmap.yaml.tmpl",
}

// Parses the YAML file located in the specified path and returns the the contents in a map stucture.
// In this implementation, using an array in the YAML root is not supported
func YamlFileToMap(filePath string) map[interface{}]interface{} {
	ext := filepath.Ext(filePath)
	if !(ext == ".yml" || ext == ".yaml") {
		panic("Not a Yaml file " + filePath)
	}
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

// Function map used to define custom behavior that can be used by the templates
func funcMapForTemplates() template.FuncMap {
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

// Structure used as a wrapper for running the templates. The template is chosen based in the OutType.
// The output is filled based on the provided writer (console, file).
type TemplateExportContext struct {
	Type      OutType
	OutWriter io.Writer
}

// Runs the template for the provided data.
func (ctx *TemplateExportContext) RunTemplate(envValues interface{}) {
	tmplPath := TYPE_TO_TEMPLATE_MAP[ctx.Type]
	tmpl := template.Must(template.New(path.Base(tmplPath)).Funcs(funcMapForTemplates()).ParseFS(resources, tmplPath))
	err := tmpl.Execute(ctx.OutWriter, envValues)
	if err != nil {
		panic("Can not execute " + err.Error())
	}
}
