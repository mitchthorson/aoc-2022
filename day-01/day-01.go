package day01

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"sort"
	"strconv"
	"strings"
)

func getElfCalories(elf []int) int {
	total := 0
	for _, v := range elf {
		total = total + v
	}
	return total
}

func getElves(rawElves []string) []int {
	var elves []int
	for _, elf := range rawElves {
		str_vals := strings.Split(elf, "\n")
		var int_vals []int
		for _, v := range str_vals {
			int_val, err := strconv.Atoi(v)
			utils.Check(err)
			int_vals = append(int_vals, int_val)
		}
		elves = append(elves, getElfCalories(int_vals))
	}
	return elves
}

func getResult1(input string) int {
	rawElves := strings.Split(input, "\n\n")
	elves := getElves(rawElves)
	result := 0
	for _, v := range elves {
		if v > result {
			result = v
		}
	}
	return result
}

func getResult2(input string) int {
	rawElves := strings.Split(input, "\n\n")
	elves := getElves(rawElves)
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return elves[0] + elves[1] + elves[2]
}

func RunTest() {
	input := utils.ReadTestInput(1)
	fmt.Printf("Day 01 test part 1 result is:\n%d\n", getResult1(input))
	fmt.Printf("\nDay 01 test part 2 result is:\n%d\n", getResult2(input))
}

func Run() {
	input := utils.ReadInput(1)
	fmt.Printf("\nDay 01 part 1 result is:\n%d\n", getResult1(input))
	fmt.Printf("\nDay 01 part 2 result is:\n%d\n", getResult2(input))
}
