package main

import (
	"fmt"
	"reflect"
)

// Structure used to traverse a YAML file. The path is the nodes path from the root (JSON path like)
// The value keeps the nodes content for that path.
type Node struct {
	Path  string
	Value interface{}
}

// Function used to traverse YAML contents based on an initial map. It also returns a map, with a
// flattened representation. The key is a JSON-like path of the (primitive) value.
func TraverseYaml(root map[interface{}]interface{}) map[string]interface{} {
	envMap := make(map[string]interface{})
	travArr := make([]Node, 0)

	for key, val := range root {
		travArr = append(travArr, Node{key.(string), val})
	}

	for len(travArr) > 0 {
		// Extract last element
		lastEl := travArr[len(travArr)-1]
		travArr = travArr[:len(travArr)-1]

		if lastEl.Value == nil {
			continue
		}

		switch kind := reflect.TypeOf(lastEl.Value).Kind(); kind {
		case reflect.Int, reflect.Bool, reflect.Float32, reflect.Float64, reflect.String:
			// Leaf values can be added to the output map
			envMap[lastEl.Path] = lastEl.Value
		case reflect.Slice:
			for idx, val := range lastEl.Value.([]interface{}) {
				travArr = append(travArr, Node{fmt.Sprintf("%s.%v", lastEl.Path, idx), val})
			}
		case reflect.Map:
			for key, val := range lastEl.Value.(map[interface{}]interface{}) {
				travArr = append(travArr, Node{fmt.Sprintf("%s.%v", lastEl.Path, key), val})
			}
		}
	}

	return envMap
}
