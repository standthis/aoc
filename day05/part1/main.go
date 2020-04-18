package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//const add = 1
//const multi = 2
//const input = 3
//const output = 4
//const halt = 99

const add, multi, input, output, halt = 1, 2, 3, 4, 99

// Global mode cases
const (
	ModePosition  = 0
	ModeImmediate = 1
)

func main() {
	content, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(run(string(content)))
}

func run(content string) int {
	var result string
	intcode := stoi(strings.Split(strings.TrimSpace(content), ","))
	//fmt.Println(intcode)

	var idx int
	for {
		//instruction = intcode[idx]
		//arg1 = intcode[idx+1]
		//arg2 = intcode[idx+2]
		//resultPos = intcode[idx+3]
		var (
			arg1      int
			arg2      int
			resultPos int
		)

		fmt.Println(intcode[idx])
		switch intcode[idx] % 100 {
		case add: // 1
			arg1, arg2, resultPos = parseOpcode(intcode, idx)
			intcode[resultPos] = arg1 + arg2
			idx += 4
		case multi: // 2
			arg1, arg2, resultPos = parseOpcode(intcode, idx)
			intcode[resultPos] = arg1 + arg2
			idx += 4
		case input: // 3
			var userInputStr string
			resultPos = intcode[idx+1]
			intcode[resultPos] = userInput(userInputStr)
			idx += 2
		case output: // 4
			switch intcode[idx] / 100 {
			case ModePosition:
				result = result + fmt.Sprint(intcode[intcode[idx+1]])
			case ModeImmediate:
				result = result + fmt.Sprint(intcode[idx+1])
			default:
				panic("invalid output result mode")
			}
			//fmt.Println(arg1)
			idx += 2
		case halt: // 99
			//fmt.Println(intcode)
			fmt.Println(result)
			return intcode[0]
		default:
			fmt.Println(intcode[idx])
			fmt.Println("Intcode error")
			fmt.Println(result)
			return 0

		}
	}

	return -1
}

func parseOpcode(intcode []int, idx int) (arg1 int, arg2 int, resultPos int) {
	var (
		arg1Mode int
		arg2Mode int
		//resultPos int
	)

	arg1Mode = (intcode[idx] % 1000) / 100
	switch arg1Mode {
	case ModePosition:
		arg1 = intcode[intcode[idx+1]]
	case ModeImmediate:
		arg1 = intcode[idx+1]
	default:
		panic("Invalid mode for arg1")
	}

	arg2Mode = (intcode[idx] % 10000) / 1000
	switch arg2Mode {
	case ModePosition:
		arg2 = intcode[intcode[idx+2]]
	case ModeImmediate:
		arg2 = intcode[idx+2]
	default:
		panic("Invalid mode for arg2")
	}

	// perform checks
	resultPos = intcode[idx+3]

	return arg1, arg2, resultPos
}

func userInput(input string) int {
	/*
		fmt.Println("Provide number: ")
		fmt.Scanln(&input)
		number, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
	*/
	return 1
}

func stoi(content []string) []int {
	var intcode = []int{}

	for _, snum := range content {
		inum, err := strconv.Atoi(snum)
		fmt.Println(inum)
		fmt.Println(snum)
		if err != nil {
			panic(err)
		}
		intcode = append(intcode, inum)
	}

	return intcode
}
