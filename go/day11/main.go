package main

import (
	"fmt"
	"log"
	"math"
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

	expandedUniverse := expand(galaxies, maxX, maxY, emptyX, emptyY)
	part1 := 0
	for a, aPos := range expandedUniverse {
		for b, bPos := range expandedUniverse {
			if a > b {
				part1 += distance(aPos, bPos)
			}
		}
	}
	fmt.Println("part1 = ", part1)
}

func distance(a, b Pos) int {
	return int(math.Abs(float64(a.X)-float64(b.X)) + math.Abs(float64(a.Y)-float64(b.Y)))
}

func expand(galaxies map[int]Pos, maxX, maxY int, emptyX []int, emptyY []int) map[int]Pos {
	output := make(map[int]Pos)

	xToAdd := make(map[int]int)
	yToAdd := make(map[int]int)

	toAdd := 0
	for x := 0; x < maxX; x++ {
		if slices.Contains(emptyX, x) {
			toAdd++
		}
		xToAdd[x] = toAdd
	}
	toAdd = 0
	for y := 0; y < maxY; y++ {
		if slices.Contains(emptyY, y) {
			toAdd++
		}
		yToAdd[y] = toAdd
	}
	for k, v := range galaxies {
		output[k] = Pos{v.X + xToAdd[v.X], v.Y + yToAdd[v.Y]}
	}
	return output
}
