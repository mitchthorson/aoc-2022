package day06

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	// "strings"
)

type SignalBuffer [4]rune

func nextSignalItem(signal *SignalBuffer, item rune) *SignalBuffer {
	signalLength := len(signal)
	for i := 1; i < len(signal); i++ {
		signal[i - 1] = signal[i]
	}
	signal[signalLength - 1] = item
	return signal
}

func checkSignal(signal *SignalBuffer) bool {
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


func GetResult1(input string) int {
	result := 0
	var signal SignalBuffer
	for i, letter := range input {
		nextSignalItem(&signal, letter)
		if checkSignal(&signal) {
			result = i + 1
			break
		}
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-06/input.txt")
	fmt.Printf("Day 06 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 06 part 2 result is:\n%s\n", GetResult2(lines))
}
