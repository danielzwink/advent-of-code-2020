package main

import (
	"advent-of-code-2020/pkg/util"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	result1, duration1 := part1()
	fmt.Printf("Part 1: %6d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2()
	fmt.Printf("Part 2: %6d (duration: %s)\n", result2, duration2)
}

func part1() (int, time.Duration) {
	start := time.Now()
	lookup := readBagLookup("07/input")

	colors := make(map[string]bool, 10)
	colors["shiny gold"] = true

	for true {
		additionalColors := make(map[string]bool, 10)

		for color := range colors {
			for bagColor, content := range lookup {
				for contentColor := range content {
					if contentColor == color {
						additionalColors[bagColor] = true
						break
					}
				}
			}
		}

		lenBefore := len(colors)
		for addColor := range additionalColors {
			colors[addColor] = true
		}
		lenAfter := len(colors)

		if lenAfter == lenBefore {
			break
		}
	}

	return len(colors) - 1, time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()
	lookup := readBagLookup("07/input")

	sum := sumBags("shiny gold", lookup)
	return sum, time.Since(start)
}

func sumBags(color string, lookup map[string]map[string]int) int {
	content, found := lookup[color]

	if !found || len(content) == 0 {
		return 0
	}

	sum := 0
	for color, size := range content {
		sum += size
		sum += size * sumBags(color, lookup)
	}
	return sum
}

func readBagLookup(day string) map[string]map[string]int {
	lines := util.ReadFile(day)

	bagRegexp, _ := regexp.Compile("([a-z]+ [a-z]+) bag[s]? contain")
	contentRegexp, _ := regexp.Compile("([0-9]+) ([a-z]+ [a-z]+) bag[s]?")

	lookup := make(map[string]map[string]int, len(lines))
	for _, line := range lines {
		bagSpec := bagRegexp.FindStringSubmatch(line)
		contentSpecs := contentRegexp.FindAllStringSubmatch(line, 10)

		content := make(map[string]int, len(contentSpecs))
		for _, contentSpec := range contentSpecs {
			count, _ := strconv.Atoi(contentSpec[1])
			content[contentSpec[2]] = count
		}
		lookup[bagSpec[1]] = content
	}
	return lookup
}
