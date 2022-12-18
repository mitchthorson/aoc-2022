package day16

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

func ExampleGetResult1() {
	input := utils.ReadFile("./test_input.txt")
	result := GetResult1(input)
	fmt.Println(result)
	//Output: 0
}
