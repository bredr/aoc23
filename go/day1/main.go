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

}
