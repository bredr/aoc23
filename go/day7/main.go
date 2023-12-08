package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	Cards        [5]int
	Score        int
	RankingScore int
}

func main() {
	b, err := os.ReadFile("day7/input")
	if err != nil {
		log.Fatal(err)
	}
	scoreMap := map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}

	str := string(b)
	lines := strings.Split(str, "\n")

	var hands []Hand
	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		score, _ := strconv.Atoi(splitLine[1])
		var cards [5]int
		for i, card := range strings.Split(splitLine[0], "") {
			cards[i] = scoreMap[card]
		}
		hands = append(hands, Hand{Cards: cards, Score: score})
	}

	scores := []struct {
		f     func(x [5]int) bool
		value int
	}{
		{isFiveOfAKind, 7},
		{isFourOfAKind, 6},
		{isFullHouse, 5},
		{isThreeOfAKind, 4},
		{isTwoPair, 3},
		{isPair, 2},
		{isHighCard, 1},
	}

	for i, hand := range hands {
		for _, score := range scores {
			if score.f(hand.Cards) {
				hands[i].RankingScore = score.value
				break
			}
		}
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].RankingScore == hands[j].RankingScore {
			for i, cardi := range hands[i].Cards {
				cardj := hands[j].Cards[i]
				if cardi != cardj {
					return cardi < cardj
				}
			}
		}
		return hands[i].RankingScore < hands[j].RankingScore
	})
	part1 := 0
	for i, hand := range hands {
		part1 += (i + 1) * hand.Score
	}
	fmt.Println("part1 = ", part1)

	scoreMapPart2 := map[string]int{"A": 13, "K": 12, "Q": 11, "J": 0, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}
	hands = []Hand{}
	for _, line := range lines {
		splitLine := strings.Split(line, " ")

		score, _ := strconv.Atoi(splitLine[1])
		var cards [5]int
		for i, card := range strings.Split(splitLine[0], "") {
			cards[i] = scoreMapPart2[card]
		}
		hands = append(hands, Hand{Cards: cards, Score: score})
	}

	for i, hand := range hands {
		hands[i].RankingScore = maxScore(hand.Cards)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].RankingScore == hands[j].RankingScore {
			for i, cardi := range hands[i].Cards {
				cardj := hands[j].Cards[i]
				if cardi != cardj {
					return cardi < cardj
				}
			}
		}
		return hands[i].RankingScore < hands[j].RankingScore
	})
	part2 := 0
	for i, hand := range hands {
		part2 += (i + 1) * hand.Score
	}
	fmt.Println("part2 = ", part2)
}

func maxScore(x [5]int) int {
	maxScore := 0
	cardValues := make([]int, 13)
	for idx := range cardValues {
		cardValues[idx] = idx + 1
	}
	scores := []struct {
		f     func(x [5]int) bool
		value int
	}{
		{isFiveOfAKind, 7},
		{isFourOfAKind, 6},
		{isFullHouse, 5},
		{isThreeOfAKind, 4},
		{isTwoPair, 3},
		{isPair, 2},
		{isHighCard, 1},
	}
	jackIdx := []int{}
	for idx, v := range x {
		if v == 0 {
			jackIdx = append(jackIdx, idx)
		}
	}
	if len(jackIdx) == 0 {
		for _, score := range scores {
			if score.f(x) && score.value > maxScore {
				maxScore = score.value
				break
			}
		}
	} else {
		np := nextProduct(cardValues, len(jackIdx))
		for {
			product := np()
			if len(product) == 0 {
				break
			}
			xcopy := [5]int{}
			copy(xcopy[:], x[:])
			for idx, v := range product {
				xcopy[jackIdx[idx]] = v
			}
			for _, score := range scores {
				if score.f(xcopy) && score.value > maxScore {
					maxScore = score.value
					break
				}
			}
		}
	}

	return maxScore
}

func isFiveOfAKind(x [5]int) bool {
	for _, v := range x {
		if v != x[0] {
			return false
		}
	}
	return true
}

func isFourOfAKind(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	for _, v := range counts {
		if v == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	if len(counts) != 2 {
		return false
	}
	for _, v := range counts {
		if v == 3 {
			return true
		}
	}
	return false
}

func isThreeOfAKind(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	for _, v := range counts {
		if v == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	pairs := 0
	for _, v := range counts {
		if v == 2 {
			pairs++
		}
	}

	return pairs == 2
}

func isPair(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	for _, v := range counts {
		if v == 2 {
			return true
		}
	}
	return false
}

func isHighCard(x [5]int) bool {
	counts := make(map[int]int)
	for _, v := range x {
		counts[v]++
	}
	return len(counts) == 5
}

func nextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}
