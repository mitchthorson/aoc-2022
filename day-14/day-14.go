package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Object interface {
	Intersects(p *Point) bool
	MaxY() int
}

type CaveMap struct {
	Floor int
	Map map[Point]struct{}
}

func newMap() *CaveMap {
	c := new(CaveMap)
	c.Map = map[Point]struct{}{}
	return c
}

func (c *CaveMap) fillPoint(p *Point) {
	c.Map[*p] = struct{}{}
}

func (c *CaveMap) fillLine(l *Line) {
	// c.Filled[*p] = struct{}{}
	for i, p := range l.Points[:len(l.Points) - 1] {
		p2 := l.Points[i + 1]
		// fill points along the x vector
		for j := utils.Min(p.X, p2.X); j <= utils.Max(p.X, p2.X); j++ {
			newP := new(Point)
			newP.X = j
			newP.Y = p.Y
			c.fillPoint(newP)
		}
		// fill points along the y vector
		for j := utils.Min(p.Y, p2.Y); j <= utils.Max(p.Y, p2.Y); j++ {
			newP := new(Point)
			newP.Y = j
			newP.X = p.X
			c.fillPoint(newP)
		}
	}
}

func (c *CaveMap) canMove(p *Point) bool {
	// check if we are hitting the floor
	if c.Floor > 0 && p.Y >= c.Floor {
		return false
	}
	_, exists := c.Map[*p]
	if exists {
		return false
	}
	return true
}

type Floor struct { Y int }

func (f Floor) Intersects(p *Point) bool {
	if p.Y == f.Y {
		return true
	}
	return false
}

func (f Floor) MaxY() int {
	return f.Y
}

func newFloor(level int) *Floor {
	return &Floor{level}
}

type Point struct{ X, Y int }

func (p *Point) String() string {
	return fmt.Sprintf("<Point %d %d>", p.X, p.Y)
}

func newPoint(input string) *Point {
	x_y_strs := strings.Split(input, ",")
	x, _ := strconv.Atoi(x_y_strs[0])
	y, _ := strconv.Atoi(x_y_strs[1])
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

type Line struct{ Points []*Point }

func newLine(points []*Point) *Line {
	l := new(Line)
	l.Points = points
	return l
}

func (l Line) Intersects(p *Point) bool {
	// x align
	for i, pointA := range l.Points {
		if i >= len(l.Points)-1 {
			break
		}
		pointB := l.Points[i+1]
		minX := utils.Min(pointA.X, pointB.X)
		maxX := utils.Max(pointA.X, pointB.X)
		minY := utils.Min(pointA.Y, pointB.Y)
		maxY := utils.Max(pointA.Y, pointB.Y)
		xIntersect := p.X >= minX && p.X <= maxX
		yIntersect := p.Y >= minY && p.Y <= maxY
		if xIntersect && yIntersect {
			return true
		}
	}
	return false
}

func (l Line) MaxY() int {
	maxY := 0
	for i, pointA := range l.Points {
		if i >= len(l.Points)-1 {
			break
		}
		pointB := l.Points[i+1]
		maxY = utils.Max(maxY, utils.Max(pointA.Y, pointB.Y))
	}
	return maxY
}

func (l *Line) String() string {
	output := "<Line "
	for _, p := range l.Points {
		output = fmt.Sprintf("%s %s", output, p)
	}
	return fmt.Sprintf("%s >", output)
}

type Sand struct {
	Pos *Point
}

func (s Sand) MaxY() int {
	return s.Pos.Y
}

func (s Sand) Intersects(p *Point) bool {
	if *s.Pos == *p {
		return true
	}
	return false

}

func newSand(pos string) *Sand {
	s := new(Sand)
	s.Pos = newPoint(pos)
	return s
}

func (s *Sand) Move(caveMap *CaveMap) bool {
	newPos := new(Point)
	newPos.X = s.Pos.X
	newPos.Y = s.Pos.Y + 1
	// if there is antersection move left and check objects again
	if caveMap.canMove(newPos) != true {
		newPos.X--
	}
	if caveMap.canMove(newPos) != true {
		newPos.X += 2
	}
	if caveMap.canMove(newPos) != true {
		// no  where to go
		return false
	}
	s.Pos = newPos
	return true
}

func parseRocks(lines []string) []*Line {
	lineGroup := make([]*Line, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " -> ")
		linePoints := make([]*Point, len(points))
		for j, point := range points {
			p := newPoint(point)
			linePoints[j] = p
		}
		lineGroup[i] = newLine(linePoints)
	}
	return lineGroup
}

func maxY(cave *CaveMap) int {
	maxY := 0
	if cave.Floor != 0 {
		return cave.Floor
	}
	for p := range cave.Map {
		maxY = utils.Max(maxY, p.Y)
	}
	return maxY
}

func pourSand(caveMap *CaveMap, sand *Sand, sandCount int) (*CaveMap, int) {
	// origin := newSand("500,0")
	moved := sand.Move(caveMap)
	if moved {
		// check for freefall
		if sand.Pos.Y > maxY(caveMap) {
			return caveMap, sandCount
		}
		return pourSand(caveMap, sand, sandCount)
	} else if sand.Pos.Y == 0 && sand.Pos.X == 500 {
		// if sand is at the origin and can't move
		sandCount = sandCount + 1
		return caveMap, sandCount
	} else {
		// if sand can't move, but its not at the entrance, fill its location in the map and do a new sand
		caveMap.fillPoint(sand.Pos)
		// recurse with next piece of sand
		sandCount = sandCount + 1
		return pourSand(caveMap, newSand("500,0"), sandCount)
	}
}

func GetResult1(input string) int {
	objects := parseRocks(strings.Split(input, "\n"))
	caveMap := newMap()
	for _, o := range objects {
		caveMap.fillLine(o)
	}
	_, result := pourSand(caveMap, newSand("500,0"), 0)
	return result
}

func GetResult2(input string) int {
	objects := parseRocks(strings.Split(input, "\n"))
	caveMap := newMap()
	for _, o := range objects {
		caveMap.fillLine(o)
	}
	caveMap.Floor = maxY(caveMap) + 2
	_, result := pourSand(caveMap, newSand("500,0"), 0)
	return result
}

func Run() {
	input := utils.ReadFile("./day-14/input.txt")
	fmt.Printf("Day 14 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 14 part 2 result is:\n%d\n", GetResult2(input))
}
