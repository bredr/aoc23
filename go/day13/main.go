package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("day13/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	blocks := strings.Split(str, "\n\n")

	part1 := 0
	for _, block := range blocks {
		field := parseBlock(block)
		part1 += findMirror(field, 0)
	}
	fmt.Println("part1 = ", part1)
	part2 := 0
	for _, block := range blocks {
		field := parseBlock(block)
		part2 += findMirror(field, 1)
	}
	fmt.Println("part2 = ", part2)
}

type Pos struct {
	X int
	Y int
}

type Block struct {
	MaxX  int
	MaxY  int
	Rocks map[Pos]struct{}
}

func parseBlock(block string) Block {
	rocks := make(map[Pos]struct{})
	for y, row := range strings.Split(block, "\n") {
		for x, v := range strings.Split(row, "") {
			if v == "#" {
				rocks[Pos{x, y}] = struct{}{}
			}
		}
	}
	return Block{len(strings.Split(block, "\n")[0]), len(strings.Split(block, "\n")), rocks}
}

func findMirror(block Block, maxMistakes int) int {
	for mirrorX := 1; mirrorX < block.MaxX; mirrorX++ {
		delta := mirrorX
		if mirrorX > block.MaxX/2 {
			delta = block.MaxX - mirrorX
		}
		mistakes := 0
	xLoop:
		for y := 0; y < block.MaxY; y++ {
			for i := 0; i < delta; i++ {
				_, right := block.Rocks[Pos{mirrorX + i, y}]
				_, left := block.Rocks[Pos{mirrorX - 1 - i, y}]
				if left != right {
					mistakes++
					if mistakes > maxMistakes {
						break xLoop
					}
				}
			}
		}
		if mistakes == maxMistakes {
			return mirrorX
		}
	}

	for mirrorY := 1; mirrorY < block.MaxY; mirrorY++ {
		delta := mirrorY
		if mirrorY > block.MaxY/2 {
			delta = block.MaxY - mirrorY
		}
		mistakes := 0
	yLoop:
		for x := 0; x < block.MaxX; x++ {
			for i := 0; i < delta; i++ {
				_, down := block.Rocks[Pos{x, mirrorY + i}]
				_, up := block.Rocks[Pos{x, mirrorY - 1 - i}]
				if up != down {
					mistakes++
					if mistakes > maxMistakes {
						break yLoop
					}
				}
			}

		}
		if mistakes == maxMistakes {
			return mirrorY * 100
		}
	}
	panic("not found mirror")
}
