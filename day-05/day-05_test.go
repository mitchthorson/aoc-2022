package day05

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	lines := strings.Split(input, "\n")
	result := GetResult1(lines)
	fmt.Println(result)
	//Output: CMZ
}
func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	lines := strings.Split(input, "\n")
	result := GetResult2(lines)
	fmt.Println(result)
	//Output: MCD
}
// func ExampleGetResult2() {
// 	input := utils.ReadFile("./test_input.txt")
// 	lines := strings.Split(input, "\n")
// 	result := GetResult2(lines)
// 	fmt.Println(result)
// 	//Output: 4
// }
