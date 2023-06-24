package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	usemap                  map[string]string
	cxsFlag                 *bool
	replFlag                *bool
	onlyBetweenBracketsFlag *bool
)

func main() {
	usemap = xsampamap

	cxsFlag = flag.Bool("c", false, "use CSX (Conlang X-SAMPA)")
	replFlag = flag.Bool("r", false, "REPL mode")
	onlyBetweenBracketsFlag = flag.Bool("b", false, "Parse only values between // or []")

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
	fmt.Print(mapxsampa(finalstr, usemap, *onlyBetweenBracketsFlag))
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
		out := mapxsampa(input, inputMap, *onlyBetweenBracketsFlag)
		fmt.Println(out)
	}
}

func mapxsampa(input string, inputMap map[string]string, onlyBetweenBrackets bool) string {
	buff := ""
	inBrackets := !onlyBetweenBrackets

	for len(input) > 0 {
		temp := string(input[0])
		if onlyBetweenBrackets {
			if temp == "[" || temp == "/" {
				inBrackets = true
			} else if temp == "]" || temp == "/" {
				inBrackets = false
			}
		}
		options := getInitMatchingKeys(temp, inputMap)
		match := findBiggestThatFits(input, options)
		toAdd := ""
		if match != "" && inBrackets {
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
