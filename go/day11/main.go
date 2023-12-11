package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	X int
	Y int
}

func main() {
	b, err := os.ReadFile("day11/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	rows := strings.Split(str, "\n")

	galaxies := make(map[int]Pos)
	galaxyCount := 0

	var emptyX []int
	var emptyY []int

	xCounts := make([]int, len(rows[0]))
	maxX := len(rows[0])
	maxY := len(rows)
	for y, row := range rows {
		if !strings.Contains(row, "#") {
			emptyY = append(emptyY, y)
		}
		for x, v := range strings.Split(row, "") {
			if v == "#" {
				xCounts[x]++
				galaxies[galaxyCount] = Pos{x, y}
				galaxyCount++
			}
		}
	}

	for x, v := range xCounts {
		if v == 0 {
			emptyX = append(emptyX, x)
		}
	}

	expandedUniverse := expand(galaxies, maxX, maxY, emptyX, emptyY, 1)
	part1 := 0
	for a, aPos := range expandedUniverse {
		for b, bPos := range expandedUniverse {
			if a > b {
				part1 += distance(aPos, bPos)
			}
		}
	}
	fmt.Println("part1 = ", part1)

	expandedUniverse2 := expand(galaxies, maxX, maxY, emptyX, emptyY, 1_000_000-1)
	part2 := 0
	for a, aPos := range expandedUniverse2 {
		for b, bPos := range expandedUniverse2 {
			if a > b {
				part2 += distance(aPos, bPos)
			}
		}
	}
	fmt.Println("part2 = ", part2)
}

func distance(a, b Pos) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func expand(galaxies map[int]Pos, maxX, maxY int, emptyX []int, emptyY []int, factor int) map[int]Pos {
	output := make(map[int]Pos)

	xToAdd := make(map[int]int)
	yToAdd := make(map[int]int)

	toAdd := 0
	for x := 0; x < maxX; x++ {
		if x > 0 && slices.Contains(emptyX, x) {
			toAdd += factor
		}
		xToAdd[x] = toAdd
	}
	toAdd = 0
	for y := 0; y < maxY; y++ {
		if y > 0 && slices.Contains(emptyY, y) {
			toAdd += factor
		}
		yToAdd[y] = toAdd
	}
	for k, v := range galaxies {
		output[k] = Pos{v.X + xToAdd[v.X], v.Y + yToAdd[v.Y]}
	}
	return output
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
