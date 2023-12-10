package main

import (
	"fmt"
	"log"
	"math"
	"os"
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

	pipesTypes := make(map[Pos]string)
	maxX := len(strings.Split(rows[0], ""))

	for y, row := range rows {
		for x, v := range strings.Split(row, "") {
			if v == "." {
				continue
			}
			if v == "S" {
				S = Pos{x, y}
			} else {
				pipesTypes[Pos{x, y}] = v
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
	s := solver{pipes}
	loop, _ := s.FindPath(S, []Pos{})
	fmt.Println("part1 = ", math.Ceil(float64(len(loop))/2))

	inLoop := make(map[Pos]string)
	for _, p := range loop {
		inLoop[p] = pipesTypes[p]
	}

	inside := []Pos{}
	for y := 0; y < len(rows); y++ {
		outside := true
		previousBend := ""
		for x := 0; x < maxX; x++ {
			p := Pos{x, y}
			pipe, ok := inLoop[p]
			if ok {
				switch pipe {
				case "|":
					outside = !outside
				case "7":
					if previousBend == "L" {
						outside = !outside
					}
					previousBend = pipe
				case "F":
					previousBend = pipe
				case "L":
					previousBend = pipe
				case "J":
					if previousBend == "F" {
						outside = !outside
					}
					previousBend = pipe
				}
			} else {
				if !outside {
					inside = append(inside, p)
				}
			}
		}
	}
	fmt.Println("part2 = ", len(inside))
}

type solver struct {
	pipes map[Pos]map[Pos]struct{}
}

func (s *solver) FindPath(S Pos, path []Pos) ([]Pos, bool) {
	if len(path) == 0 {
		for _, p := range []Pos{{S.X + 1, S.Y}, {S.X - 1, S.Y}, {S.X, S.Y - 1}, {S.X, S.Y + 1}} {
			if pipe, ok := s.pipes[p]; ok {
				if _, ok := pipe[S]; ok {
					nextPath, ok := s.FindPath(S, append(path, p))
					if ok {
						return nextPath, true
					}
				}
			}
		}
	}
	current := path[len(path)-1]
	for p := range s.pipes[current] {
		if len(path) > 2 && p.X == S.X && p.Y == S.Y {
			// do end of thing
			return path, true
		}
		if _, ok := s.pipes[p]; ok {
			delete(s.pipes[p], current)
			nextPath, ok := s.FindPath(S, append(path, p))
			if ok {
				return nextPath, true
			}
		}
	}
	return path, false
}
