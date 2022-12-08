package day07

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	// "strings"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult1(input)
	fmt.Println(result)
	//Output: 95437 
}
func ExampleGetResult2() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult2(input)
	fmt.Println(result)
	//Output: 24933642 
}
