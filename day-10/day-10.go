package day10

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strconv"
	"strings"
)

type CPU struct {
	X       int
	program []string
}

func (c *CPU) tick() {
	if len(c.program) < 1 {
		return
	}
	instruction := c.program[0]
	c.program = c.program[1:]
	switch instruction {
	case "noop":
		return
	case "addx":
		return
	default:
		val, err := strconv.Atoi(instruction)
		utils.Check(err)
		c.X += val
	}
}

type CRT struct {
	w, h, i int
	cpu     *CPU
	line    []bool
}

func (c *CRT) draw() {
	spriteMin, spriteMax := c.cpu.X-1, c.cpu.X+2
	if c.i >= spriteMin && c.i < spriteMax {
		c.line[c.i] = true
	} else {
		c.line[c.i] = false
	}
	c.i = (c.i + 1) % c.w
}

func (c *CRT) String() string {
	var line string
	for _, val := range c.line {
		if val == true {
			line = fmt.Sprintf("%s%s", line, "#")
		} else {
			line = fmt.Sprintf("%s%s", line, ".")
		}
	}
	return line
}

func getSignalStrength(cycle, signal int) int {
	return cycle * signal
}

func newCPU(program []string) *CPU {
	return &CPU{1, program}
}

func newCRT(w, h int, cpu *CPU) *CRT {
	return &CRT{w, h, 0, cpu, make([]bool, w)}
}

func GetResult1(input string) int {
	program := strings.Fields(input)
	cpu := newCPU(program)
	result := 0
	for i := range program {
		if (i+20+1)%40 == 0 {
			cycle := i + 1
			result += getSignalStrength(cycle, cpu.X)
		}
		cpu.tick()
	}
	return result
}

func GetResult2(input string) string {
	program := strings.Fields(input)
	cpu := newCPU(program)
	crt := newCRT(40, 6, cpu)
	result := ""
	for i := range program {
		crt.draw()
		if (i+1)%crt.w == 0 {
			result = fmt.Sprintf("%s%s\n", result, crt)
		}
		cpu.tick()
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-10/input.txt")
	fmt.Printf("Day 10 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 10 part 2 result is:\n%s\n", GetResult2(input))
}
