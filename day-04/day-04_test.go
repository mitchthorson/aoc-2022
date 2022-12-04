package day04

import (
	"fmt"
	"strings"
	"github.com/mitchthorson/aoc-2022/utils"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	lines := strings.Split(input, "\n")
	result := GetResult1(lines)
	fmt.Println(result)
	//Output: 2
}
func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	lines := strings.Split(input, "\n")
	result := GetResult2(lines)
	fmt.Println(result)
	//Output: 4
}
