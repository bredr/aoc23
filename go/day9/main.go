package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day9/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")
	part1 := 0
	for _, line := range lines {
		numbers := stringToIntSlice(line)
		part1 += solve(numbers)
	}

	fmt.Println("part1 = ", part1)
	part2 := 0
	for _, line := range lines {
		numbers := stringToIntSlice(line)
		slices.Reverse(numbers)
		part2 += solve(numbers)
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

func solve(x []int) int {
	diffs := make([]int, len(x)-1)
	same := true
	for i := range diffs {
		diffs[i] = x[i+1] - x[i]
		if i > 1 && diffs[i] != diffs[i-1] {
			same = false
		}
	}
	if same {
		return x[len(x)-1] + diffs[0]
	}
	return x[len(x)-1] + solve(diffs)
}
