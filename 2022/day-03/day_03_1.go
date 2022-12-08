package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := strings.Split(string(input), "\n")

	scores := map[string]int{}
	for r, i := 'a', 1; r <= 'z'; r, i = r+1, i+1 {
		scores[string(r)] = i
		scores[string(unicode.ToUpper(r))] = i + 26
	}

	totalPriority := 0
	for _, v := range rucksacks {
		compartmentSize := len(v) / 2
		compartment1 := v[:compartmentSize]
		compartment2 := v[compartmentSize:]

		// Dedupe
		var s strings.Builder
		for _, c := range compartment1 {
			if !strings.ContainsRune(s.String(), c) {
				s.WriteRune(c)
			}
		}

		rucksackPriority := 0
		for _, c := range s.String() {
			if strings.ContainsRune(compartment2, c) {
				rucksackPriority += scores[string(c)]
			}
		}

		totalPriority += rucksackPriority
	}

	fmt.Println(totalPriority)
}
