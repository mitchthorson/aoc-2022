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
		newTime := time-distance
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

func partitionValves(valves []string) [][2][]string {
	partitions := [][2][]string{}
	for i := 0; i < 1<<len(valves); i+=2 {
		partition := [2][]string{}
		// i might need help with this. 

	}

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

func Run() {
	input := utils.ReadFile("./day-16/input.txt")
	fmt.Printf("Day 15 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 15 part 2 result is:\n%d\n", GetResult2(input, 4000000))
}
