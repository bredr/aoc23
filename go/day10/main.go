package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
)

type Pos struct {
	X int
	Y int
}

func main() {
	b, err := os.ReadFile("day10/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	rows := strings.Split(str, "\n")

	pipes := make(map[Pos]map[Pos]struct{})
	var S Pos

	for y, row := range rows {
		for x, v := range strings.Split(row, "") {
			if v == "." {
				continue
			}
			if v == "S" {
				S = Pos{x, y}
			} else {
				switch v {
				case "-":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x - 1, y}: {}, {x + 1, y}: {}}
				case "|":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x, y + 1}: {}, {x, y - 1}: {}}
				case "7":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x, y + 1}: {}, {x - 1, y}: {}}
				case "L":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x + 1, y}: {}, {x, y - 1}: {}}
				case "F":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x + 1, y}: {}, {x, y + 1}: {}}
				case "J":
					pipes[Pos{x, y}] = map[Pos]struct{}{{x - 1, y}: {}, {x, y - 1}: {}}
				}
			}
		}
	}
	loop, _ := findPath(S, pipes, []Pos{})
	fmt.Println("part1 = ", math.Ceil(float64(len(loop))/2))

}

func findPath(S Pos, pipes map[Pos]map[Pos]struct{}, path []Pos) ([]Pos, bool) {
	if len(path) == 0 {
		for _, p := range []Pos{{S.X + 1, S.Y}, {S.X - 1, S.Y}, {S.X, S.Y - 1}, {S.X, S.Y + 1}} {
			if pipe, ok := pipes[p]; ok {
				if _, ok := pipe[S]; ok {
					nextPath, ok := findPath(S, pipes, append(path, p))
					if ok {
						return nextPath, true
					}
				}
			}
		}
	}
	current := path[len(path)-1]
	for p := range pipes[current] {
		if len(path) > 2 && reflect.DeepEqual(p, S) {
			// do end of thing
			return path, true
		}
		if pipe, ok := pipes[p]; ok {
			if _, ok := pipe[current]; ok {
				if !contains(path, p) {
					nextPath, ok := findPath(S, pipes, append(path, p))
					if ok {
						return nextPath, true
					}
				}
			}
		}
	}
	return path, false
}

func contains(s []Pos, e Pos) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}
