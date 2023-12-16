package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("day16/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	var grid [][]rune
	for _, line := range strings.Split(str, "\n") {
		grid = append(grid, []rune(line))
	}

	s := solver{grid, len(grid[0]) - 1, len(grid) - 1}
	energised := s.Traverse(Vec{-1, 0}, Vec{1, 0}, make(map[Vec]map[Vec]struct{}))
	fmt.Println("part1 = ", len(energised))

	part2 := 0
	for y := 0; y < s.MaxY; y++ {
		energised := s.Traverse(Vec{-1, y}, Vec{1, 0}, make(map[Vec]map[Vec]struct{}))
		if len(energised) > part2 {
			part2 = len(energised)
		}
		energised = s.Traverse(Vec{s.MaxX + 1, y}, Vec{-1, 0}, make(map[Vec]map[Vec]struct{}))
		if len(energised) > part2 {
			part2 = len(energised)
		}
	}
	for x := 0; x < s.MaxX; x++ {
		energised := s.Traverse(Vec{x, -1}, Vec{0, 1}, make(map[Vec]map[Vec]struct{}))
		if len(energised) > part2 {
			part2 = len(energised)
		}
		energised = s.Traverse(Vec{x, s.MaxY + 1}, Vec{0, -1}, make(map[Vec]map[Vec]struct{}))
		if len(energised) > part2 {
			part2 = len(energised)
		}
	}
	fmt.Println("part2 = ", part2)
}

type Vec struct {
	X, Y int
}

func PrintGrid(grid [][]rune, paths map[Vec]map[Vec]struct{}) {
	fmt.Println()

	for y, row := range grid {
		for x := range row {
			if _, ok := paths[Vec{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

type solver struct {
	Grid       [][]rune
	MaxX, MaxY int
}

type Beam struct {
	Pos       Vec
	Direction Vec
}

func (s *solver) Traverse(position Vec, direction Vec, energised map[Vec]map[Vec]struct{}) map[Vec]map[Vec]struct{} {
	// loop protection
	if position.X >= 0 && position.Y >= 0 && position.X <= s.MaxX && position.Y <= s.MaxY {
		if v, ok := energised[position]; ok {
			if _, ok := v[direction]; ok {
				return energised
			} else {
				energised[position][direction] = struct{}{}
			}
		} else {
			energised[position] = make(map[Vec]struct{})
			energised[position][direction] = struct{}{}
		}
	}

	next := Vec{position.X + direction.X, position.Y + direction.Y}
	if next.X < 0 || next.X > s.MaxX || next.Y < 0 || next.Y > s.MaxY {
		return energised
	}
	switch s.Grid[next.Y][next.X] {
	case '.':
		return s.Traverse(next, direction, energised)
	case '/':
		return s.Traverse(next, Vec{direction.Y * -1, direction.X * -1}, energised)
	case '\\':
		if direction.X == 0 {
			return s.Traverse(next, Vec{direction.Y, direction.X}, energised)
		}
		return s.Traverse(next, Vec{direction.Y, direction.X}, energised)
	case '-':
		if direction.Y == 0 {
			return s.Traverse(next, direction, energised)
		}
		a := s.Traverse(next, Vec{-1, 0}, energised)
		b := s.Traverse(next, Vec{1, 0}, energised)
		for k, v := range a {
			for direction := range v {
				if _, ok := energised[k]; ok {
					energised[k][direction] = struct{}{}
				} else {
					energised[k] = make(map[Vec]struct{})
					energised[k][direction] = struct{}{}
				}
			}
		}
		for k, v := range b {
			for direction := range v {
				if _, ok := energised[k]; ok {
					energised[k][direction] = struct{}{}
				} else {
					energised[k] = make(map[Vec]struct{})
					energised[k][direction] = struct{}{}
				}
			}
		}
		return energised
	case '|':
		if direction.X == 0 {
			return s.Traverse(next, direction, energised)
		}
		a := s.Traverse(next, Vec{0, 1}, energised)
		b := s.Traverse(next, Vec{0, -1}, energised)
		for k, v := range a {
			for direction := range v {
				if _, ok := energised[k]; ok {
					energised[k][direction] = struct{}{}
				} else {
					energised[k] = make(map[Vec]struct{})
					energised[k][direction] = struct{}{}
				}
			}
		}
		for k, v := range b {
			for direction := range v {
				if _, ok := energised[k]; ok {
					energised[k][direction] = struct{}{}
				} else {
					energised[k] = make(map[Vec]struct{})
					energised[k][direction] = struct{}{}
				}
			}
		}
		return energised
	default:
		panic("not meant to be here")
	}
}
