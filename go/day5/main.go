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
	b, err := os.ReadFile("day5/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	blocks := strings.Split(str, "\n\n")
	reNumber := regexp.MustCompile(`\d+`)

	var seeds []int
	for _, v := range reNumber.FindAllString(blocks[0], -1) {
		value, _ := strconv.Atoi(v)
		seeds = append(seeds, value)
	}

	var maps [][]Mapper
	for _, block := range blocks[1:] {
		maps = append(maps, mapMap(block))
	}

	// part 1
	part1 := math.MaxInt
	for _, seed := range seeds {
		y := seed
		for _, m := range maps {

			for _, r := range m {
				if y >= r.Source && y <= r.Source+r.Range {
					y = (y - r.Source) + r.Dest
					break
				}
			}
		}
		if y < part1 {
			part1 = y
		}
	}
	fmt.Println("part1 = ", part1)
}

type Mapper struct {
	Source int
	Dest   int
	Range  int
}

func mapMap(block string) []Mapper {
	var out []Mapper
	reNumber := regexp.MustCompile(`\d+`)
	for _, row := range strings.Split(block, "\n")[1:] {
		values := make([]int, 3)
		for i, v := range reNumber.FindAllString(row, -1) {
			value, _ := strconv.Atoi(v)
			values[i] = value
		}
		out = append(out, Mapper{Source: values[1], Dest: values[0], Range: values[2]})
	}
	return out
}
