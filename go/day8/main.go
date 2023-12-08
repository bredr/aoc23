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
	nodeRe := regexp.MustCompile("[A-Z]+")
	graph := make(map[string][2]string)
	for _, line := range strings.Split(blocks[1], "\n") {
		matches := nodeRe.FindAllString(line, 3)
		graph[matches[0]] = [2]string{matches[1], matches[2]}
	}

	steps := 0
	position := "AAA"
loop:
	for {
		for _, instruction := range instructions {
			i := 0
			if instruction == "R" {
				i = 1
			}
			position = graph[position][i]
			steps++
			if position == "ZZZ" {
				break loop
			}
		}
	}
	fmt.Println("part1 = ", steps)

}
