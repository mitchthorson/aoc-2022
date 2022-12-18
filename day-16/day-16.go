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
			fmt.Println("tunnel to", tId)
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

func GetResult1(input string) int {
	t := parseTunnels(strings.Split(input, "\n"))
	fmt.Println(t["AA"])
	// Instead of this queue, it would seem we need to find
	// the shortest paths to a set of weighted edges
	// A Floyd-Warshall algorithm can help with this:
	// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
	// queue of paths to check
	todo := []Step{{30, 0, make(map[string]struct{}), []string{"AA"}}}
	result := 0
	for len(todo) > 0 {
		step := todo[0]
		todo = todo[1:]
		if step.Minutes == 0 {
			if step.Points > result {
				result = step.Points
			}
			break
		}
		currentValve := step.Route[len(step.Route)-1]
		// next step will have one less minute
		nextMinutes := step.Minutes - 1
		// check if the current valve is already open
		_, valveOpen := step.OpenValves[currentValve]
		// if its not open, and its worth more than zero
		// add path where we turn it on
		if !valveOpen && t[currentValve].Rate > 0 {
			valves := step.OpenValves
			nextPoints := step.Points + t[currentValve].Rate*nextMinutes
			next := newStep(nextMinutes, nextPoints, valves, step.Route)
			valves[currentValve] = struct{}{}
			todo = append(todo, *next)
		}
		// now add all the possible move paths
		for _, nextValve := range t[currentValve].Tunnels {
			valves := step.OpenValves
			next := newStep(nextMinutes, step.Points, valves, append(step.Route, nextValve))
			todo = append(todo, *next)
		}

	}

	return result
}

func Run() {
	input := utils.ReadFile("./day-16/test_input.txt")
	fmt.Printf("Day 15 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 15 part 2 result is:\n%d\n", GetResult2(input, 4000000))
}
