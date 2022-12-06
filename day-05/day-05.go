package day05

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"sort"
	"strings"
)

type Instruction struct{ amount, origin, destination int }

type CrateStacks struct {
	stacks map[int]*Stack
}

func (cs *CrateStacks) printStacks() {
	for k, v := range cs.stacks {
		fmt.Println(k)
		current := v.head
		for current.next != nil {
			fmt.Printf("-%c\n", current.content)
			current = current.next
		}
	}
}

func (cs *CrateStacks) getStacks() []int {
	stackIds := []int{}
	for k := range cs.stacks {
		stackIds = append(stackIds, k)
	}
	sort.Ints(stackIds)
	return stackIds
}

func (cs *CrateStacks) insertToStack(stackId int, content byte) *CrateStacks {
	if cs.stacks == nil {
		cs.stacks = map[int]*Stack{}
	}
	stack, ok := cs.stacks[stackId]
	if !ok {
		stack = &Stack{}
		cs.stacks[stackId] = stack
	}
	c := new(Crate)
	c.content = content
	stack.insertItem(c)
	return cs
}

func (cs *CrateStacks) moveCrates(instruction *Instruction) *CrateStacks {
	for i := 0; i < instruction.amount; i++ {
		crate := cs.stacks[instruction.origin].takeItem()
		cs.stacks[instruction.destination].insertItem(crate)
	}
	return cs
}

func (cs *CrateStacks) moveCratesGrouped(instruction *Instruction) *CrateStacks {
	var movedCrates []*Crate
	for i := 0; i < instruction.amount; i++ {
		crate := cs.stacks[instruction.origin].takeItem()
		movedCrates = append(movedCrates, crate)
	}
	// put them into the new stack in reverse order
	for i := len(movedCrates) - 1; i >= 0; i-- {
		cs.stacks[instruction.destination].insertItem(movedCrates[i])
	}
	return cs
}

type Crate struct {
	content byte
	next    *Crate
}

type Stack struct {
	head *Crate
}

func (s *Stack) insertItem(item *Crate) {
	item.next = s.head
	s.head = item
}

func (s *Stack) takeItem() *Crate {
	c := s.head
	s.head = s.head.next
	return c
}

func parseStacks(rawStacks []string) *CrateStacks {
	stacks := new(CrateStacks)
	// inputColumns := []byte{}
	for i := len(rawStacks) - 1; i >= 0; i-- {
		line := rawStacks[i]
		for i := 1; i < len(line); i += 4 {
			val := line[i]
			stackIndex := i/4 + 1
			if val == ' ' {
				continue
			}
			stacks.insertToStack(stackIndex, val)
		}
	}
	return stacks
}

func parseInstructions(rawInstructions []string) []*Instruction {
	instructionList := []*Instruction{}
	for _, instruction := range rawInstructions {
		if instruction == "" {
			continue
		}
		inst := Instruction{}
		fmt.Sscanf(instruction, "move %d from %d to %d", &inst.amount, &inst.origin, &inst.destination)
		instructionList = append(instructionList, &inst)
	}
	return instructionList
}

func parseInput(rawInput []string) (*CrateStacks, []*Instruction) {
	outputs := map[int][]string{}
	outputIdx := 0
	for _, line := range rawInput {
		if line == "" {
			outputIdx++
		}
		outputs[outputIdx] = append(outputs[outputIdx], line)
	}
	return parseStacks(outputs[0]), parseInstructions(outputs[1])
}

func GetResult1(input []string) string {
	stacks, instructions := parseInput(input)
	for _, instruction := range instructions {
		stacks.moveCrates(instruction)
	}
	result := ""
	for _, stack := range stacks.getStacks() {
		result = fmt.Sprintf("%s%c", result, stacks.stacks[stack].head.content)
	}
	return result
}

func GetResult2(input []string) string {
	stacks, instructions := parseInput(input)
	for _, instruction := range instructions {
		stacks.moveCratesGrouped(instruction)
	}
	result := ""
	for _, stack := range stacks.getStacks() {
		result = fmt.Sprintf("%s%c", result, stacks.stacks[stack].head.content)
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-05/input.txt")
	lines := strings.Split(input, "\n")
	fmt.Printf("Day 05 part 1 result is:\n%s\n", GetResult1(lines))
	fmt.Printf("Day 05 part 2 result is:\n%s\n", GetResult2(lines))
}
