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

	//part 2
	var seedRanges [][2]int
	for i := 0; i < len(seeds)/2; i++ {
		seedRanges = append(seedRanges, [2]int{seeds[2*i], seeds[2*i] + seeds[2*i+1] - 1})
	}
	part2 := math.MaxInt
	ranges := seedRanges
	for _, m := range maps {
		next := [][2]int{}
		for {
			if len(ranges) == 0 {
				break
			}
			rang := ranges[0]
			ranges = ranges[1:]
			notInRange := true
			for _, r := range m {
				if rang[0] < r.Source+r.Range && rang[1] >= r.Source {
					next = append(next, [2]int{(max(rang[0], r.Source) - r.Source) + r.Dest, (min(rang[1], r.Source+r.Range-1) - r.Source) + r.Dest})
					if rang[0] < r.Source {
						ranges = append(ranges, [2]int{rang[0], r.Source - 1})
					}
					if rang[1] >= (r.Source + r.Range) {
						ranges = append(ranges, [2]int{r.Source + r.Range, rang[1]})
					}
					notInRange = false
					break
				}
			}
			if notInRange {
				next = append(next, rang)
			}
		}
		ranges = next
	}
	for _, r := range ranges {
		if r[0] < part2 {
			part2 = r[0]
		}
	}
	fmt.Println("part2 = ", part2)
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
