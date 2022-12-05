package day05

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

func ExampleGetResult1() {
	fmt.Println("hi")
	input := utils.ReadLines("./test_input.txt")
	result := GetResult1(input)
	fmt.Println(result)
	//Output: CMZ
}
// func ExampleGetResult2() {
// 	input := utils.ReadFile("./test_input.txt")
// 	lines := strings.Split(input, "\n")
// 	result := GetResult2(lines)
// 	fmt.Println(result)
// 	//Output: 4
// }
