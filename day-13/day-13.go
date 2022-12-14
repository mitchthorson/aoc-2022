package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

// func GetResult2(input string) int {
// 	result := 0
// 	return result
// }

type signalPacket struct {
	Value int
	Children []*signalPacket
	Parent *signalPacket
}

func (s *signalPacket) Append(newPacket *signalPacket) {
	s.Children = append(s.Children, newPacket)
	newPacket.Parent = s
}

func (s *signalPacket) String() string {
	if len(s.Children) < 1 {
		if s.Value == 0 {
			return "[]"
		}
		return fmt.Sprintf("%d", s.Value)
	}
	output := "["
	for _, c := range s.Children {
		output = fmt.Sprintf("%s %s", output, c)
	}
	output = fmt.Sprintf("%s%s", output, "]")
	return output
}

func cleanValue(item rune) int {
	packetValue, err := strconv.Atoi(string(item))
	utils.Check(err)
	return packetValue
}

func (s *signalPacket) fromString(input string) *signalPacket {
	outerLen := len(input) - 1
	innerInput := input[1:outerLen]
	items := strings.ReplaceAll(innerInput, ",", "")
	current := s
	for _, item := range items {
		fmt.Println("parsing item", string(item))
		// parsing issue when i get []]
		// tries to recurse without moving up a level
		// [ is 91
		// ] is 93
		listStart := item == 91
		listEnd := item == 93
		if listStart == true {
			fmt.Println("adding child")
			// fmt.Printf("List starting up: %s\n", item)
			childPacket := new(signalPacket)
			current.Append(childPacket)
			current = childPacket
			continue
		}
		if listEnd == true {
			fmt.Println("ending child")
			fmt.Println(current)
			current = current.Parent
			continue
		}
		// this is either a value or a value and end of list
		val := cleanValue(item)
		childPacket := new(signalPacket)
		childPacket.Value = val
		current.Append(childPacket)
	}

	return s
}

func compare(left, right *signalPacket) int {
	fmt.Println("comparing", left, right)
	if left.Value > 0 && right.Value > 0 {
		if left.Value < right.Value {
			fmt.Println("left is smaller")
			return 1
		} else if left.Value > right.Value {
			fmt.Println("right is smaller")
			return -1
		}
		fmt.Println("equal")
		return 0
	}
	if left.Value > 0 {
		// if left is an integer, put it into a list
		newVal := new(signalPacket)
		newVal.Append(&*left)
		return compare(newVal, right)
	}
	if right.Value > 0 {
		// if right is an integer, put it into a list
		newVal := new(signalPacket)
		newVal.Append(&*right)
		return compare(left, newVal)
	}
	for i, l := range left.Children {
		if i >= len(right.Children) {
			fmt.Println("left ran out of items")
			return -1
		}
		c := compare(l, right.Children[i])
		if c < 0 {
			return -1
		} 
		if c > 0 {
			return 1
		}
		if i == len(left.Children) - 1 && i < len(right.Children) - 1 {
			fmt.Println("right ran out of items")
			return 1
		}
	}
	// if len(left.Children) > len(right.Children) {
	// 	return -1
	// } else if len(right.Children) > len(left.Children) {
	// 	return 1
	// }
	return 0
}

func parseInput(input string) [][2]*signalPacket {
	pairs := strings.Split(input, "\n\n")
	parsedPairs := make([][2]*signalPacket, len(pairs))
	// var parsedPair [2]signalPacket
	for i, pair := range pairs[:1] {
		lines := strings.Split(pair, "\n")
		leftPackets := new(signalPacket)
		leftPackets.fromString(lines[0])
		rightPackets := new(signalPacket)
		rightPackets.fromString(lines[1])
		parsedPairs[i][0] = leftPackets
		parsedPairs[i][1] = rightPackets
	}
	return parsedPairs
}

func GetResult1(input string) int {
	pairs := parseInput(input)
	result := 0
	for i, p := range pairs[:1] {
		fmt.Println("pair", i+1)
		fmt.Println(p[0])
		fmt.Println(p[1])
		// if compare(p[0], p[1]) > 0 {
		// 	fmt.Println("correct", i + 1)
		// 	result += (i + 1)
		// }
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-13/input.txt")
	fmt.Printf("Day 13 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 13 part 2 result is:\n%d\n", GetResult2(input))
}
