package main

import (
	"fmt"
	"os"
	"github.com/mitchthorson/aoc-2022/day-01"
)

func main() {
	fmt.Println("Welcome to AOC 2022!")
	day := os.Args[1]
	fmt.Printf("Running day %s\n", day)
	if day == "1" {
		day01.RunTest()
		day01.Run()
	}
}
