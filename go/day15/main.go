package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func main() {
	b, err := os.ReadFile("day15/input")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)

	part1 := 0
	for _, step := range strings.Split(str, ",") {
		part1 += hash(step)
	}
	fmt.Println("part1 = ", part1)

	boxes := make(map[int]*linkedhashmap.Map)
	for _, step := range strings.Split(str, ",") {
		instruction := getInstruction(step)
		if _, ok := boxes[instruction.Box]; !ok {
			boxes[instruction.Box] = linkedhashmap.New()
		}
		if instruction.FocalLength != nil {
			boxes[instruction.Box].Put(instruction.Label, *instruction.FocalLength)
		} else {
			boxes[instruction.Box].Remove(instruction.Label)
		}
	}

	part2 := 0
	for i := 0; i < 256; i++ {
		if lenses, ok := boxes[i]; ok {
			for j, focalLength := range lenses.Values() {
				part2 += (i + 1) * (j + 1) * focalLength.(int)
			}
		}
	}
	fmt.Println("part2 = ", part2)

}

func hash(x string) int {
	currentValue := 0
	for _, char := range x {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

type Instruction struct {
	Box         int
	Label       string
	FocalLength *int
}

func getInstruction(x string) Instruction {
	lettersRe := regexp.MustCompile(`[a-z]+`)
	labelString := lettersRe.FindString(x)
	box := hash(labelString)
	if strings.ContainsRune(x, '-') {
		return Instruction{box, labelString, nil}
	}
	numberRe := regexp.MustCompile(`\d+`)
	rawNumber := numberRe.FindString(x)
	focalLength, err := strconv.Atoi(rawNumber)
	if err != nil {
		panic(err)
	}
	return Instruction{box, labelString, &focalLength}
}
