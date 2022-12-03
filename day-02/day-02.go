package day02

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

func GetResult1(rounds []string) int {
	result := 0
	for _, round := range rounds {
		fields := strings.Fields(round)
		opponent := int(fields[0][0] - 'A')
		player := int(fields[1][0] - 'X')
		// add score for play
		result += player + 1
		switch {
		// winning case 
		case player == (opponent+1)%3:
			result += 6
		// draw case
		case player == opponent:
			result += 3
		}
	}
	return result
}

func GetResult2(rounds []string) int {
	result := 0
	for _, round := range rounds {
		fields := strings.Fields(round)
		opponent := int(fields[0][0] - 'A')
		switch fields[1] {
		case "X":
			result += (opponent+2)%3 + 1
		case "Y":
			result += opponent + 1 + 3
		case "Z":
			result += (opponent+1)%3 + 1 + 6

		}
	}
	return result
}

func Run() {
	input := utils.ReadInput(2)
	rounds := strings.Split(input, "\n")
	fmt.Printf("Day 02 part 1 result is:\n%d\n", GetResult1(rounds))
	fmt.Printf("\nDay 02 part 2 result is:\n%d\n", GetResult2(rounds))
}
