package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"unicode"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := strings.Split(string(input), "\n")
	groupCount := int(math.Ceil(float64(len(rucksacks) / 3)))

	scores := map[int32]int{}
	for r, i := 'a', 1; r <= 'z'; r, i = r+1, i+1 {
		scores[r] = i
		scores[unicode.ToUpper(r)] = i + 26
	}

	totalPriority := 0

	for i := 0; i < groupCount; i++ {
		groupRucksacks := rucksacks[i*3 : (i+1)*3]

		var s strings.Builder
		for _, c := range groupRucksacks[0] {
			if !strings.ContainsRune(s.String(), c) {
				s.WriteRune(c)
			}
		}

		var badge int32
		for _, c := range s.String() {
			if strings.ContainsRune(groupRucksacks[1], c) && strings.ContainsRune(groupRucksacks[2], c) {
				badge = c
				break
			}
		}

		totalPriority += scores[badge]
	}

	fmt.Println(totalPriority)
}
