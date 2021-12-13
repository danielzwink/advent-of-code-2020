package main

import (
	"advent-of-code-2020/pkg/util"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	numbers := getExpenseReport()

	result1, duration1 := part1(numbers)
	fmt.Printf("Part 1: %10d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(numbers)
	fmt.Printf("Part 2: %10d (duration: %s)\n", result2, duration2)
}

const NecessarySum = 2020

func part1(numbers []int) (int, time.Duration) {
	start := time.Now()
	sort.Ints(numbers)

	for indexA, summandA := range numbers {
		requiredSummandB := NecessarySum - summandA

		for indexB, summandB := range numbers {
			if indexA == indexB {
				continue
			}
			if summandB == requiredSummandB {
				return summandA * summandB, time.Since(start)
			}
			if summandB > requiredSummandB {
				break
			}
		}
	}
	return 0, time.Since(start)
}

func part2(numbers []int) (int, time.Duration) {
	start := time.Now()
	sort.Ints(numbers)

	for indexA, summandA := range numbers {
		for indexB, summandB := range numbers {
			if indexA == indexB {
				continue
			}
			requiredSummandC := NecessarySum - summandA - summandB
			if requiredSummandC < 0 {
				continue
			}

			for indexC, summandC := range numbers {
				if indexA == indexC || indexB == indexC {
					continue
				}
				if summandC == requiredSummandC {
					return summandA * summandB * summandC, time.Since(start)
				}
				if summandC > requiredSummandC {
					break
				}
			}
		}
	}
	return 0, time.Since(start)
}

func getExpenseReport() []int {
	list := util.ReadFile("01")

	expenseReport := make([]int, 0, len(list))
	for _, value := range list {
		number, _ := strconv.Atoi(value)
		expenseReport = append(expenseReport, number)
	}
	return expenseReport
}
