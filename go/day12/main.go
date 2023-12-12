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
	s := perturbations.Solver{Cache: make(map[string]int)}

	part1 := 0
	for _, line := range lines {
		counts := stringToIntSlice(strings.Split(line, " ")[1])
		record := strings.TrimSpace(strings.Split(line, " ")[0])

		p := s.ValidArrangements([]rune(record), counts, 0)
		part1 += p
	}
	fmt.Println("part1 = ", part1)

	part2 := 0
	for _, line := range lines {
		counts := []int{}
		baseCounts := stringToIntSlice(strings.Split(line, " ")[1])
		parts := []string{}
		for i := 0; i < 5; i++ {
			parts = append(parts, strings.TrimSpace(strings.Split(line, " ")[0]))
			counts = append(counts, baseCounts...)
		}
		record := strings.Join(parts, "?")

		p := s.ValidArrangements([]rune(record), counts, 0)
		part2 += p
	}
	fmt.Println("part2 = ", part2)
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
