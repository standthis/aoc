package main

import (
    "fmt"
//    "strconv"
)

var modules = [...]int { 148454, 118001, 98851, 51106, 113158, 139801, 126884, 63241, 71513, 60490, 131129, 76176, 74841, 73589, 130062, 77999, 140758, 98182, 101049, 80951, 75759, 92666, 142078, 89196, 124613, 134713, 75618, 62680, 141366, 108899, 88419, 133385, 90018, 123521, 51919, 58191, 109523, 106012, 94564, 61103, 72803, 66309, 143380, 113708, 146037, 135176, 131177, 77109, 108287, 72170, 87055, 121407, 126216, 139520, 120675, 103833, 130708, 74029, 149840, 117122, 105745, 81186, 51331, 72686, 52095, 72612, 76915, 104859, 114009, 69714, 130716, 133106, 73911, 79766, 56647, 98035, 103504, 93728, 111546, 57637, 68064, 62803, 72759, 144845, 80084, 139247, 139905, 112807, 87844, 149388, 76795, 135703, 120523, 137422, 108335, 60206, 133851, 138574, 141740, 74398 }

func fuel(mass int) int {
    result := mass / 3 - 2
    return result
}

func total() int {
    sum := 0
    for _, mod := range modules {
        sum += moreFuel(mod)
    }
    return sum
}

//Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
func part1() {
    fmt.Println(total())
}

func moreFuel(mod int) int {
    if result := fuel(mod); result < 0 {
        return 0
    } else {
        return result + moreFuel(result)
    }
}

func part2() {
    fmt.Println(total())
}

func main() {
    part1()
}
