package main

import (
	"advent-of-code-2020/pkg/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	seatIDs := readSeatIDs()

	result1, duration1 := part1(seatIDs)
	fmt.Printf("Part 1: %10d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(seatIDs)
	fmt.Printf("Part 2: %10d (duration: %s)\n", result2, duration2)
}

func part1(seatIDs []int) (int, time.Duration) {
	start := time.Now()

	maxSeatID := seatIDs[len(seatIDs)-1]
	return maxSeatID, time.Since(start)
}

func part2(seatIDs []int) (int, time.Duration) {
	start := time.Now()

	offset := seatIDs[0]
	for i, seatID := range seatIDs {
		if offset+i != seatID {
			return offset + i, time.Since(start)
		}
	}
	return 0, time.Since(start)
}

func readSeatIDs() []int {
	lines := util.ReadFile("05")

	seatIDs := make([]int, len(lines))
	for i, line := range lines {
		row, _ := strconv.ParseInt(convertToBitField(line[:7]), 2, 8)
		seat, _ := strconv.ParseInt(convertToBitField(line[7:]), 2, 8)
		seatIDs[i] = int(row*8 + seat)
	}
	sort.Ints(seatIDs)
	return seatIDs
}

func convertToBitField(input string) string {
	builder := strings.Builder{}
	for i := 0; i < 7-len(input); i++ {
		builder.WriteString("0")
	}
	for _, r := range input {
		if r == 'F' || r == 'L' {
			builder.WriteString("0")
		} else if r == 'B' || r == 'R' {
			builder.WriteString("1")
		}
	}
	return builder.String()
}
