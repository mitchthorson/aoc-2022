package day16

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

type TunnelMap map[string]Valve

type Valve struct {
	Id      string
	Rate    int
	Tunnels []string
}

func (v Valve) getValue(minutes int) int {
	return v.Rate * minutes
}

func newValve(id string, rate int) *Valve {
	v := new(Valve)
	v.Id = id
	v.Rate = rate
	return v
}

func parseTunnels(lines []string) TunnelMap {
	tunnelMap := TunnelMap{}
	for _, line := range lines {
		splitLine := strings.Split(line, ";")
		valveInput := splitLine[0]
		var id string
		var rate int
		fmt.Sscanf(valveInput, "Valve %2s has flow rate=%d", &id, &rate)
		v := newValve(id, rate)
		tunnelInput := splitLine[1]
		tunnelInput = strings.TrimPrefix(tunnelInput, " tunnel leads to valve ")
		tunnelInput = strings.TrimPrefix(tunnelInput, " tunnels lead to valves ")
		for _, tId := range strings.Split(tunnelInput, ", ") {
			v.Tunnels = append(v.Tunnels, tId)
		}
		tunnelMap[v.Id] = *v
	}
	return tunnelMap
}

type Step struct {
	Minutes    int
	Points     int
	OpenValves map[string]struct{}
	Route      []string
}

func newStep(min, p int, open map[string]struct{}, route []string) *Step {
	s := new(Step)
	s.Minutes = min
	s.Points = p
	s.OpenValves = open
	s.Route = route
	return s
}

func contains(list []string, item string) bool {
	for _, l := range list {
		if l == item {
			return true
		}
	}
	return false
}

func calculateDistances(tunnelMap TunnelMap) map[string]map[string]int {
	distances := map[string]map[string]int{}
	INF := 1 << 30

	// build the matrix of all values set to INF
	for v1 := range tunnelMap {
		distances[v1] = map[string]int{}
		for v2 := range tunnelMap {
			distances[v1][v2] = INF
		}
	}
	// we know some of the distances are 1 from the input, set those here
	for v1, valve := range tunnelMap {
		for _, v2 := range valve.Tunnels {
			distances[v1][v2] = 1
		}
	}
	// now for the Floyd-Warshall algorithm
	// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
	for k := range tunnelMap {
		for i := range tunnelMap {
			for j := range tunnelMap {
				distances[i][j] = utils.Min(distances[i][j], distances[i][k]+distances[k][j])
			}
		}
	}
	return distances
}

func removeValve(valves []string, v string) []string {
	result := make([]string, 0, len(valves))
	for _, valve := range valves {
		if valve == v {
			continue
		}
		result = append(result, valve)
	}
	return result
}

// findPath uses depth-first search
func findPath(currentValve string, time, pressure int, toOpen []string, distances map[string]map[string]int, tunnelMap TunnelMap) int {
	max := pressure
	for _, destinationValve := range toOpen {
		// add one for opening valve
		distance := distances[currentValve][destinationValve] + 1
		newTime := time - distance
		if newTime > 0 {
			pathResult := findPath(destinationValve, newTime, pressure+newTime*tunnelMap[destinationValve].Rate, removeValve(toOpen, destinationValve), distances,
				tunnelMap)
			if pathResult > max {
				max = pathResult
			}
		}
	}
	return max
}

func divideList(list []string) [2][]string {
	halfLen := len(list) / 2
	partA := make([]string, 0, halfLen)
	partB := make([]string, 0, halfLen)
	for i := 0; i < halfLen; i++ {
		partA = append(partA, list[i])
		partB = append(partB, list[i+halfLen])
	}
	// if list has an odd number, add the extra item to list a
	if len(list)%2 > 0 {
		partA = append(partA, list[len(list)-1])
	}
	return [2][]string{partA, partB}
}

// this partitioning method doesn't work, it grossly underestimates the number of possible combinations.
// it also does not seem to always return the same partitions, which if i'm being honest confuses me a lot.
// it seems like it should return the same (incorrect) subset every time but it doesn't
// i see also that one of the problems is this assumes partitions of equal length
// that is probably an incorrect assumption. it also means this doesn't handle odd numbered input length
// but i still don't get the race condition
func partitionValves(valves []string) [][2][]string {
	halfSplit := divideList(valves)
	partitions := [][2][]string{halfSplit}
	for i := 0; i < len(valves)/2; i++ {
		for j := 0; j < len(valves)/2; j++ {
			splitList := divideList(valves)
			// swap item i from list a with item j from list b
			itemA := splitList[0][i]
			itemB := splitList[1][j]
			splitList[0][i] = itemB
			splitList[1][j] = itemA
			partitions = append(partitions, splitList)
		}
	}
	return partitions
}

// i have tried and failed to understand how to do this partitioning correctly
// this solution is borrowed from : https://github.com/lucianoq/adventofcode/blob/master/2022/16/main2.go
// i understand the basic gist of it, but its still beyond me.
func partition(list []string) [][2][]string {
	p := [][2][]string{}
	// i+=2 will generate half of the partitions.
	// We skip those because they'd be already
	// calculated by their symmetrical.
	for i := uint64(0); i < 1<<len(list); i += 2 {
		part := [2][]string{}
		for j := 0; j < len(list); j++ {
			if i&(1<<j) != 0 {
				part[0] = append(part[0], list[j])
			} else {
				part[1] = append(part[1], list[j])
			}
		}
		p = append(p, part)
	}
	return p
}

// found another solution for partinioning on stack overflow in python
// https://stackoverflow.com/questions/40709488/all-possibilities-to-split-a-list-into-two-lists
// here is my attempt to re-implement, which also seems to work
func partition2(list []string) [][2][]string {
	results := [][2][]string{}
	patterns := [][]int{}
	for i := 0; i < 1<<(len(list)-1); i++ {
		pattern := make([]int, 0, len(list))
		for j := 0; j < len(list)-1; j++ {
			pattern = append(pattern, i/(1<<j)%2)
		}
		patterns = append(patterns, pattern)
	}
	for _, p := range patterns {
		list1 := []string{list[len(list) - 1]}
		list2 := []string{}
		for i, val := range p {
			if val == 1 {
				list1 = append(list1, list[i])
			} else {
				list2 = append(list2, list[i])
			}
		}
		results = append(results, [2][]string{list1, list2})
	}
	return results
}

func GetResult1(input string) int {
	tunnels := parseTunnels(strings.Split(input, "\n"))
	valvesToOpen := []string{}
	for id, v := range tunnels {
		if v.Rate > 0 {
			valvesToOpen = append(valvesToOpen, id)
		}
	}
	distances := calculateDistances(tunnels)
	result := findPath("AA", 30, 0, valvesToOpen, distances, tunnels)
	return result
}
func GetResult2(input string) int {
	tunnels := parseTunnels(strings.Split(input, "\n"))
	valvesToOpen := []string{}
	for id, v := range tunnels {
		if v.Rate > 0 {
			valvesToOpen = append(valvesToOpen, id)
		}
	}
	distances := calculateDistances(tunnels)
	partitions := partition2(valvesToOpen)
	result := 0
	for _, p := range partitions {
		total := findPath("AA", 26, 0, p[0], distances, tunnels) + findPath("AA", 26, 0, p[1], distances, tunnels)
		if total > result {
			result = total
		}
	}
	return result
}

func Run() {
	input := utils.ReadFile("./day-16/input.txt")
	fmt.Printf("Day 16 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 16 part 2 result is:\n%d\n", GetResult2(input))
}
