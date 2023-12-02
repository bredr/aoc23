package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Round struct {
	Blue  int
	Red   int
	Green int
}
type Game struct {
	ID     int
	Rounds []Round
}

func main() {
	b, err := os.ReadFile("day2/input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	raw := strings.Split(str, "\n")
	var games []Game
	idRegex := regexp.MustCompile(`\d+`)
	for _, line := range raw {
		match := idRegex.FindString(line)
		ID, err := strconv.Atoi(match)
		if err != nil {
			log.Fatal(err)
		}
		game := Game{ID: ID}
		for _, roundRaw := range strings.Split(strings.Split(line, ":")[1], ";") {
			round := Round{}
			for _, countRaw := range strings.Split(roundRaw, ",") {
				match := idRegex.FindString(countRaw)
				count, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				if strings.Contains(countRaw, "red") {
					round.Red = count
				} else if strings.Contains(countRaw, "blue") {
					round.Blue = count
				} else if strings.Contains(countRaw, "green") {
					round.Green = count
				}
			}
			game.Rounds = append(game.Rounds, round)
		}
		games = append(games, game)
	}
	// part 1
	part1 := 0
	for _, game := range games {
		max := maxRound(game.Rounds)
		if max.Red <= 12 && max.Green <= 13 && max.Blue <= 14 {
			part1 += game.ID
		}
	}
	fmt.Println("part1 = ", part1)

	// part 2
	part2 := 0
	for _, game := range games {
		max := maxRound(game.Rounds)
		part2 += max.Red * max.Green * max.Blue
	}
	fmt.Println("part2 = ", part2)
}

func maxRound(rounds []Round) Round {
	out := Round{}
	for _, r := range rounds {
		if r.Red > out.Red {
			out.Red = r.Red
		}
		if r.Blue > out.Blue {
			out.Blue = r.Blue
		}
		if r.Green > out.Green {
			out.Green = r.Green
		}
	}
	return out
}
