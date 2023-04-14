package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func mapxsampa(input string) string {
	buff := ""
	for len(input) > 0 {
		temp := string(input[0])
		options := getInitMatchingKeys(temp, xsampamap)
		match := findBiggestThatFits(input, options)
		toAdd := ""
		if match != "" {
			toAdd = xsampamap[match]
		} else {
			toAdd = temp
			match = temp
		}
		buff += toAdd
		input = input[len(match):]
	}
	return removePhs(buff)
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	finalstr := ""
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("Error reading input:", err)
			return
		}
		finalstr += str
	}
	finalstr = strings.TrimRight(finalstr, "\n")
	fmt.Print(mapxsampa(finalstr))
}
