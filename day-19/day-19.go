package day19

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

// I think what I need to do here is similar to the other graph problem 
// where I used the Floyd-Warshall algorithm to 
// find the shortest paths in a tree
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// but in this case there isn't a fixed distance between nodes,
// so it is going to take some creative thinking about how to prune branches


func GetResult1(input string) int {
	return 0
}

func Run() {
	input := utils.ReadFile("./day-16/input.txt")
	fmt.Printf("Day 19 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 19 part 2 result is:\n%d\n", GetResult2(input))
}
