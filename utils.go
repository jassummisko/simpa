package main

import (
	"sort"
	"strings"
)

func removePhs(s string) string {
	var buff string = ""
	for _, char := range s {
		if char != '◌' {
			buff = buff + string(char)
		}
	}
	return buff
}

func matchFirst(s string, s2 string) bool {
	if len(s) > len(s2) {
		return false
	}
	return s == s2[:len(s)]
}

func findBiggestThatFits(s string, list []string) string {
	for _, el := range sortStringsByLength(list) {
		if strings.HasPrefix(s, el) {
			return el
		}
	}
	return ""
}

func sortStringsByLength(list []string) []string {
	sort.Slice(list, func(i, j int) bool {
		l1, l2 := len(list[i]), len(list[j])
		if l1 != l2 {
			return l1 > l2
		}
		return list[i] > list[j]
	})
	return list
}

func getInitMatchingKeys(s string, m map[string]string) []string {
	keys := []string{}
	for k := range m {
		if matchFirst(s, k) {
			keys = append(keys, k)
		}
	}
	return keys
}

func copyMapWithDifferences(input map[string]string, differences map[string]string) map[string]string {
	buffer := make(map[string]string)
	for k, v := range input {
		buffer[k] = v
	}
	for k, v := range differences {
		buffer[k] = v
	}
	return buffer
}
