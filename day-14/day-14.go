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

type Floor struct { Y int }

func (f Floor) Intersects(p *Point) bool {
	if p.Y == f.Y {
		fmt.Println("intersected floor", p)
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

func (s *Sand) Move(objects []Object) bool {
	newPos := new(Point)
	newPos.X = s.Pos.X
	newPos.Y = s.Pos.Y + 1
	for _, obj := range objects {
		// if there is antersection move left and check objects again
		if obj.Intersects(newPos) {
			fmt.Println("obj present at ", newPos, "moving left")
			newPos.X--
			for _, obj2 := range objects {
				if obj2.Intersects(newPos) {
					newPos.X += 2
					for _, obj3 := range objects {
						if obj3.Intersects(newPos) {
							// no where to go
							return false
						}
					}
				}
			}
		}
	}
	s.Pos = newPos
	return true
}

func parseRocks(lines []string) []Object {
	lineGroup := make([]Object, len(lines))
	for i, line := range lines {
		points := strings.Split(line, " -> ")
		linePoints := make([]*Point, len(points))
		for j, point := range points {
			p := newPoint(point)
			linePoints[j] = p
		}
		lineGroup[i] = *newLine(linePoints)
	}
	return lineGroup
}

func maxY(objects []Object) int {
	maxY := 0
	for _, o := range objects {
		objmaxY := o.MaxY()
		if objmaxY > maxY {
			maxY = objmaxY
		}
	}
	return maxY
}

func dropSand(objects []Object) ([]Object, bool) {
	origin := newSand("500,0")
	sand := newSand("500,0")
	for {
		moved := sand.Move(objects)
		if moved {
			if sand.Pos.Y > maxY(objects) {
				return objects, false
			}
		}  else if *sand.Pos == *origin.Pos {
			return objects, false
		} else {
			objects = append(objects, *sand)
			return objects, true
		}
	}

}

func GetResult1(input string) int {
	objects := parseRocks(strings.Split(input, "\n"))
	result := 0
	room := true
	for room == true {
		objects, room = dropSand(objects)
		if room == true {
			result += 1
		}
	}
	return result
}

func GetResult2(input string) int {
	objects := parseRocks(strings.Split(input, "\n"))
	floor := newFloor(maxY(objects) + 2)
	fmt.Println("floor", floor)
	objects = append(objects, floor)
	result := 0
	room := true
	for room == true {
		objects, room = dropSand(objects)
		if room == true {
			result += 1
		}
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-14/input.txt")
	fmt.Printf("Day 14 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 14 part 2 result is:\n%d\n", GetResult2(input))
}
