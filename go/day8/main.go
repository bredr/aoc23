package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	b, err := os.ReadFile("day8/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	blocks := strings.Split(str, "\n\n")

	instructions := strings.Split(blocks[0], "")
	nodeRe := regexp.MustCompile("[A-Z0-9]+")
	graph := make(map[string][2]string)
	for _, line := range strings.Split(blocks[1], "\n") {
		matches := nodeRe.FindAllString(line, 3)
		graph[matches[0]] = [2]string{matches[1], matches[2]}
	}

	part1 := 0
	position := "AAA"
loop:
	for {
		for _, instruction := range instructions {
			i := 0
			if instruction == "R" {
				i = 1
			}
			position = graph[position][i]
			part1++
			if position == "ZZZ" {
				break loop
			}
		}
	}
	fmt.Println("part1 = ", part1)

	var positions []string
	var steps []int
	for k := range graph {
		if k[2] == 'A' {
			positions = append(positions, k)
			steps = append(steps, 0)
		}
	}
	for idx, position := range positions {
		step := 0
	loop2:
		for {
			for _, instruction := range instructions {
				i := 0
				if instruction == "R" {
					i = 1
				}
				position = graph[position][i]
				step++
				if position[2] == 'Z' {
					break loop2
				}
			}
		}
		steps[idx] = step
	}
	fmt.Println("part2 = ", LCM(steps[0], steps[1], steps[2:]...))

}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
