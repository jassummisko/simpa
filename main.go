package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var usemap map[string]string
	cxsFlag := flag.Bool("c", false, "use CSX (Conlang X-SAMPA)")
	flag.Parse()
	if *cxsFlag {
		usemap = cxsmap
	} else {
		usemap = xsampamap
	}
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
	fmt.Print(mapxsampa(finalstr, usemap))
}

func mapxsampa(input string, inputMap map[string]string) string {
	buff := ""
	for len(input) > 0 {
		temp := string(input[0])
		options := getInitMatchingKeys(temp, inputMap)
		match := findBiggestThatFits(input, options)
		toAdd := ""
		if match != "" {
			toAdd = inputMap[match]
		} else {
			toAdd = temp
			match = temp
		}
		buff += toAdd
		input = input[len(match):]
	}
	return removePhs(buff)
}
