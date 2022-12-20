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
	jets     []string
	step     int
	shapeIdx int
	jetIdx   int
	history  map[[2]int]struct{}
	topo     []int
}

// TODO
// impement topo map that gets updated with each shape
func (c *Chamber) getTopo() []int {
	topo := make([]int, c.W)
	for i := 0; i < c.W; i++ {
		// each column
		baseHeight := c.HighestPoint()
		fmt.Println("baseheight", baseHeight)
		s := newShape("#", i, baseHeight)
		fmt.Println("scanning column", i)
		scan: for {
			for _, s2 := range c.Fallen {
				if s.Collides(s2) {
					fmt.Println("found surface")
					fmt.Println(s2)
					break scan
				}
			}
			if s.Pos.Y < 0 {
				break scan
			}
			fmt.Println("nothing found, scan continues down")
			// go down
			s.Pos.Y--
		}
		topo[i] = baseHeight - s.Pos.Y
	}
	return topo
}
// implement function to insert shapes based on a given 
// topo map at a given height

func (c *Chamber) String() string {
	return fmt.Sprintf("<Chamber %d, %s, fallen: %d, height: %d>", c.W, c.Active, len(c.Fallen), c.HighestPoint())
}

func (c *Chamber) JetPush() {
	dir := c.jets[c.jetIdx]
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

// returns a boolean that is true if a new shape is created
func (c *Chamber) ShapeDrop() bool {
	if c.Active.MinY() == 0 {
		c.Fallen = append(c.Fallen, c.Active)
		return true
	}
	c.Active.Pos.Y--
	for _, fs := range c.Fallen {
		if c.Active.Collides(fs) == true {
			// put it back to the previous position
			c.Active.Pos.Y++
			c.Fallen = append(c.Fallen, c.Active)
			return true
		}
	}
	return false
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
	s := newShape(rawShapes[c.shapeIdx], 2, currentHighest+4)
	s.Pos.Y += s.Height()
	c.shapeIdx = (c.shapeIdx + 1) % len(rawShapes)
	return s
}

// returns a boolean if the pattern repeats
func (c *Chamber) Next() bool {
	if c.step%2 == 0 {
		c.JetPush()
		c.jetIdx = (c.jetIdx + 1) % len(c.jets)
	} else {
		newShape := c.ShapeDrop()
		if newShape {
			// fmt.Println("new shape", c.shapeIdx, c.jetIdx)
			// update the history here, checking if the
			// current shape and wind index pairs have been seen before
			historyKey := [2]int{c.shapeIdx, c.jetIdx}
			_, seen := c.history[historyKey]
			c.history[historyKey] = struct{}{}
			c.Active = c.NextShape()
			if seen {
				return true
			}
		}
	}
	c.step++
	return false
}

func newChamber(w int, jets string) *Chamber {
	c := new(Chamber)
	c.W = w
	c.jets = strings.Split(jets, "")
	c.step = 0
	c.Active = c.NextShape()
	c.history = map[[2]int]struct{}{{0, 0}: {}}
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
	numShapes := 1
	c := newChamber(7, input)
	for {
		repeat := c.Next()
		if repeat {
			fmt.Println("we have a repeat, let's speed things up")
			shapesDropped, highest := len(c.Fallen), c.HighestPoint()
			multiples := numShapes / shapesDropped
			fmt.Println("dropped", shapesDropped, "high", highest, "can be repeated", multiples)
			fmt.Println(shapesDropped * multiples)
			// c.
			break
		}
		if len(c.Fallen) == numShapes {
			break
		}
	}
	fmt.Println(c.Fallen[0])
	fmt.Println(c.getTopo())
	return c.HighestPoint() + 1
}

func Run() {
	input := utils.ReadFile("./day-17/input.txt")
	// fmt.Printf("Day 17 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 17 part 2 result is:\n%d\n", GetResult2(input))
}
