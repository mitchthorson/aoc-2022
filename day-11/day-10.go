package day11

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Monkey struct {
	items     []int
	operation string
	test      int
	targets   []int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("<Monkey items:%v operation:%s test:%d targets:%v>", m.items, m.operation, m.test, m.targets)
}

func newMonkey() *Monkey {
	m := new(Monkey)
	m.items = []int{}
	m.targets = []int{}
	return m
}

func doOperation(operationText string, inputVal int) int {
	operationFields := strings.Fields(operationText)
	fmt.Println(operationFields)
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

func parseMonkeys(input string) {
	monkeyStrs := strings.Split(input, "\n\n")
	var monkeys []*Monkey
	for _, m := range monkeyStrs {
		monkey := newMonkey()
		for i, line := range strings.Split(m, "\n") {
			// skip first line of each monkey block
			if i == 0 {
				continue
			}
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
	fmt.Println(monkeys)
}

func GetResult1(input string) int {
	parseMonkeys(input)
	result := 0
	return result
}

func Run() {
	input := utils.ReadFile("./day-10/input.txt")
	fmt.Printf("Day 10 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 10 part 2 result is:\n%s\n", GetResult2(input))
}
