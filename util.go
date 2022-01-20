package main

import (
	"sort"
	"strings"
)

func SortedEnvVarMap(in map[string]interface{}) map[string]interface{} {
	envMapKeys := make([]string, 0)
	for k := range in {
		envMapKeys = append(envMapKeys, k)
	}
	sort.Strings(envMapKeys)
	orderedMap := make(map[string]interface{})
	for _, key := range envMapKeys {
		orderedMap[ToEnv(key)] = in[key]
	}
	return orderedMap
}

func ToEnv(in string) string {
	str := strings.ToUpper(strings.ReplaceAll(in, ".", "_"))
	str = strings.ReplaceAll(str, "[", "_")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, "-", "")
	return str
}
