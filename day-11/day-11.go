package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Monkey struct {
	id        int
	items     []int
	count     int
	operation string
	test      int
	targets   []int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("<Monkey %d items: %v count: %d>", m.id, len(m.items), m.count)
}

func (m *Monkey) throwItem(item int) {
	m.items = append(m.items, item)
}

func newMonkey(id int) *Monkey {
	m := new(Monkey)
	m.id = id
	m.count = 0
	m.items = []int{}
	m.targets = []int{}
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
					monkey.items = append(monkey.items, itemValue)
				}
			case "Operation":
				monkey.operation = strings.TrimPrefix(content, "new = old ")
			case "Test":
				testVal, err := strconv.Atoi(strings.TrimPrefix(content, "divisible by "))
				utils.Check(err)
				monkey.test = testVal
			default:
				strTarget := strings.Replace(content, "throw to monkey ", "", 1)
				intTarget, _ := strconv.Atoi(strTarget)
				monkey.targets = append(monkey.targets, intTarget)
			}
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func doRound(monkeys []*Monkey) []*Monkey {
	for _, m := range monkeys {
		for _, item := range m.items {
			// update value based on operation
			item = doOperation(m.operation, item)
			// monkey inspection count goes up
			m.count++
			// monkey gets bored
			item = item / 3
			// monkey checks test
			var targetIndex int
			if item%m.test == 0 {
				targetIndex = m.targets[0]
			} else {
				targetIndex = m.targets[1]
			}
			monkeys[targetIndex].throwItem(item)
		}
		// reset items list
		m.items = []int{}
	}
	return monkeys
}

func GetResult1(input string) int {
	monkeys := parseMonkeys(input)
	numRounds := 20
	for i := 0; i < numRounds; i++ {
		monkeys = doRound(monkeys)
	}
	counts := make([]int, len(monkeys))
	for i, m := range monkeys {
		counts[i] = m.count
	}
	sort.Ints(counts)
	top := counts[len(counts)-2:]
	result := top[0] * top[1]
	return result
}

func Run() {
	input := utils.ReadFile("./day-11/input.txt")
	fmt.Printf("Day 10 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 10 part 2 result is:\n%s\n", GetResult2(input))
}
