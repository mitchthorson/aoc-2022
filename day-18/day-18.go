package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Lava struct {
	X, Y, Z [2]int
	Points [][3]int
}

func newLava() *Lava {
	INF := 1<<30
	l := new(Lava)
	defaultDimension := [2]int{INF, 0}
	l.X, l.Y, l.Z = defaultDimension, defaultDimension, defaultDimension
	return l
}

func (l *Lava) addPoint(p [3]int) {
	l.Points = append(l.Points, p)
	l.X[0], l.X[1] = utils.Min(l.X[0], p[0]), utils.Max(l.X[1], p[0])
	l.Y[0], l.Y[1] = utils.Min(l.Y[0], p[1]), utils.Max(l.Y[1], p[1])
	l.Z[0], l.Z[1] = utils.Min(l.Z[0], p[2]), utils.Max(l.Z[1], p[2])
}

func newPoint(input string) [3]int {
	p := [3]int{}
	for i, v := range strings.Split(input, ",") {
		val, err := strconv.Atoi(v)
		utils.Check(err)
		p[i] = val
	}
	return p
}

func scanSurface(l *Lava) int {
	total := 0
	// scan across x,y
	for a := l.X[0]
	return total
}

func GetResult1(input string) int {
	l := newLava()
	for _, p := range strings.Split(input, "\n") {
		point := newPoint(p)
		l.addPoint(point)
	}
	fmt.Println(l)
	return 0
}

func Run() {
	input := utils.ReadFile("./day-16/input.txt")
	fmt.Printf("Day 18 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 16 part 2 result is:\n%d\n", GetResult2(input))
}
