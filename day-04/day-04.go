package day04

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strconv"
	"strings"
)

type Assignment struct { start, stop int }

func (a *Assignment) overlaps(otherAssignment *Assignment) bool {
	if otherAssignment.start <= a.stop && otherAssignment.stop >= a.start {
		return true
	}
	return false
}

func (a *Assignment) contains(otherAssignment *Assignment) bool {
	if otherAssignment.start >= a.start && otherAssignment.stop <= a.stop {
		return true
	}
	return false
}

func (a *Assignment) fromStr(inputStr string) *Assignment {
	startstop := strings.Split(inputStr, "-")
	start, err := strconv.Atoi(startstop[0])
	utils.Check(err)
	stop, err := strconv.Atoi(startstop[1])
	utils.Check(err)
	a.start = start
	a.stop = stop
	return a
}

func getPair(a string) *[2]*Assignment {
	var assignmentPair [2]*Assignment
	intputPairs := strings.Split(a, ",")
	for i, p := range intputPairs {
		assignment := Assignment{}
		assignment.fromStr(p)
		assignmentPair[i] = &assignment
	}
	return &assignmentPair
}

func GetResult1(inputAssignments []string) int {
	result := 0
	for _, a := range inputAssignments {
		assignmentPair := getPair(a)
		if assignmentPair[0].contains(assignmentPair[1]) || assignmentPair[1].contains(assignmentPair[0]) {
			result++
		}

	}
	return result
}

func GetResult2(inputAssignments []string) int {
	result := 0
	for _, a := range inputAssignments {
		assignmentPair := getPair(a)
		if assignmentPair[0].overlaps(assignmentPair[1]) {
			result++
		}

	}
	return result
}

func Run() {
	input := utils.ReadInput(4)
	inputAssignments := strings.Split(input, "\n")
	fmt.Printf("Day 04 part 1 result is:\n%d\n", GetResult1(inputAssignments))
	fmt.Printf("Day 04 part 2 result is:\n%d\n", GetResult2(inputAssignments))
}
