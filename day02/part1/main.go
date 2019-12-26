package main

import (
	"io/ioutil"
    "fmt"
    "strings"
    "os"
    "strconv"
)

const add = 1
const multi = 2
const halt = 99

func run(content string) int {
    intcode := stoi(strings.Split(strings.TrimSpace(content), ","))
    fmt.Println(intcode)

    for i := 0; i < len(intcode); i += 4 {
        if intcode[i] == add {
            intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
        } else if intcode[i] == multi {
            intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
        } else if intcode[i] == halt {
            fmt.Println(intcode)
            return intcode[0]
        }else {
            fmt.Println("Intcode error")
        }
    }

    return -1
}

func stoi(content []string) []int {
    var intcode = []int{}

    for _, snum := range content {
        inum, err := strconv.Atoi(snum)
        if err != nil {
            panic(err)
        }
        intcode = append(intcode, inum)
    }
    return intcode
}

func main() {
    content, err := ioutil.ReadFile("../input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(run(string(content)))
}
