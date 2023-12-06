package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day6/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	blocks := strings.Split(str, "\n")
	reNumber := regexp.MustCompile(`\d+`)
	var times []int
	var distances []int

	for _, n := range reNumber.FindAllString(blocks[0], -1) {
		v, _ := strconv.Atoi(n)
		times = append(times, v)
	}
	for _, n := range reNumber.FindAllString(blocks[1], -1) {
		v, _ := strconv.Atoi(n)
		distances = append(distances, v)
	}
	part1 := 1
	for i, t := range times {
		d := distances[i]
		n1 := math.Ceil((float64(t) - math.Sqrt(float64(t*t-4*d))) / 2)
		n2 := math.Floor((float64(t) + math.Sqrt(float64(t*t-4*d))) / 2)
		part1 *= int(1 + n2 - n1)
	}
	fmt.Println("part1 = ", part1)

	t, _ := strconv.Atoi(strings.Join(reNumber.FindAllString(blocks[0], -1), ""))
	d, _ := strconv.Atoi(strings.Join(reNumber.FindAllString(blocks[1], -1), ""))
	n1 := math.Ceil((float64(t) - math.Sqrt(float64(t*t-4*d))) / 2)
	n2 := math.Floor((float64(t) + math.Sqrt(float64(t*t-4*d))) / 2)
	part2 := int(1 + n2 - n1)

	fmt.Println("part2 = ", part2)
}
