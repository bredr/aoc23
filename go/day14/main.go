package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	X int
	Y int
}

func main() {
	b, err := os.ReadFile("day14/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	rows := strings.Split(str, "\n")

	grid := [][]rune{}
	for _, row := range rows {
		r := []rune{}
		for _, v := range []rune(row) {
			r = append(r, v)
		}
		grid = append(grid, r)
	}

	part1 := 0
	for x := 0; x < len(grid[0]); x++ {
		column := []rune{}
		for _, row := range grid {
			column = append(column, row[x])
		}

		rolledColumn := rollNorth(column)
		part1 += loadNorth(rolledColumn)
	}
	fmt.Println("part1 = ", part1)
}

func rollNorth(column []rune) []rune {
	moves := 0
	for i, v := range column {
		if v == 'O' {
			if i > 0 && column[i-1] == '.' {
				column[i-1] = 'O'
				column[i] = '.'
				moves++
			}
		}
	}
	if moves > 0 {
		return rollNorth(column)
	}
	return column
}

func loadNorth(column []rune) int {
	sum := 0
	for i, v := range column {
		if v == 'O' {
			sum += (len(column) - i)
		}
	}
	return sum
}
