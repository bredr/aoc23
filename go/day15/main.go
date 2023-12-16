package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
