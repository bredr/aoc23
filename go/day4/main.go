package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	b, err := os.ReadFile("day4/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	raw := strings.Split(str, "\n")
	reNumber := regexp.MustCompile(`\d+`)
	part1 := 0
	for _, card := range raw {
		card = strings.Split(card, ":")[1]
		numbers := strings.Split(card, "|")
		winningNumbers := reNumber.FindAllString(numbers[0], -1)
		yourNumbers := reNumber.FindAllString(numbers[1], -1)
		winners := len(Intersection[string](winningNumbers, yourNumbers))
		if winners > 0 {
			part1 += int(math.Pow(2, float64(winners-1)))
		}
	}
	fmt.Println("part1 = ", part1)
}

func Intersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}
