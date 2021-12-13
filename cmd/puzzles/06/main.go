package main

import (
	"advent-of-code-2020/pkg/util"
	"bufio"
	"fmt"
	"time"
)

func main() {
	groups := readGroups()

	result1, duration1 := part1(groups)
	fmt.Printf("Part 1: %10d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(groups)
	fmt.Printf("Part 2: %10d (duration: %s)\n", result2, duration2)
}

func part1(groups [][]string) (int, time.Duration) {
	start := time.Now()

	sum := 0
	for _, group := range groups {
		counts := characterCounts(group)
		sum += len(counts)
	}
	return sum, time.Since(start)
}

func part2(groups [][]string) (int, time.Duration) {
	start := time.Now()

	sum := 0
	for _, group := range groups {
		size := len(group)
		counts := characterCounts(group)

		for _, count := range counts {
			if count == size {
				sum++
			}
		}
	}
	return sum, time.Since(start)
}

func characterCounts(group []string) map[rune]int {
	runes := make(map[rune]int, 0)
	for _, answers := range group {
		for _, r := range answers {
			runes[r]++
		}
	}
	return runes
}

func readGroups() [][]string {
	file := util.OpenFile("06")
	defer file.Close()

	groups := make([][]string, 0)
	answers := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			groups = append(groups, answers)
			answers = make([]string, 0)
		} else {
			answers = append(answers, text)
		}
	}
	groups = append(groups, answers)
	answers = make([]string, 0)
	return groups
}
