package day06

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

func checkSignal(signal string) bool {
	signalRunes := map[rune]interface{}{}
	for _, r := range signal {
		_, ok := signalRunes[r]
		if ok || r == 0 {
			return false
		}
		signalRunes[r] = 1
	}
	return true
}

func findStart(signal string, length int) int {
	left := 0
	right := length
	for range signal {
		signalBuffer := signal[left:right]
		if checkSignal(signalBuffer) {
			return right
		}
		left++
		right++
	}
	return 0
}

func GetResult1(input string) int {
	return findStart(input, 4)
}
func GetResult2(input string) int {
	return findStart(input, 14)
}

func Run() {
	input := utils.ReadFile("./day-06/input.txt")
	fmt.Printf("Day 06 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 06 part 2 result is:\n%d\n", GetResult2(input))
}
