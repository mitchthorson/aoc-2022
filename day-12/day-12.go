package day12

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	// "sort"
	"strings"
)

type PointMap struct {
	W, H       int
	Start, End *Point
	Points     []*Point
}

type Point struct {
	X, Y, H int
}

func (p *Point) Scan(m *PointMap) []*Point {
	paths := []*Point{}
	// can i calculate the point index in the slice?
	pointIndex := p.X + (p.Y * m.W)
	// up
	if p.Y > 0 {
		paths = append(paths, m.Points[pointIndex-m.W])
	}
	// right
	if p.X < m.W-1 {
		paths = append(paths, m.Points[pointIndex+1])
	}
	// down
	if p.Y < m.H-1 {
		paths = append(paths, m.Points[pointIndex+m.W])
	}
	// left
	if p.X > 0 {
		paths = append(paths, m.Points[pointIndex-1])
	}
	return paths

}

func (p *Point) String() string {
	return fmt.Sprintf("<Point x=%d y=%d h=%d>", p.X, p.Y, p.H)
}

// func (p *Point) IsPath(p2 *Point) bool {
// 	if(p2.H < p.H+2) {
//
// 	}
// }

func newPoint(x, y int, c rune) *Point {
	p := new(Point)
	h := parsePointHeight(c)
	p.X = x
	p.Y = y
	p.H = h
	return p
}

type PathPoint struct {
	Point *Point
	// Parent *PathTree
	// Children []*PathTree
	Depth int
}

// func (p *PathTree) appendTo(newPath *PathTree) *PathTree {
// 	// p.Parent = newPath
// 	p.Depth = newPath.Depth + 1
// 	// newPath.Children = append(newPath.Children, p)
// 	return p
// }

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
	m.H = len(lines)
	m.W = len(lines[0])
	for y, line := range lines {
		for x, char := range line {
			p := newPoint(x, y, char)
			points = append(points, p)
			if char == 'S' {
				m.Start = p
			}
			if char == 'E' {
				m.End = p
			}
		}
	}
	m.Points = points
	return m
}

func (p *Point) canStep(p2 *Point) bool {
	if p2.H-p.H < 2 {
		return true
	}
	return false
}
func (p *Point) canStepReverse(p2 *Point) bool {
	if p.H-p2.H < 2 {
		return true
	}
	return false
}

// func newTree(p *Point) *PathTree {
// 	tree := new(PathTree)
// 	tree.Point = p
// 	return tree
// }

// func (tree *PathTree) hasParentPoint(p *Point) bool {
// 	if tree == nil {
// 		return false
// 	}
// 	return tree.Point == p || tree.Parent.hasParentPoint(p)
// }

func (start *PathPoint) walkToEnd(m *PointMap, isEnd func(*Point, *PointMap) bool, reverse bool) int {
	pathHistory := map[Point]struct{}{*start.Point: {}}
	pathQueue := []*PathPoint{start}
	var curr PathPoint
	for len(pathQueue) > 0 {
		curr, pathQueue = *pathQueue[0], pathQueue[1:]
		if isEnd(curr.Point, m) {
			return curr.Depth
		}
		neighbors := curr.Point.Scan(m)
		for _, n := range neighbors {
			// if the neighbor cant be stepped to, skip
			if !reverse && !curr.Point.canStep(n) {
				continue
			}
			if reverse && !curr.Point.canStepReverse(n) {
				continue
			}
			// if the neighbor has been visited, skip
			if _, ok := pathHistory[*n]; ok {
				continue
			}
			pathHistory[*n] = struct{}{}
			pathQueue = append(pathQueue, &PathPoint{Point: n, Depth: curr.Depth + 1})
		}

	}
	return 0
}

func GetResult2(input string) int {
	pointMap := parseMap(input)
	pathPoint := new(PathPoint)
	pathPoint.Point = pointMap.End
	pathPoint.Depth = 0
	isEnd := func(p *Point, m *PointMap) bool {
		return p.H == 0
	}
	result := pathPoint.walkToEnd(pointMap, isEnd, true)
	return result
}

func GetResult1(input string) int {
	pointMap := parseMap(input)
	pathPoint := new(PathPoint)
	pathPoint.Point = pointMap.Start
	pathPoint.Depth = 0
	isEnd := func(p *Point, m *PointMap) bool {
		return p == m.End
	}
	result := pathPoint.walkToEnd(pointMap, isEnd, false)
	return result
}

func Run() {
	input := utils.ReadFile("./day-12/input.txt")
	fmt.Printf("Day 11 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 11 part 2 result is:\n%d\n", GetResult2(input))
}
