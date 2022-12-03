package main

import (
	"fmt"
	"os"
	"github.com/mitchthorson/aoc-2022/day-01"
	"github.com/mitchthorson/aoc-2022/day-02"
	"github.com/mitchthorson/aoc-2022/day-03"
)

func main() {
	fmt.Println("Welcome to AOC 2022!")
	day := os.Args[1]
	fmt.Printf("Running day %s\n\n", day)
	if day == "1" {
		day01.Run()
	}
	if day == "2" {
		day02.Run()
	}
	if day == "3" {
		day03.Run()
	}
}
