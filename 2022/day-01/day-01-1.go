package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	trimmedInput := strings.Trim(string(input), "\n")
	trimmedInput = strings.Trim(trimmedInput, " ")
	perElf := strings.Split(trimmedInput, "\n\n")

	var perElfCals []int
	for i := 0; i < len(perElf); i++ {
		cals := strings.Split(perElf[i], "\n")
		calsTotal := 0
		for j := 0; j < len(cals); j++ {
			intval, err := strconv.Atoi(cals[j])
			if err != nil {
				fmt.Printf("Conversion issue at index %d\n", i)
				panic(err)
			}
			calsTotal += intval
		}
		perElfCals = append(perElfCals, calsTotal)
	}

	perElfCalsSorted := perElfCals[:]
	sort.Ints(perElfCalsSorted)
	fmt.Println(perElfCalsSorted[len(perElfCalsSorted)-1])
}
