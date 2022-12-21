package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Point [3]int

type Lava struct {
	W, H, D int
	Points  []Point
	Matrix   []bool
}

func (l *Lava) FillPoint(x, y, z int) {
	pointIdx := l.W*l.H*z + l.W*y + x
	l.Matrix[pointIdx] = true
}

func (l *Lava) HasPoint(p Point) bool {
	x, y, z := p[0], p[1], p[2]
	if x < 0 || y < 0 || z < 0 || x > l.W-1 || y > l.H-1 || z > l.D-1 {
		return false
	}
	pointIdx := l.W*l.H*z + l.W*y + x
	return l.Matrix[pointIdx]
}

func (l *Lava) GetPoint(pointIdx int) Point {
	// reverse of this: pointIdx := l.W*l.H*z + l.W*y + x
	x := pointIdx % l.W * l.H
	y := (pointIdx - x) % l.W
	z := (pointIdx - x - y*l.W) % l.H
	return Point{x, y, z}
}

func max(points []Point, v int) int {
	max := 0
	for _, p := range points {
		max = utils.Max(max, p[v])
	}
	return max
}

func min(points []Point, v int) int {
	min := 1 << 30
	for _, p := range points {
		min = utils.Min(min, p[v])
	}
	return min
}

func newLava(points []Point) *Lava {
	l := new(Lava)
	l.W = max(points, 0) + 1
	l.H = max(points, 1) + 1
	l.D = max(points, 2) + 1
	l.Points = points
	l.Matrix = make([]bool, l.W*l.H*l.D)
	for _, p := range points {
		l.FillPoint(p[0], p[1], p[2])
	}
	return l
}

func newPoint(input string) Point {
	p := Point{}
	for i, v := range strings.Split(input, ",") {
		val, err := strconv.Atoi(v)
		utils.Check(err)
		p[i] = val
	}
	return p
}

func (l *Lava) GetEmptyNeighbors(p Point) []Point {
	n := make([]Point, 0, 6)
	for i := 0; i < 3; i++ {
		a := Point{p[0], p[1], p[2]}
		b := Point{p[0], p[1], p[2]}
		a[i]++
		b[i]--
		if !(l.HasPoint(a)) {
			n = append(n, a)
		}
		if !(l.HasPoint(b)) {
			n = append(n, b)
		}
	}
	return n
}

func (l *Lava) GetSurfaceArea(checkBubbles bool) int {
	area := 0
	for _, p := range l.Points {
		// loop through each dimension
		for _, n := range l.GetEmptyNeighbors(p) {
			if checkBubbles {
				// fmt.Println("Checking if empty space is reachable", n)
				if l.IsReachable(n) {
					area++
				}
			} else {
				area++
			}
		}
	}
	return area
}

func (l *Lava) IsEdge(p Point) bool {
	if p[0] <= 0 || p[0] >= l.W || p[1] <= 0 || p[1] >= l.H || p[2] <= 0 || p[2] >= l.D {
		return true
	}
	return false
}

func (l *Lava) IsReachable(p Point) bool {
	// if point is on the outside edge of the lava droplet, return true
	if l.IsEdge(p) {
		return true
	}
	// otherwise, create a queue of all the neighboring points
	queue := l.GetEmptyNeighbors(p)
	checked := map[Point]struct{}{}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if l.IsEdge(item) {
			return true
		}
		for _, n := range l.GetEmptyNeighbors(item) {
			_, hasBeenChecked := checked[item]
			if !hasBeenChecked {
				queue = append(queue, n)
			}
		}
		checked[item] = struct{}{}
	}
	return false
}

func parsePoints(input string) []Point {
	lines := strings.Split(input, "\n")
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		p := newPoint(line)
		points = append(points, p)
	}
	return points
}

func GetResult1(input string) int {
	points := parsePoints(input)
	l := newLava(points)
	return l.GetSurfaceArea(false)
}

func GetResult2(input string) int {
	points := parsePoints(input)
	l := newLava(points)
	return l.GetSurfaceArea(true)
}

func Run() {
	input := utils.ReadFile("./day-18/input.txt")
	fmt.Printf("Day 18 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 18 part 2 result is:\n%d\n", GetResult2(input))
}
