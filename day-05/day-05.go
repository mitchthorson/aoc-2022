package day05

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

type Instruction struct{ amount, origin, destination int }

type Crate struct {
	content string
	next *Crate
}

type Stack struct {
	head *Crate
}

func (s *Stack) insertItem(item string) {
	c := new(Crate)
	c.content = item
	c.next = s.head
	s.head = c
}

func (s *Stack) appendItem(item string) *Stack {
	c := new(Crate)
	c.content = item
	if s.head == nil {
		s.head = c
		return s
	}
	current := s.head
	for current != nil {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = c
	return s
}

func (s *Stack) takeItem() *Crate {
	s.head = s.head.next
	return s.head
}

func parseStacks(rawStacks []string) *map[int]*Stack {
	stacks := map[int]*Stack{}
	// inputColumns := []string{}
	for _, line := range rawStacks {
		// this fields logic has one big problem, which is the first field in a line is not always aligned with column 1
		// need to somehow count the whitespace as columns
		// fields := strings.Fields(line)
		for i := 0; i < len(line); i += 4 {
			fmt.Println(line[i:i+3])
		}
		// for fieldIndex, f := range fields {
			// if strings.Contains(f, "[") {
			// 	var item string
			// 	fmt.Println(f)
			// 	fmt.Sscanf(f, "[%s]", &item)
			// 	fmt.Println(item)
			// 	if val, ok := stacks[fieldIndex + 1]; ok {
			// 		val.appendItem(item)
			// 	} else {
			// 		newStack := new(Stack)
			// 		newStack.appendItem(item)
			// 		stacks[fieldIndex + 1] = newStack
			// 	}
			// }
		// }
	}
	return &stacks
}

func parseInstructions(rawInstructions []string) []*Instruction {
	instructionList := []*Instruction{}
	for _, instruction := range rawInstructions {
		var inst Instruction
		fmt.Println(instruction)
		_, err := fmt.Sscanf(instruction, "move %d from %d to %d", &inst.amount, &inst.origin, &inst.destination)
		utils.Check(err)
		instructionList = append(instructionList, &inst)
	}
	return instructionList
}

func parseInput(rawInput []string) ([]string, []string) {
	outputs := map[int][]string{}
	outputIdx := 0
	for _, line := range rawInput {
		if line == "" {
			outputIdx++
		}
		outputs[outputIdx] = append(outputs[outputIdx], line)
	}
	return outputs[0], outputs[1]
}

func GetResult1(input []string) string {
	stackInput, instructionInput := parseInput(input)
	fmt.Println(len(stackInput))
	fmt.Println(len(instructionInput))
	result := ""
	instructions := parseInstructions(instructionInput)
	stacks := parseStacks(stackInput)
	fmt.Println(instructions, stacks)
	for k,v := range *stacks {
		fmt.Println(k, v.head.content)
	}
	return result
}
func Run() {
	input := utils.ReadLines("./day-05/test_input.txt")
	fmt.Printf("Day 05 part 1 result is:\n%s\n", GetResult1(input))
	// fmt.Printf("Day 04 part 2 result is:\n%d\n", GetResult2(inputAssignments))
}
