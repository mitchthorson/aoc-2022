package day03

import (
	"fmt"
	"github.com/mitchthorson/aoc-2022/utils"
	"strings"
)

func getDuplicate(rawInventory string) int {
	// create a map of all the items in the first compartment
	firstCompartmentItems := map[byte]struct{}{}
	for i := 0; i < len(rawInventory)/2; i++ {
		firstCompartmentItems[rawInventory[i]] = struct{}{}
	}
	// loop over items in the 2nd compartment
	// and check if they exist in the map of the first
	for i := len(rawInventory) / 2; i < len(rawInventory); i++ {
		_, ok := firstCompartmentItems[rawInventory[i]]
		if ok {
			return getPriority(rawInventory[i])
		}
	}
	return 0
}

func getBadge(group []string) int {
	groupItems := map[byte]int{}
	for _, inv := range group {
		items := map[byte]struct{}{}
		for i := 0; i < len(inv); i++ {
			items[inv[i]] = struct{}{}
		}
		for k := range items {
			_, ok := groupItems[k]
			if ok && groupItems[k] == 2 {
				return getPriority(k)
			} else if ok {
				groupItems[k] += 1
			} else {
				groupItems[k] = 1
			}
		}
	}
	return 0
}

func getPriority(item byte) int {
	if item < 'a' {
		return int(item-'A') + 27
	}
	return int(item-'a') + 1
}

func GetResult1(inventories []string) int {
	result := 0
	for _, inv := range inventories {
		result += getDuplicate(inv)
	}
	return result
}

func GetResult2(inventories []string) int {
	result := 0
	for i := 0; i < len(inventories); i += 3 {
		result += getBadge(inventories[i: i + 3])
	}
	return result
}

func Run() {
	input := utils.ReadInput(3)
	inventories := strings.Split(input, "\n")
	fmt.Printf("Day 03 part 1 result is:\n%d\n", GetResult1(inventories))
	fmt.Printf("Day 03 part 2 result is:\n%d\n", GetResult2(inventories))
}
