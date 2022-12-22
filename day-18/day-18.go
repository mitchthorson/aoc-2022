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
	Cubes []bool
}

func (l *Lava) FillPoint(x, y, z int) {
	pointIdx := l.W*l.H*z + l.W*y + x
	l.Cubes[pointIdx] = true
}
func (l *Lava) GetPoint(x, y, z int) bool {
	if x < 0 || y < 0 || z < 0 || x > l.W - 1 || y > l.H - 1 || z > l.D - 1 {
		return false
	}
	pointIdx := l.W*l.H*z + l.W*y + x
	return l.Cubes[pointIdx]
}

func max(points []Point, v int) int {
	max := 0
	for _, p := range points {
		max = utils.Max(max, p[v])
	}
	return max
}

func min(points []Point, v int) int {
	min := 1<<30
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
	l.Cubes = make([]bool, l.W*l.H*l.D)
	for _, p := range points {
		l.FillPoint(p[0], p[1], p[2])
	}
	return l
}

// func (l *Lava) addPoint(p [3]int) {
// 	l.Points = append(l.Points, p)
// 	l.X[0], l.X[1] = utils.Min(l.X[0], p[0]), utils.Max(l.X[1], p[0])
// 	l.Y[0], l.Y[1] = utils.Min(l.Y[0], p[1]), utils.Max(l.Y[1], p[1])
// 	l.Z[0], l.Z[1] = utils.Min(l.Z[0], p[2]), utils.Max(l.Z[1], p[2])
// }

func newPoint(input string) Point {
	p := Point{}
	for i, v := range strings.Split(input, ",") {
		val, err := strconv.Atoi(v)
		utils.Check(err)
		p[i] = val
	}
	return p
}

// func scanSurface(l *Lava) int {
// 	total := 0
// 	// scan across x,y
// 	for a := l.X[0]
// 	return total
// }

func (l *Lava) GetEmptyNeighbors(p Point) []Point {
	n := make([]Point, 0, 6)
	for i := 0; i < 3; i++ {
		a := Point{p[0], p[1], p[2]}
		b := Point{p[0], p[1], p[2]}
		a[i]++
		b[i]--
		if !(l.GetPoint(a[0], a[1], a[2])) {
			n = append(n, a)
		}
		if !(l.GetPoint(b[0], b[1], b[2])) {
			n = append(n, b)
		}
	}
	return n
}

func (l *Lava) GetSurfaceArea(points []Point, checkBubbles bool) int {
	area := 0
	for _, p := range points {
		// loop through each dimension
		for _, n := range l.GetEmptyNeighbors(p) {
			if !l.GetPoint(n[0], n[1], n[2]) {
				if checkBubbles {
					if l.IsReachable(n) {
						area++
					}
				} else {
					area++ 
				}
			}
		}
	}
	return area
}

func (l *Lava) IsEdge(p Point) bool {
	if p[0]<=0 || p[0] >= l.W || p[1] <= 0 || p[1] >= l.H || p[2] <= 0 || p[2] >= l.D {
		return true
	}
	return false
}

func (l *Lava) IsReachable(p Point) bool {
	// @TODO implement IsReachable
	// should return true only if it can walk 
	// to the edge of the lava droplet
	if l.IsEdge(p) {
		return true
	}
	checked := map[Point]struct{}{}
	// if point is on the outside edge of the lava droplet, return true
	// create a queue of all the neighboring points
	queue := l.GetEmptyNeighbors(p)
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		fmt.Println("checking", item)
		if l.IsEdge(item) {
			fmt.Println("reached edge")
			return true
		}
		checked[item] = struct{}{}
		for _, n := range l.GetEmptyNeighbors(item) {
			_, hasBeenChecked := checked[item]
			if !hasBeenChecked {
				queue = append(queue[1:], n)
			}
		}
	}
	fmt.Println("ran out of places to look")
	return false
}

func parsePoints(input string) []Point {
	lines := strings.Split(input, "\n")
	points := make([]Point,0,len(lines))
	for _, line := range lines {
		p := newPoint(line)
		points = append(points, p)
	}
	return points
}

func GetResult1(input string) int {
	points := parsePoints(input)
	l := newLava(points)
	return l.GetSurfaceArea(points, false)
}
func GetResult2(input string) int {
	points := parsePoints(input)
	l := newLava(points)
	return l.GetSurfaceArea(points, true)
}

func Run() {
	input := utils.ReadFile("./day-18/input.txt")
	fmt.Printf("Day 18 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 18 part 2 result is:\n%d\n", GetResult2(input))
}
