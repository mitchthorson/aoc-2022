package day12

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

type PointMap struct {
	W,H int
	Points []*Point
}

func (m *PointMap) Scan(p *Point) []*Point {
	paths := []*Point{}
	// can i calculate the point index in the slice?
	pointIndex := p.X + (p.Y * m.W)
	fmt.Println("Point index calculated: ", pointIndex)
	// up
	if p.Y > 1 {
		paths = append(paths, m.Points[pointIndex - m.W])
	}
	// right
	if p.X < m.W - 1 {
		paths = append(paths, m.Points[pointIndex + 1])
	}
	// down
	if p.Y < m.H - 1 {
		paths = append(paths, m.Points[pointIndex + m.W])
	}
	// left
	if p.X > 0 {
		paths = append(paths, m.Points[pointIndex - 1])
	}
	return paths

}

type Point struct {
	X, Y, H int 
	Start, End bool
}

func (p *Point) String() string {
	return fmt.Sprintf("<Point x=%d y=%d h=%d start=%t end=%t>", p.X, p.Y, p.H, p.Start, p.End)
}

// func (p *Point) IsPath(p2 *Point) bool {
// 	if(p2.H < p.H+2) {
//
// 	}
// }

func newPoint(x, y int, c rune) *Point {
	p := new(Point)
	h := parsePointHeight(c)
	s := c == 'S'
	e := c == 'E'
	p.X = x
	p.Y = y
	p.H = h
	p.Start = s
	p.End = e
	return p
}

type PathTree struct {
	Point Point
	Paths[]*PathTree
}

func (p *PathTree) append(newPath *PathTree) *PathTree {
	p.Paths = append(p.Paths, newPath)
	return p
}

func parsePointHeight(input rune) int {
	if input == 'S' {
		return 0
	}
	if input == 'E' {
		return 25
	}

	return int(input - 'a')
}

func parseMap(input string) *PointMap {
	m := new(PointMap)
	points := []*Point{}
	lines := strings.Split(input, "\n")
	fmt.Println(lines)
	m.H = len(lines)
	m.W = len(lines[0])
	fmt.Println(m.W, m.H)
	for y, line := range lines {
		for x, char := range line {
			p := newPoint(x, y, char)
			points = append(points, p)
		}
	}
	m.Points = points
	return m
}

func GetResult2(input string) int {
	return 0
}

func GetResult1(input string) int {
	pointMap := parseMap(input)
	fmt.Println(pointMap)
	return 0
}

func Run() {
	input := utils.ReadFile("./day-12/test_input.txt")
	fmt.Printf("Day 11 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 11 part 2 result is:\n%d\n", GetResult2(input))
}
