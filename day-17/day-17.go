package day17

import (
	"fmt"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Point struct{ X, Y int }

func (p *Point) String() string {
	return fmt.Sprintf("<Point %d %d>", p.X, p.Y)
}

func newPoint(x, y int) *Point {
	p := new(Point)
	p.X, p.Y = x, y
	return p
}

type Shape struct {
	Pos    *Point
	Points []*Point
}

func (s *Shape) TranslatePoint(p *Point) *Point {
	return newPoint(s.Pos.X+p.X, s.Pos.Y-p.Y)
}

func (s1 *Shape) Collides(s2 *Shape) bool {
	// compare every point in s1 to every point in s2
	// if any are the same, return true, else return false
	for _, p1 := range s1.Points {
		for _, p2 := range s2.Points {
			tp1, tp2 := s1.TranslatePoint(p1), s2.TranslatePoint(p2)
			if *tp1 == *tp2 {
				return true
			}
		}
	}
	return false
}

func (s *Shape) MaxY() int {
	max := 0
	for _, p := range s.Points {
		tp := s.TranslatePoint(p)
		if tp.Y > max {
			max = tp.Y
		}
	}
	return max
}

func (s *Shape) MinY() int {
	min := 1 << 30
	for _, p := range s.Points {
		tp := s.TranslatePoint(p)
		if tp.Y < min {
			min = tp.Y
		}
	}
	return min
}

func (s *Shape) MaxX() int {
	max := 0
	for _, p := range s.Points {
		tp := s.TranslatePoint(p)
		if tp.X > max {
			max = tp.X
		}
	}
	return max
}

func (s *Shape) MinX() int {
	min := 1 << 30
	for _, p := range s.Points {
		tp := s.TranslatePoint(p)
		if tp.X < min {
			min = tp.X
		}
	}
	return min
}

func (s *Shape) Height() int {
	return s.MaxY() - s.MinY()
}

func newShape(input string, x, y int) *Shape {
	s := new(Shape)
	s.Pos = newPoint(x, y)
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == '#' {
				p := newPoint(x, y)
				s.Points = append(s.Points, p)
			}
		}
	}
	return s
}

type Chamber struct {
	W        int
	Active   *Shape
	Fallen   []*Shape
	Jets     []string
	Step     int
	ShapeIdx int
	JetIdx   int
}

func (c *Chamber) String() string {
	return fmt.Sprintf("<Chamber %d, %s, fallen: %d, %v, %d>", c.W, c.Active, len(c.Fallen), c.Jets, c.Step)
}

func (c *Chamber) JetPush() {
	dir := c.Jets[c.JetIdx]
	c.JetIdx = (c.JetIdx + 1) % len(c.Jets)
	if dir == ">" {
		// if the max x of the shape is against the wall, return here
		if c.Active.MaxX() >= c.W-1 {
			return
		}
		c.Active.Pos.X++
		// check if it would collide with any shapes
		for _, s := range c.Fallen {
			if c.Active.Collides(s) == true {
				// if it does, put it back
				c.Active.Pos.X--
			}
		}
	}
	if dir == "<" {
		if c.Active.MinX() > 0 {
			c.Active.Pos.X--
			// check if it would collide with any shapes
			for _, s := range c.Fallen {
				if c.Active.Collides(s) == true {
					// if it does, put it back
					c.Active.Pos.X++
				}
			}
		}
	}
}

func (c *Chamber) ShapeDrop() {
	if c.Active.MinY() == 0 {
		c.Fallen = append(c.Fallen, c.Active)
		c.Active = c.NextShape()
		return
	}
	c.Active.Pos.Y--
	for _, fs := range c.Fallen {
		if c.Active.Collides(fs) == true {
			// put it back to the previous position
			c.Active.Pos.Y++
			c.Fallen = append(c.Fallen, c.Active)
			c.Active = c.NextShape()
			return
		}
	}
}

func (c *Chamber) HighestPoint() int {
	h := -1
	for _, s := range c.Fallen {
		if s.MaxY() > h {
			h = s.MaxY()
		}
	}
	return h
}

func (c *Chamber) NextShape() *Shape {
	currentHighest := c.HighestPoint()
	rawShapes := getShapes()
	s := newShape(rawShapes[c.ShapeIdx], 2, currentHighest+4)
	s.Pos.Y += s.Height()
	c.ShapeIdx = (c.ShapeIdx + 1) % len(rawShapes)
	return s
}

func (c *Chamber) Next() {
	if c.Step%2 == 0 {
		c.JetPush()
	} else {
		c.ShapeDrop()
	}
	c.Step++
}

func newChamber(w int, jets string) *Chamber {
	c := new(Chamber)
	c.W = w
	c.Jets = strings.Split(jets, "")
	c.Step = 0
	c.Active = c.NextShape()
	return c
}

func getShapes() []string {
	rawShapes := `####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`
	return strings.Split(rawShapes, "\n\n")
}

func GetResult1(input string) int {
	numShapes := 2022
	c := newChamber(7, input)
	for {
		c.Next()
		if len(c.Fallen) == numShapes {
			break
		}
	}
	return c.HighestPoint() + 1
}

func GetResult2(input string) int {
	// for part 2, i suspect that the pattern between the 
	// wind and the rocks repeats, giving us something we
	// can figure out based on that.
	// just need to determine the exact repitition point,
	// then use the shape count and the height to compute the final answer
	numShapes := 1000000000000
	c := newChamber(7, input)
	for {
		c.Next()
		if len(c.Fallen) == numShapes {
			break
		}
	}
	return c.HighestPoint() + 1
}

func Run() {
	input := utils.ReadFile("./day-17/input.txt")
	fmt.Printf("Day 16 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 16 part 2 result is:\n%d\n", GetResult2(input))
}
