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
	"github.com/mitchthorson/aoc-2022/day-15"
	"github.com/mitchthorson/aoc-2022/day-16"
	"github.com/mitchthorson/aoc-2022/day-17"
)

func main() {
	fmt.Println("Welcome to AOC 2022!")
	day := os.Args[1]
	fmt.Printf("Running day %s\n\n", day)
	if day == "1" {
		day01.Run()
		return
	}
	if day == "2" {
		day02.Run()
		return
	}
	if day == "3" {
		day03.Run()
		return
	}
	if day == "4" {
		day04.Run()
		return
	}
	if day == "5" {
		day05.Run()
		return
	}
	if day == "6" {
		day06.Run()
		return
	}
	if day == "7" {
		day07.Run()
		return
	}
	if day == "8" {
		day08.Run()
		return
	}
	if day == "9" {
		day09.Run()
		return
	}
	if day == "10" {
		day10.Run()
		return
	}
	if day == "11" {
		day11.Run()
		return
	}
	if day == "12" {
		day12.Run()
		return
	}
	if day == "13" {
		day13.Run()
		return
	}
	if day == "14" {
		day14.Run()
		return
	}
	if day == "15" {
		day15.Run()
		return
	}
	if day == "16" {
		day16.Run()
		return
	}
	if day == "17" {
		day17.Run()
		return
	}
	fmt.Println("Day not setup", day)
}
