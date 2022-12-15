package main

import (
	"fmt"
	"os"
	"github.com/mitchthorson/aoc-2022/day-01"
	"github.com/mitchthorson/aoc-2022/day-02"
	"github.com/mitchthorson/aoc-2022/day-03"
	"github.com/mitchthorson/aoc-2022/day-04"
	"github.com/mitchthorson/aoc-2022/day-05"
	"github.com/mitchthorson/aoc-2022/day-06"
	"github.com/mitchthorson/aoc-2022/day-07"
	"github.com/mitchthorson/aoc-2022/day-08"
	"github.com/mitchthorson/aoc-2022/day-09"
	"github.com/mitchthorson/aoc-2022/day-10"
	"github.com/mitchthorson/aoc-2022/day-11"
	"github.com/mitchthorson/aoc-2022/day-12"
	"github.com/mitchthorson/aoc-2022/day-13"
	"github.com/mitchthorson/aoc-2022/day-14"
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
	if day == "4" {
		day04.Run()
	}
	if day == "5" {
		day05.Run()
	}
	if day == "6" {
		day06.Run()
	}
	if day == "7" {
		day07.Run()
	}
	if day == "8" {
		day08.Run()
	}
	if day == "9" {
		day09.Run()
	}
	if day == "10" {
		day10.Run()
	}
	if day == "11" {
		day11.Run()
	}
	if day == "12" {
		day12.Run()
	}
	if day == "13" {
		day13.Run()
	}
	if day == "14" {
		day14.Run()
	}
}
