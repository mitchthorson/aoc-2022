package day08

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mitchthorson/aoc-2022/utils"
)

type Tree struct {
	x, y, height int
}

type TreeGrid struct {
	width, height int
	trees         []*Tree
}

func (tg *TreeGrid) checkVisible(treeIndex int) bool {
	tree := tg.trees[treeIndex]
	dirs := []string{"left", "up", "right", "down"}
	for _, dir := range dirs {
		switch dir {

		case "left":
			// left
			v := treeIndex - 1
			visible := true
			for (v % tg.width) < tg.width-1 {
				if tg.trees[v].height >= tree.height {
					visible = false
					break
				}
				v--
			}
			if visible == true {
				return true
			}
		case "up":
			v := treeIndex - tg.width
			visible := true
			for v > 0 {
				if tg.trees[v].height >= tree.height {
					visible = false
					break
				}
				v -= tg.width
			}
			if visible == true {
				return true
			}
		case "right":
			v := treeIndex + 1
			visible := true
			for v%tg.width > 0 {
				if tg.trees[v].height >= tree.height {
					visible = false
					break
				}
				v++
			}
			if visible == true {
				return true
			}
		case "down":
			v := treeIndex + tg.width
			visible := true
			for v < len(tg.trees) {
				if tg.trees[v].height >= tree.height {
					visible = false
					break
				}
				v += tg.width
			}
			if visible == true {
				return true
			}
		}
	}
	return false
}

func (tg *TreeGrid) getScenicScore(treeIndex int) int {
	tree := tg.trees[treeIndex]
	dirs := []string{"left", "up", "right", "down"}
	var scores [4]int
	for i, dir := range dirs {
		scores[i] = 0
		switch dir {

		case "left":
			// left
			if tree.x == 0 {
				continue
			}
			v := treeIndex - 1
			for (v % tg.width) < tg.width-1 && v >= 0 {
				scores[i]++
				if tg.trees[v].height >= tree.height {
					break
				}
				v--
			}
		case "up":
			if tree.y == 0 {
				continue
			}
			v := treeIndex - tg.width
			for v > 0 {
				scores[i]++
				if tg.trees[v].height >= tree.height {
					break
				}
				v -= tg.width
			}
		case "right":
			if tree.x == tg.width - 1 {
				continue
			}
			v := treeIndex + 1
			for v%tg.width > 0 {
				scores[i]++
				if tg.trees[v].height >= tree.height {
					break
				}
				v++
			}
		case "down":
			if tree.x == tg.height - 1 {
				continue
			}
			v := treeIndex + tg.width
			for v < len(tg.trees) {
				scores[i]++
				if tg.trees[v].height >= tree.height {
					break
				}
				v += tg.width
			}
		}
	}
	return scores[0] * scores[1] * scores[2] * scores[3]
}

func GetResult2(input string) int {
	lines := strings.Split(input, "\n")
	tg := new(TreeGrid)
	tg.height = len(lines)
	tg.width = len(lines[0])
	tg.trees = []*Tree{}
	for y, line := range lines {
		for x, char := range line {
			t := new(Tree)
			t.x = x
			t.y = y
			t.height = int(char - '0')
			tg.trees = append(tg.trees, t)
		}
	}
	treeScores := []int{}
	for i := range tg.trees {
		treeScores = append(treeScores, tg.getScenicScore(i))
	}
	sort.Ints(treeScores)
	return treeScores[len(treeScores) - 1]

}

func GetResult1(input string) int {
	lines := strings.Split(input, "\n")
	tg := new(TreeGrid)
	tg.height = len(lines)
	tg.width = len(lines[0])
	tg.trees = []*Tree{}
	for y, line := range lines {
		for x, char := range line {
			t := new(Tree)
			t.x = x
			t.y = y
			t.height = int(char - '0')
			tg.trees = append(tg.trees, t)
		}
	}
	visibleTrees := 0
	for i, tree := range tg.trees {
		// if tree is on the edge, it is visible
		if tree.x == 0 || tree.y == 0 || tree.x == tg.width-1 || tree.y == tg.height-1 {
			visibleTrees++
			continue
		}
		if tg.checkVisible(i) == true {
			visibleTrees++
		}
	}
	return visibleTrees
}

func Run() {
	input := utils.ReadFile("./day-08/input.txt")
	fmt.Printf("Day 08 part 1 result is:\n%d\n", GetResult1(input))
	fmt.Printf("Day 07 part 2 result is:\n%d\n", GetResult2(input))
}
