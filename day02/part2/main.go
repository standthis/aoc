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

func run(content string, noun int, verb int) int {
    intcode := stoi(strings.Split(strings.TrimSpace(content), ","))
    intcode[1] = noun
    intcode[2] = verb
    for i := 0; i < len(intcode); i += 4 {
        opcode := intcode[i]
        if opcode == add {
            intcode[intcode[i+3]] = intcode[intcode[i+1]] + intcode[intcode[i+2]]
        } else if opcode == multi {
            intcode[intcode[i+3]] = intcode[intcode[i+1]] * intcode[intcode[i+2]]
        } else if opcode == halt {
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

func gravityAssist(content string) int {
    for noun := 0; noun < 100; noun++ { 
        for verb := 0; verb < 100; verb++ { 
            if run(content, noun, verb) == 19690720 {
                return 100 * noun + verb
            }
        }
    }
    return -100
}

func main() {
    content, err := ioutil.ReadFile("../input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(gravityAssist(string(content)))
}
