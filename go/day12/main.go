package main

import (
	"aoc23/day12/perturbations"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day12/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")

	part1 := 0
	for _, line := range lines {
		counts := stringToIntSlice(strings.Split(line, " ")[1])
		record := strings.TrimSpace(strings.Split(line, " ")[0])

		p := perturbations.Perturbations(record, counts)
		part1 += len(p)
	}
	fmt.Println("part1 = ", part1)
}

func stringToIntSlice(x string) []int {
	reN := regexp.MustCompile(`-?\d+`)
	var numbers []int
	for _, m := range reN.FindAllString(x, -1) {
		v, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, v)
	}
	return numbers
}
