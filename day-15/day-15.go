package day15

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

type Point struct{ X, Y int }

func newPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}

type BeaconMap struct {
	zones map[Point]struct{}
}

func (b *BeaconMap) maskSensorRow(s *Sensor, row int, skipBeacon bool, minMax []int) {
	cx, cy, d := s.Pos.X, s.Pos.Y, s.Dist
	xSize := d - utils.Abs(row-cy)
	xStart := cx - xSize
	xEnd := cx + xSize
	if len(minMax) == 2 {
		xStart = utils.Min(utils.Max(xStart, minMax[0]), minMax[1])
		xEnd = utils.Max(utils.Min(xEnd, minMax[1]), minMax[0])
	}

	for x := xStart; x <= xEnd; x++ {
		p := newPoint(x, row)
		// skip actual beacon
		if skipBeacon && p.X == s.Beacon.X && p.Y == s.Beacon.Y {
			continue
		}
		b.zones[*p] = struct{}{}
	}
}

func newBeaconMap() *BeaconMap {
	bm := new(BeaconMap)
	bm.zones = make(map[Point]struct{})
	return bm
}

type Sensor struct {
	Pos    *Point
	Beacon *Point
	Dist   int
}

func (s *Sensor) String() string {
	return fmt.Sprintf("<Sensor x=%d y=%d dist=%d>", s.Pos.X, s.Pos.Y, s.Dist)
}

func newSensor(sx, sy, bx, by int) *Sensor {
	s := new(Sensor)
	s.Pos = newPoint(sx, sy)
	s.Dist = distance(sx, sy, bx, by)
	s.Beacon = newPoint(bx, by)
	return s
}

func (s *Sensor) getRowMax(row int) int {
	cx, cy, d := s.Pos.X, s.Pos.Y, s.Dist
	xSize := d - utils.Abs(row-cy)
	return cx + xSize
}

func distance(x1, y1, x2, y2 int) int {
	return utils.Abs(x2-x1) + utils.Abs(y2-y1)
}

func parseSensors(lines []string) []*Sensor {
	sensors := make([]*Sensor, 0, len(lines))
	for _, line := range lines {
		var sx, sy, bx, by int
		// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		s := newSensor(sx, sy, bx, by)
		sensors = append(sensors, s)
	}
	return sensors
}

func GetResult1(input string, row int) int {
	s := strings.Split(input, "\n")
	sensors := parseSensors(s)
	beaconMap := newBeaconMap()
	for _, s := range sensors {
		beaconMap.maskSensorRow(s, row, true, []int{})
	}
	return len(beaconMap.zones)
}
func GetResult2(input string, max int) int {
	s := strings.Split(input, "\n")
	sensors := parseSensors(s)
	// loop through rows
	for y := 0; y <= max; y++ {
		// loop through points in row
	pointLoop:
		for x := 0; x <= max; x++ {
			p := Point{x, y}
			for _, s := range sensors {
				if distance(s.Pos.X, s.Pos.Y, x, y) <= s.Dist {
					// here we find ourselves inside the zone for s
					// what we need to do is jump to the end of the
					// zone for this row, which we can calculate
					// and then skip to the next point loop and keep going
					x = s.getRowMax(y)
					continue pointLoop
				}
			}
			// if we end up here, we searched through every
			//beacon and found an undetexted point
			return p.X*4000000 + y
		}
	}
	return 0
}

func Run() {
	input := utils.ReadFile("./day-15/input.txt")
	fmt.Printf("Day 15 part 1 result is:\n%d\n", GetResult1(input, 2000000))
	fmt.Printf("Day 15 part 2 result is:\n%d\n", GetResult2(input, 4000000))
}
