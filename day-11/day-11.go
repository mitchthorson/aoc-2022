package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Monkey struct {
	Id        int
	Items     []int
	Count     int
	Operation string
	Test      int
	Targets   []int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("<Monkey %d items: %v count: %d>", m.Id, len(m.Items), m.Count)
}

func (m *Monkey) throwItem(item int) {
	m.Items = append(m.Items, item)
}

func newMonkey(id int) *Monkey {
	m := new(Monkey)
	m.Id = id
	m.Count = 0
	m.Items = []int{}
	m.Targets = []int{}
	return m
}

func doOperation(operationText string, inputVal int) int {
	operationFields := strings.Fields(operationText)
	var operationVal int
	if operationFields[1] == "old" {
		operationVal = inputVal
	} else {
		parsedVal, err := strconv.Atoi(operationFields[1])
		utils.Check(err)
		operationVal = parsedVal
	}
	var result int
	switch operationFields[0] {
	case "+":
		result = inputVal + operationVal
	case "*":
		result = inputVal * operationVal
	default:
		panic("Operation not implemented")
	}
	return result
}

func parseLinePrefix(inputLine string) (string, string) {
	lSplit := strings.Split(inputLine, ":")
	if len(lSplit) != 2 {
		panic(fmt.Sprintf("Input not parsed correctly: %s", inputLine))
	}
	return strings.Trim(lSplit[0], " "), strings.Trim(lSplit[1], " ")
}

func parseMonkeys(input string) []*Monkey {
	monkeyStrs := strings.Split(input, "\n\n")
	var monkeys []*Monkey
	for i, m := range monkeyStrs {
		monkey := newMonkey(i)
		// skip first line of each monkey block
		for _, line := range strings.Split(m, "\n")[1:] {
			prefix, content := parseLinePrefix(line)
			switch prefix {
			case "Starting items":
				items := strings.Split(content, ", ")
				for _, item := range items {
					itemValue, _ := strconv.Atoi(item)
					monkey.Items = append(monkey.Items, itemValue)
				}
			case "Operation":
				monkey.Operation = strings.TrimPrefix(content, "new = old ")
			case "Test":
				testVal, err := strconv.Atoi(strings.TrimPrefix(content, "divisible by "))
				utils.Check(err)
				monkey.Test = testVal
			default:
				strTarget := strings.Replace(content, "throw to monkey ", "", 1)
				intTarget, _ := strconv.Atoi(strTarget)
				monkey.Targets = append(monkey.Targets, intTarget)
			}
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func doRound(monkeys []*Monkey, worryReduction func(w int) int) []*Monkey {
	for _, m := range monkeys {
		for _, item := range m.Items {
			// update value based on operation
			item = doOperation(m.Operation, item)
			// monkey inspection count goes up
			m.Count++
			// monkey gets bored
			item = worryReduction(item)
			// monkey checks test
			var targetIndex int
			if item%m.Test == 0 {
				targetIndex = m.Targets[0]
			} else {
				targetIndex = m.Targets[1]
			}
			monkeys[targetIndex].throwItem(item)
		}
		// reset items list
		m.Items = []int{}
	}
	return monkeys
}

func monkeyBusiness(monkeys []*Monkey, numRounds int, stressReduction func(int) int) int {
	for i := 0; i < numRounds; i++ {
		monkeys = doRound(monkeys, stressReduction)
	}
	counts := make([]int, len(monkeys))
	for i, m := range monkeys {
		counts[i] = m.Count
	}
	sort.Ints(counts)
	top := counts[len(counts)-2:]
	return top[0] * top[1]
}

func GetResult1(input string) int {
	monkeys := parseMonkeys(input)
	numRounds := 20
	stressReduction := func(w int) int {
		return w / 3
	}
	return monkeyBusiness(monkeys, numRounds, stressReduction)
}

func GetResult2(input string) int {
	monkeys := parseMonkeys(input)
	numRounds := 10000
	stressLimit := 1
	for _,m := range monkeys {
		stressLimit *= m.Test
	}
	stressReduction := func(w int) int {
		return w % stressLimit
	}
	return monkeyBusiness(monkeys, numRounds, stressReduction)
}

func Run() {
	input := utils.ReadFile("./day-11/input.txt")
	fmt.Printf("Day 11 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 11 part 2 result is:\n%d\n", GetResult2(input))
}
