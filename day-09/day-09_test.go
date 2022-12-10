package day09

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	// "strings"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult1(input)
	fmt.Println(result)
	//Output: 13
}

func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult2(input)
	input2 := utils.ReadFile("./test_2_input.txt")
	result2 := GetResult2(input2)
	fmt.Println(fmt.Sprintf("%d %d", result, result2))
	//Output: 1 36
}
