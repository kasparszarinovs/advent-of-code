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

	inputMap := map[string]map[string]int{
		"X": {
			"state": loss,
			"A":     scisors,
			"B":     rock,
			"C":     paper,
		},
		"Y": {
			"state": draw,
			"A":     rock,
			"B":     paper,
			"C":     scisors,
		},
		"Z": {
			"state": win,
			"A":     paper,
			"B":     scisors,
			"C":     rock,
		},
	}

	totalPts := 0
	for _, v := range games {
		r := strings.Fields(v)
		pts := inputMap[r[1]]["state"] + inputMap[r[1]][r[0]]
		totalPts += pts
	}

	fmt.Println(totalPts)
}
