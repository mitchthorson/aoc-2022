package day02

import (
	"fmt"
	"strings"
	"github.com/mitchthorson/aoc-2022/utils"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	rounds := strings.Split(input, "\n")
	result := GetResult1(rounds)
	fmt.Println(result)
	// Output: 15
}

func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	rounds := strings.Split(input, "\n")
	result := GetResult2(rounds)
	fmt.Println(result)
	// Output: 12
}

// func TestGetResult2(t *testing.T) {
// 	input := utils.ReadFile("./test_input.txt")
// 	expectedResult := 45000
// 	result := getResult2(input)
// 	if result != expectedResult {
// 		t.Errorf("Wanted %d, got %d", expectedResult, result)
// 	}
// }
