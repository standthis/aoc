package main

import (
	"fmt"
	"io/ioutil"
	"os"
	//"reflect"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("../input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := checkRange(string(content))
	fmt.Println(len(result))
}

func accepted(input string) bool {
	components := strings.Split(input, "")
	previous := "0"
	seen := make(map[string]int)

	for _, c := range components {
		if previous > c {
			return false
		}
		seen[c] = seen[c] + 1
		previous = c
	}

	fmt.Println(seen)
	for _, v := range seen {
		if v == 2 {
			return true
		}
	}
	return false
}

func checkRange(inputRange string) []int {
	therange := make([]int, 0)
	result := make([]int, 0)
	inputRange = strings.TrimSpace(inputRange)
	tmp := strings.Split(inputRange, "-")
	for _, t := range tmp {
		i, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println(err)
		}
		therange = append(therange, i)
	}
	for i := therange[0]; i < therange[1]; i++ {
		inputString := strconv.Itoa(i)
		if accept := accepted(inputString); accept == true {
			fmt.Println(inputString)
			result = append(result, i)
		}
	}
	return result
}
