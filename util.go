package main

import "sort"

func SortedMap(in map[string]interface{}) map[string]interface{} {
	envMapKeys := make([]string, 0)
	for k := range in {
		envMapKeys = append(envMapKeys, k)
	}
	sort.Strings(envMapKeys)
	orderedMap := make(map[string]interface{})
	for _, key := range envMapKeys {
		orderedMap[key] = in[key]
	}
	return orderedMap
}
