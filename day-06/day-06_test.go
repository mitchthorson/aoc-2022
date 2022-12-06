package day06

import (
	"fmt"
	// "github.com/mitchthorson/aoc-2022/utils"
	// "strings"
)

func ExampleGetResult1() {
	// input := utils.ReadFile("./test_input.txt")
	result := ""
	testInputs := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	for _, t := range testInputs {
		result = fmt.Sprintf("%s %d", result, GetResult1(t))
	}
	fmt.Println(result)
	//Output: 7 5 6 10 11
}
func ExampleGetResult2() {
	result := ""
	testInputs := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	for _, t := range testInputs {
		result = fmt.Sprintf("%s %d", result, GetResult2(t))
	}
	fmt.Println(result)
	//Output: 19 23 23 29 26
}
