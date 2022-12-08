package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	games := strings.Split(string(input), "\n")

	rock := 1
	paper := 2
	scisors := 3
	win := 6
	draw := 3
	loss := 0

	inputMap := map[string]int{
		"A": rock,
		"B": paper,
		"C": scisors,
		"X": rock,
		"Y": paper,
		"Z": scisors,
	}

	totalPts := 0
	for _, v := range games {
		r := strings.Fields(v)
		pts := inputMap[r[1]]

		if inputMap[r[0]] == inputMap[r[1]] {
			pts += draw
		} else if (inputMap[r[0]] == rock && inputMap[r[1]] == paper) ||
			(inputMap[r[0]] == paper && inputMap[r[1]] == scisors) ||
			(inputMap[r[0]] == scisors && inputMap[r[1]] == rock) {
			pts += win
		} else {
			pts += loss
		}

		totalPts += pts
	}

	fmt.Println(totalPts)
}
