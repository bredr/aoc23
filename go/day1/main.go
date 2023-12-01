package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day1/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	raw := strings.Split(str, "\n")
	r := regexp.MustCompile(`\d`)

	part1 := 0
	for _, line := range raw {
		digits := r.FindAllString(line, -1)
		value, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			log.Fatal(err)
		}
		part1 += value
	}
	fmt.Println("part1 = ", part1)
	part2 := 0
	for _, line := range raw {
		digits := findWithOverlaps(line)
		first := digits[0]
		last := digits[len(digits)-1]
		value, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		part2 += value
	}
	fmt.Println("part2 = ", part2)
}

func findWithOverlaps(x string) []string {
	r := regexp.MustCompile(`^(\d|zero|one|two|three|four|five|six|seven|eight|nine)`)
	lookup := map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	var found []string
	for i := 0; i < len(x); i++ {
		match := r.FindString(x[i:])
		if match != "" {
			if value, ok := lookup[match]; ok {
				match = value
			}
			found = append(found, match)

		}
	}
	return found
}
