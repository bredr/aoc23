package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("day3/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	raw := strings.Split(str, "\n")
	var grid [][]string

	for _, row := range raw {
		row = strings.TrimSpace(row)
		grid = append(grid, strings.Split(row, ""))
	}

	part1 := 0
	for y, row := range grid {
		rawNumber := ""
		isPart := false
		for x, v := range row {
			if strings.ContainsAny(v, "0123456789") {
				rawNumber += v

				// check all adjacent if part
				for xx := max(x-1, 0); xx <= min(len(row)-1, x+1); xx++ {
					for yy := max(y-1, 0); yy <= min(len(grid)-1, y+1); yy++ {
						if !(xx == x && yy == y) {
							if !strings.ContainsAny(grid[yy][xx], ".0123456789") {
								isPart = true
							}
						}
					}
				}
			} else {
				if isPart && rawNumber != "" {
					number, err := strconv.Atoi(rawNumber)
					if err != nil {
						panic(err)
					}
					part1 += number
				}
				rawNumber = ""
				isPart = false
			}
		}
		if isPart && rawNumber != "" {
			number, err := strconv.Atoi(rawNumber)
			if err != nil {
				panic(err)
			}
			part1 += number
		}
	}
	fmt.Println("part1 = ", part1)
}
