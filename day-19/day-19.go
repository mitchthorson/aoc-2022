package day19

import (
	"fmt"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

// I think what I need to do here is similar to the other graph problem
// where I used the Floyd-Warshall algorithm to
// find the shortest paths in a tree
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
// but in this case there isn't a fixed distance between nodes,
// so it is going to take some creative thinking about how to prune branches

type Blueprint struct {
	Id                         int
	Ore, Clay, Obsidian, Geode [3]int
}

func parseBlueprint(input string) *Blueprint {
	bpSplit := strings.Split(input, ":")
	var bpId int
	fmt.Sscanf(bpSplit[0], "Blueprint %d", &bpId)
	bp := &Blueprint{bpId, [3]int{}, [3]int{}, [3]int{}, [3]int{}}
	for _, botCost := range strings.Split(bpSplit[1], ".") {
		var botType, resourceType string
		ore, extra := 0, 0
		fmt.Sscanf(botCost, " Each %s robot costs %d ore and %d %s.", &botType, &ore, &extra, &resourceType)
		cost := [3]int{ore, 0, 0}
		if resourceType == "clay" {
			cost[1] = extra
		}
		if resourceType == "obsidian" {
			cost[2] = extra
		}
		if botType == "ore" {
			bp.Ore = cost
		}
		if botType == "clay" {
			bp.Clay = cost
		}
		if botType == "obsidian" {
			bp.Obsidian = cost
		}
		if botType == "geode" {
			bp.Geode = cost
		}
	}
	return bp
}

type Factory struct {
	Bp        *Blueprint
	Time      int
	Rates     map[string]int
	Resources map[string]int
}

func newFactory(bp *Blueprint, time int) *Factory {
	defaultRates := map[string]int{
		"ore":      1,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	defaultResources := map[string]int{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
	return &Factory{bp, time, defaultRates, defaultResources}
}

func (current *Factory) addBot(botType string) *Factory {
	// generic settings for next Factory
	nextFactory := newFactory(current.Bp, current.Time)
	for k,v := range current.Rates {
		nextFactory.Rates[k] = v
	}
	for k,v := range current.Resources {
		nextFactory.Resources[k] = v
	}
	timeIncrement := 0
	// fmt.Println("adding bot of type", botType)
	// fmt.Println("current ore level", current.Resources["ore"])

	if botType == "ore" {
		timeIncrement = utils.Max((current.Bp.Ore[0] - current.Resources["ore"]) / current.Rates["ore"], 0) + 1
		nextFactory.Time = current.Time + timeIncrement
		nextFactory.Resources["ore"] = current.Resources["ore"] - current.Bp.Ore[0]
	} else if botType == "clay" {
		timeIncrement = utils.Max((current.Bp.Clay[0] - current.Resources["ore"]) / current.Rates["ore"], 0) + 1
		nextFactory.Time = current.Time + timeIncrement
		nextFactory.Resources["ore"] = current.Resources["ore"] - current.Bp.Clay[0]
	} else if botType == "obsidian" {
		timeIncrement = utils.Max(utils.Max((current.Bp.Obsidian[0]-current.Resources["ore"])/current.Rates["ore"], (current.Bp.Obsidian[1]-current.Resources["clay"])/current.Rates["clay"]), 0) + 1
		nextFactory.Time = current.Time + timeIncrement
		nextFactory.Resources["ore"] = current.Resources["ore"] - current.Bp.Obsidian[0]
		nextFactory.Resources["clay"] = current.Resources["clay"] - current.Bp.Obsidian[1]
	} else if botType == "geode" {
		timeIncrement = utils.Max(utils.Max((current.Bp.Geode[0]-current.Resources["ore"])/current.Rates["ore"], (current.Bp.Geode[2]-current.Resources["obsidian"])/current.Rates["obsidian"]), 0) + 1
		nextFactory.Time = current.Time + timeIncrement
		nextFactory.Resources["ore"] = current.Resources["ore"] - current.Bp.Geode[0]
		nextFactory.Resources["obsidian"] = current.Resources["obsidian"] - current.Bp.Geode[2]

	} else {
		panic(fmt.Sprintf("error adding bot of type %s\n", botType))
	}
	// and we increment each resource based on the current rate
	nextFactory.Resources["ore"] += current.Rates["ore"] * timeIncrement
	nextFactory.Resources["clay"] += current.Rates["clay"] * timeIncrement
	nextFactory.Resources["obsidian"] += current.Rates["clay"] * timeIncrement
	nextFactory.Resources["geode"] += current.Rates["geode"] * timeIncrement
	// the ore generation rate is increased by 1
	nextFactory.Rates[botType]++
	return nextFactory
}

func findMaxGeodes(factory *Factory) int {
	maxGeodes := factory.Resources["geode"]
	// if there is time remaining, create a queue of all the next
	// steps that can be taken
	// create a clone of the current factory to update
	checked := map[[8]int]int{}
	queue := []*Factory{factory}
	for len(queue) > 0 {
		// search for max
		current := queue[len(queue)-1]
		if len(queue) > 1 {
			queue = queue[:len(queue)-2]
		} else {
			queue = []*Factory{}
		}
		// @TODO need to do a better job at catching the end cases
		// need to prune branches that would jump past 24 minutes
		// also need to compute the results of harvesting for the remaining minutes
		// when a new bot can't be made
		if current.Time > 24 {
			continue
		}
		if current.Time == 24 {
			maxGeodes = utils.Max(current.Resources["geode"], maxGeodes)
			continue
		}
		checkKey := [8]int{current.Resources["ore"], current.Resources["clay"], current.Resources["obsidian"], current.Resources["geode"], current.Rates["ore"], current.Rates["clay"], current.Rates["obsidian"], current.Rates["geode"]}
		checkedMinute, alreadyChecked := checked[checkKey]
		if alreadyChecked {
			if checkedMinute < current.Time {
				continue
			}
		}
		checked[checkKey] = current.Time
		// check an ore bot
		queue = append(queue, current.addBot("ore"))
		queue = append(queue, current.addBot("clay"))
		if current.Rates["clay"] > 0 {
			queue = append(queue, current.addBot("obsidian"))
		}
		if current.Rates["obsidian"] > 0 {
			queue = append(queue, current.addBot("geode"))
		}
	}
	return maxGeodes
}

func GetResult1(input string) int {
	for _, bpInput := range strings.Split(input, "\n") {
		blueprint := parseBlueprint(bpInput)
		fact := newFactory(blueprint, 1)
		fmt.Println(findMaxGeodes(fact))
	}
	return 0
}

func Run() {
	input := utils.ReadFile("./day-19/test_input.txt")
	fmt.Printf("Day 19 part 1 result is:\n%d\n", GetResult1(input))
	// fmt.Printf("Day 19 part 2 result is:\n%d\n", GetResult2(input))
}
