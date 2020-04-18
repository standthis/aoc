package main

import (
	"fmt"
	//"reflect"
	//"io/ioutils"
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(fmt.Errorf("Error opening file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wirePaths := make([]string, 0, 2)
	for scanner.Scan() {
		wirePaths = append(wirePaths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("Error reading file: %s", err))
	}

	grid := make(map[Point]string)

	for i, path := range wirePaths {
		identifier := strconv.Itoa(i)
		if err := ParsePath(path, identifier, grid); err != nil {
			panic(fmt.Errorf("Error parsing path %s: %s", identifier, err))
		}
	}
	distances := FindIntersectionDistances(grid)
	if len(distances) == 0 {
		panic(fmt.Errorf("No intersections found!"))
	}
	sort.Ints(distances)
	fmt.Printf("Smallest distance: %d\n", distances[0])
}

func ParsePath(wirePath string, wireIdentifier string, grid map[Point]string) error {
	pathComponents := strings.Split(wirePath, ",")
	position := Point{0, 0}

	for _, component := range pathComponents {
		direction := component[0]
		distance, err := strconv.Atoi(component[1:])
		if err != nil {
			return err
		}

		for i := 0; i < distance; i++ {
			switch string(direction) {
			case "U":
				position.Y++
			case "D":
				position.Y--
			case "L":
				position.X--
			case "R":
				position.X++
			}

			psVal := grid[position]
			if psVal == wireIdentifier {
				//fmt.Println(grid[position])
				continue
			}
			grid[position] = psVal + wireIdentifier
		}
	}

	return nil
}

func FindIntersectionDistances(grid map[Point]string) []int {
	var intersections []int
	for k, v := range grid {
		if len(v) > 1 {
			result := abs(k.X) + abs(k.Y)
			intersections = append(intersections, result)
		}
	}
	return intersections
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindFewestSteps(grid map[Point]string) int {
	return 5
}
