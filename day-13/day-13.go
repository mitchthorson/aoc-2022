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
	output := ""
	if s.Value >= 0 {
		output = fmt.Sprintf("%s%d", output, s.Value)
	}
	if s.Value == -1 && len(s.Children) == 0 {
		return "[]"
	}
	if len(s.Children) > 0 {
		output = fmt.Sprintf("%s[", output)
	}
	for i, c := range s.Children {
		if i > 0 {
			output = fmt.Sprintf("%s,", output)
		}
		output = fmt.Sprintf("%s%s", output, c)
	}
	if len(s.Children) > 0 {
		output = fmt.Sprintf("%s]", output)
	}


	// if len(s.Children) < 1 {
	// 	if s.Value == -1 {
	// 		return "[]"
	// 	}
	// 	return fmt.Sprintf("%d", s.Value)
	// }
	// output := "["
	// output = fmt.Sprintf("%s%s", output, "]")
	return output
}

func cleanValue(item string) int {
	cleanString := strings.ReplaceAll(strings.ReplaceAll(item, "[", ""), "]", "")
	if cleanString == "" {
		return -1
	}
	packetValue, err := strconv.Atoi(string(cleanString))
	utils.Check(err)
	return packetValue
}

func newPacket() *signalPacket {
	p := new(signalPacket)
	p.Value = -1
	return p
}

func (s *signalPacket) fromString(input string) *signalPacket {
	outerLen := len(input) - 1
	innerInput := input[1:outerLen]
	items := strings.Split(innerInput, ",")
	current := s
	for _, item := range items {
		listStart := strings.HasPrefix(item, "[")
		listEnd := strings.HasSuffix(item, "]")
		if listStart == true {
			for _, c := range item {
				if c == '[' {
						childPacket := newPacket()
						current.Append(childPacket)
						current = childPacket
				}

			}
		}
		val := cleanValue(item)
		if val >= 0 {
			childPacket := newPacket()
			childPacket.Value = val
			current.Append(childPacket)
		}
		if listEnd == true {
			for _, c := range item {
				if c == ']' {
					current = current.Parent
				}

			}
		}
	}

	return s
}

func compare(left, right *signalPacket) int {
	// fmt.Println("comparing", left, right)
	// actual values, not containers
	if len(left.Children) == 0 && len(right.Children) == 0 {
		fmt.Println("values", left.Value, right.Value)
		if left.Value < right.Value {
			// fmt.Println("left is smaller")
			return 1
		} else if left.Value > right.Value {
			// fmt.Println("right is smaller")
			return -1
		}
		// fmt.Println("left and right are equal")
		return 0
	}
	// @ TODO this handling is all messed up now, need to fix 
	if len(right.Children) > 0 && len(left.Children) == 0 && left.Value != -1 {
		// if left is an integer, put it into a list
		fmt.Println("left is a value, right is a list, making left a sublist")
		newVal := newPacket()
		newVal.Append(&*left)
		return compare(newVal, right)
	}
	if len(left.Children) > 0 && len(right.Children) == 0 && right.Value != -1 {
		// if right is an integer, put it into a list
		fmt.Println("right is a value, left is a list, making right a sublist")
		newVal := newPacket()
		newVal.Append(&*right)
		return compare(left, newVal)
	}
	for i, l := range left.Children {
		if i >= len(right.Children) {
			fmt.Println("right ran out of items")
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
			fmt.Println("left ran out of items")
			return 1
		}
		fmt.Println("no decision, next child")
	}
	// in case left value is empty
	// if left.Value == -1 && len(left.Children) == 0 {
	// 	fmt.Println("left is out of items")
	// 	return 1
	// }
	// if len(left.Children) > len(right.Children) {
	// 	return -1
	// } else if len(right.Children) > len(left.Children) {
	// 	return 1
	// }
	fmt.Println("hitting default tie")
	return 0
}

func parseInput(input string) [][2]*signalPacket {
	pairs := strings.Split(input, "\n\n")
	parsedPairs := make([][2]*signalPacket, len(pairs))
	// var parsedPair [2]signalPacket
	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		leftPackets := newPacket()
		leftPackets.fromString(lines[0])
		rightPackets := newPacket()
		rightPackets.fromString(lines[1])
		parsedPairs[i][0] = leftPackets
		parsedPairs[i][1] = rightPackets
	}
	return parsedPairs
}

func GetResult1(input string) int {
	pairs := parseInput(input)
	result := 0
	for i, p := range pairs {
		if i == 50 {
			fmt.Println("pair", i+1)
			fmt.Println(p[0])
			fmt.Println(p[1])
			if compare(p[0], p[1]) > 0 {
				fmt.Println("correct", i + 1)
				result += (i + 1)
			}
		}
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-13/input.txt")
	fmt.Printf("Day 13 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 13 part 2 result is:\n%d\n", GetResult2(input))
}
