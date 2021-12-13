package main

import (
	"advent-of-code-2020/pkg/util"
	"fmt"
	"time"
)

func main() {
	treeMap := getTreeMap()

	result1, duration1 := part1(treeMap)
	fmt.Printf("Part 1: %10d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(treeMap)
	fmt.Printf("Part 2: %10d (duration: %s)\n", result2, duration2)
}

func part1(treeMap [][]int) (int, time.Duration) {
	start := time.Now()

	trees := countTrees(treeMap, 3, 1)
	return trees, time.Since(start)
}

func part2(treeMap [][]int) (int, time.Duration) {
	start := time.Now()

	trees1 := countTrees(treeMap, 1, 1)
	trees2 := countTrees(treeMap, 3, 1)
	trees3 := countTrees(treeMap, 5, 1)
	trees4 := countTrees(treeMap, 7, 1)
	trees5 := countTrees(treeMap, 1, 2)
	return trees1 * trees2 * trees3 * trees4 * trees5, time.Since(start)
}

func countTrees(treeMap [][]int, right, down int) (trees int) {
	mapHeight := len(treeMap)
	mapWidth := len(treeMap[0])

	col := 0
	row := 0
	for row < mapHeight {
		if col >= mapWidth {
			col -= mapWidth
		}

		trees += treeMap[row][col]
		row += down
		col += right
	}
	return trees
}

func getTreeMap() [][]int {
	list := util.ReadFile("03")

	treeMap := make([][]int, 0, 10)
	for _, line := range list {
		trees := make([]int, 0, len(line))
		for _, v := range line {
			if v == '#' {
				trees = append(trees, 1)
			} else {
				trees = append(trees, 0)
			}
		}
		treeMap = append(treeMap, trees)
	}
	return treeMap
}
