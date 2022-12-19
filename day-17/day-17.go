package day17

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
)

type Point struct { X, Y int}

func newPoint(x, y int) *Point {
	p := new(Point)
	p.X, p.Y = x, y
	return p
}

type Shape struct { 
	Pos *Point
	Points []*Point
}

func (s *Shape) TranslatePoint(p *Point) *Point {
	return newPoint(s.Pos.X + p.X, s.Pos.Y + p.Y)
}

func (s1 *Shape) Collides(s2 *Shape) bool {
	// compare every point in s1 to every point in s2
	// if any are the same, return true, else return false
	for _, p1 := range s1.Points {
		for _, p2 := range s2.Points {
			if *s1.TranslatePoint(p1) == *s2.TranslatePoint((p2)) {
				return true
			}
		}
	}
	return false
}

type Chamber struct {
	W int
	Active *Shape
	Fallen []*Shape
	Jets string
}

func newChamber(w int, jets string) *Chamber {
	c := new(Chamber)
	c.W = w
	c.Jets = jets
	return c
}

func getShapes() {
	rawShapes := utils.ReadFile("./shapes.txt")
	fmt.Println(rawShapes)
}

func GetResult1(input string) int {
	getShapes()
	return 0
}

func Run() {
	input := utils.ReadFile("./day-17/test_input.txt")
	fmt.Printf("Day 16 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 16 part 2 result is:\n%d\n", GetResult2(input))
}
