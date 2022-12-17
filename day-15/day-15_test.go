package day15

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult1(input, 10)
	fmt.Println(result)
	//Output: 26
}
func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult2(input, 20)
	fmt.Println(result)
	//Output: 56000011
}
