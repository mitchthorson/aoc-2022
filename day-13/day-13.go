package day13

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type packet struct {
	Value int
	Children []*packet
	Parent *packet
}

// proud of thinking to implement the interface for Sort()
func (s *packet) Len() int {
	return len(s.Children)
}
func (s *packet) Less(i, j int) bool {
	c := compare(s.Children[i], s.Children[j]) 
	if c >= 0 {
		return true
	}
	return false
}
func (s *packet) Swap(i, j int) {
	iPacket := s.Children[i]
	jPacket := s.Children[j]
	s.Children[i] = jPacket
	s.Children[j] = iPacket
}

func (s *packet) GetDividerPackets() (*packet, *packet) {
	dividerPacker1 := "[[2]]"
	dividerPacker2 := "[[6]]"
	dividerPacket1 := newPacket()
	dividerPacket1.FromString(dividerPacker1)
	dividerPacket2 := newPacket()
	dividerPacket2.FromString(dividerPacker2)
	return dividerPacket1, dividerPacket2
}

func (s *packet) AddDividerPackets() {
	k1,k2 := s.GetDividerPackets()
	s.Append(k1)
	s.Append(k2)
}

func (s *packet) Append(newPacket *packet) {
	s.Children = append(s.Children, newPacket)
	newPacket.Parent = s
}

func (s *packet) String() string {
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
	return output
}

func (s *packet) IsValue() bool {
	if s.Value > -1 && len(s.Children) == 0 {
		return true
	}
	return false
}

func (s *packet) FindDividerPackets() (int, int) {
	keyA, keyB := s.GetDividerPackets()
	indexA, indexB := 0, 0
	for i, p := range s.Children {
		if p.String() == keyA.String() {
			indexA = i
		}
		if p.String() == keyB.String() {
			indexB = i
		}
	}

	return indexA + 1, indexB + 1

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

func newPacket() *packet {
	p := new(packet)
	p.Value = -1
	return p
}

func (s *packet) FromString(input string) *packet {
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

func compare(left, right *packet) int {
	leftIsV, rightIsV := left.IsValue(), right.IsValue()
	// if both are values, compare directly
	if leftIsV && rightIsV {
		if left.Value < right.Value {
			return 1
		}
		if left.Value > right.Value {
			return -1
		}
		return 0
	}
	// if left is a value and right is a list, put left in a sublist
	if leftIsV && !rightIsV {
		subList := newPacket()
		subList.Append(left)
		return compare(subList, right)
	}
	// if left is a list and right is a value, put right in a sublist
	if !leftIsV && rightIsV {
		subList := newPacket()
		subList.Append(right)
		return compare(left, subList)
	}

	// if we reach here, both items are lists
	for i, leftChild := range left.Children {
		// if index is out of right range
		if i >= len(right.Children) {
			return -1
		}
		// compare children directly
		c := compare(leftChild, right.Children[i])

		// if children are not equal return the comparison
		if c != 0 {
			return c
		}
		// if they are equal keep looping

	}
	
	// if we reach here, all compared children are equal, lets see if one is longer
	lenLeft, lenRight := len(left.Children), len(right.Children)
	// left is longer than right
	if lenLeft > lenRight {
		return -1
	}
	// right is longer than left
	if lenLeft < lenRight {
		return 1
	}

	// everything is equal
	return 0
}

func parsePairs(input string) [][2]*packet {
	pairs := strings.Split(input, "\n\n")
	parsedPairs := make([][2]*packet, len(pairs))
	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		leftPackets := newPacket()
		leftPackets.FromString(lines[0])
		rightPackets := newPacket()
		rightPackets.FromString(lines[1])
		parsedPairs[i][0] = leftPackets
		parsedPairs[i][1] = rightPackets
	}
	return parsedPairs
}

func parsePackets(input string) *packet {
	cleaned := strings.ReplaceAll(input, "\n\n", "\n")
	packets := strings.Split(cleaned, "\n")
	parsedPackets := make([]*packet, len(packets))
	for i, packet := range packets {
		p := newPacket()
		p.FromString(packet)
		parsedPackets[i] = p
	}
	all := newPacket()
	all.Children = parsedPackets
	return all
}

func GetResult1(input string) int {
	pairs := parsePairs(input)
	result := 0
	for i, p := range pairs {
		if compare(p[0], p[1]) > 0 {
			result += (i + 1)
		}
	}
	return result
}

func GetResult2(input string) int {
	packets := parsePackets(input)
	packets.AddDividerPackets()
	sort.Sort(packets)
	p1, p2 := packets.FindDividerPackets()
	return p1 * p2
}

func Run() {
	input := utils.ReadFile("./day-13/input.txt")
	fmt.Printf("Day 13 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 13 part 2 result is:\n%d\n", GetResult2(input))
}
