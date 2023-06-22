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
	usemap := xsampamap
	cxsFlag := flag.Bool("c", false, "use CSX (Conlang X-SAMPA)")
	replFlag := flag.Bool("r", false, "REPL mode")
	flag.Parse()

	if *cxsFlag {
		usemap = cxsmap
	}

	if *replFlag {
		replmode(usemap)
		return
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

func replmode(inputMap map[string]string) {
	var input string
	var history []string
	for {
		fmt.Scan(&input)
		if input == "---" {
			history = append(history, input)
			break
		}
		out := mapxsampa(input, inputMap)
		fmt.Println(out)
	}
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
