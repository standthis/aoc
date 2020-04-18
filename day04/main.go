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
	content, err := ioutil.ReadFile("input.txt")

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
	double := false

	for _, c := range components {
		if previous > c {
			return false
		}
		if previous == c {
			double = true
		}
		previous = c
	}

	return double
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
			result = append(result, i)
		}
	}
	return result
}
