package day09

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

type Pos struct {
	x, y int
}

func (p *Pos) String() string {
	return fmt.Sprintf("Pos<%d, %d>", p.x, p.y)
}

type History map[Pos]interface{}

type Knot struct {
	id      int
	pos     Pos
	next    *Knot
	history *History
}

func newKnot(id int) *Knot {
	h := make(History)
	h[Pos{0, 0}] = 1
	return &Knot{id, Pos{0, 0}, nil, &h}
}

func (k *Knot) String() string {
	return fmt.Sprintf("Knot<%d: %s, %t>", k.id, &k.pos, (k.next == nil))
}

type Rope struct {
	head *Knot
	tail *Knot
}

func (r *Rope) getKnots() []Knot {
	k := []Knot{}
	c := r.head
	for c != nil {
		k = append(k, *c)
		c = c.next
	}
	return k
}

func (r *Rope) doMoves(moves []string) int {
	for _, move := range moves {
		var dir string
		var num int
		fmt.Sscanf(move, "%s %d", &dir, &num)
		r.move(dir, num)
	}
	return len(*r.tail.history)
}

func newRope(num int) *Rope {
	r := new(Rope)
	r.head = newKnot(0)
	c := r.head
	for i := 0; i < num-1; i++ {
		c.next = newKnot(i + 1)
		c = c.next
		r.tail = c
	}
	return r
}

func (r *Rope) move(dir string, num int) *Rope {
	for i := 0; i < num; i++ {
		switch dir {
		case "R":
			r.head.pos.x++
		case "U":
			r.head.pos.y++
		case "D":
			r.head.pos.y--
		case "L":
			r.head.pos.x--
		}
		currentKnot := r.head
		for {
			_, ok := (*currentKnot.history)[currentKnot.pos]
			if !ok {
				(*currentKnot.history)[currentKnot.pos] = 1
			}
			if currentKnot.next == nil {
				break
			}
			// check x
			if utils.Abs(currentKnot.pos.x-currentKnot.next.pos.x) > 1 {
				currentKnot.next.pos.x += normalizeDistance(currentKnot.pos.x - currentKnot.next.pos.x)
				currentKnot.next.pos.y += normalizeDistance(currentKnot.pos.y - currentKnot.next.pos.y)
			}
			// check y
			if utils.Abs(currentKnot.pos.y-currentKnot.next.pos.y) > 1 {
				currentKnot.next.pos.y += normalizeDistance(currentKnot.pos.y - currentKnot.next.pos.y)
				currentKnot.next.pos.x += normalizeDistance(currentKnot.pos.x - currentKnot.next.pos.x)
			}
			if currentKnot.next == nil {
			}
			currentKnot = currentKnot.next
		}
	}

	return r
}

func normalizeDistance(d int) int {
	if d == 0 {
		return 0
	}
	if d > 0 {
		return 1
	}
	return -1
}

// func moveAmount()

// func (r *Rope) String() string {
// 	return fmt.Sprintf()
//
// }

func getBoundingPositions(moves []string) []Pos {
	r := newRope(1)
	for _, move := range moves {
		var dir string
		var num int
		fmt.Sscanf(move, "%s %d", &dir, &num)
		r.move(dir, num)
	}
	return []Pos{}
	// return len(*r.tail.history)
}

func GetResult1(input string) int {
	moves := strings.Split(input, "\n")
	rope := newRope(2)
	result := rope.doMoves(moves)
	return result
}

func GetResult2(input string) int {
	moves := strings.Split(input, "\n")
	rope := newRope(10)
	result := rope.doMoves(moves)
	return result
}

func Run() {
	input := utils.ReadFile("./day-09/input.txt")
	fmt.Printf("Day 09 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 09 part 2 result is:\n%d\n", GetResult2(input))
}
