package main

import (
	"bytes"
	"encoding/gob"
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

		rolledColumn := roll(column)
		part1 += loadNorth(rolledColumn)
	}
	fmt.Println("part1 = ", part1)
	s := solver{make(map[string]PointInTime)}
	var cache *int
	periodicity := 0
	offset := 0
	target := 1000000000
	for c := 1; c <= target; c++ {
		grid, cache = s.Spin(grid, c)
		if periodicity > 0 {
			if (c-offset)%periodicity == (target-offset)%periodicity {
				break
			}
		}
		if cache != nil {
			periodicity = c - *cache
			offset = *cache
		}
	}
	part2 := 0
	for x := 0; x < len(grid[0]); x++ {
		column := []rune{}
		for _, row := range grid {
			column = append(column, row[x])
		}
		part2 += loadNorth(column)
	}
	fmt.Println("part2 = ", part2)
}

type PointInTime struct {
	Grid       [][]rune
	FirstIndex int
}

type solver struct {
	cache map[string]PointInTime
}

func hash(matrix [][]rune) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(matrix)
	return b.String()
}

func (s *solver) Spin(matrix [][]rune, i int) ([][]rune, *int) {
	key := hash(matrix)
	if v, ok := s.cache[key]; ok {
		return v.Grid, &v.FirstIndex
	}
	for i := 0; i < 4; i++ {
		matrix = reverse(transpose(rollGrid(matrix)))
	}
	s.cache[key] = PointInTime{make([][]rune, len(matrix)), i}
	for i, v := range matrix {
		s.cache[key].Grid[i] = make([]rune, len(matrix[0]))
		copy(s.cache[key].Grid[i], v)
	}
	return matrix, nil
}

func printGrid(matrix [][]rune) {
	for _, row := range matrix {
		for _, v := range row {
			fmt.Print(string([]rune{v}))
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func reverse(matrix [][]rune) [][]rune {
	output := make([][]rune, len(matrix))
	for i, row := range matrix {
		c := make([]rune, len(row))
		copy(c, row)
		slices.Reverse(c)
		output[i] = c
	}
	return output
}

func transpose(matrix [][]rune) [][]rune {
	output := make([][]rune, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		row := make([]rune, len(matrix))
		for j := 0; j < len(matrix[0]); j++ {
			row[j] = matrix[j][i]
		}
		output[i] = row
	}
	return output
}

func rollGrid(matrix [][]rune) [][]rune {
	output := make([][]rune, len(matrix))
	for y := range output {
		output[y] = make([]rune, len(matrix[0]))
	}

	for x := 0; x < len(matrix[0]); x++ {
		column := []rune{}
		for _, row := range matrix {
			column = append(column, row[x])
		}
		for y, v := range roll(column) {
			output[y][x] = v
		}
	}
	return output
}

func roll(column []rune) []rune {
	output := make([]rune, len(column))
	copy(output, column)
	moves := 0
	for i, v := range output {
		if v == 'O' {
			if i > 0 && output[i-1] == '.' {
				output[i-1] = 'O'
				output[i] = '.'
				moves++
			}
		}
	}
	if moves > 0 {
		return roll(output)
	}
	return output
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
