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
	fmt.Printf("Part 1: %5d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2()
	fmt.Printf("Part 2: %5d (duration: %s)\n", result2, duration2)
}

func part1() (int, time.Duration) {
	start := time.Now()

	instructions := readInstructions("08/input")
	accumulator, _ := executeInstructions(instructions)
	return accumulator, time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()

	instructions := readInstructions("08/input")

	i := 0
	for true {
		if instructions[i].Operation == "acc" {
			i++
			continue
		}

		modified := fixInstructions(instructions, i)
		accumulator, success := executeInstructions(modified)
		if success {
			return accumulator, time.Since(start)
		}
		i++
	}

	return 0, time.Since(start)
}

func fixInstructions(instructions []Instruction, i int) []Instruction {
	current := make([]Instruction, 0, len(instructions))
	current = append(current, instructions...)

	instruction := &current[i]
	switch instruction.Operation {
	case "jmp":
		instruction.Operation = "nop"
	case "nop":
		instruction.Operation = "jmp"
	}
	return current
}

func executeInstructions(instructions []Instruction) (int, bool) {
	accumulator, i, length := 0, 0, len(instructions)

	for true {
		if i >= length {
			return accumulator, true
		}

		instruction := &instructions[i]

		if instruction.Executed {
			return accumulator, false
		}

		switch instruction.Operation {
		case "acc":
			accumulator += instruction.Value
			i++
		case "jmp":
			i += instruction.Value
		case "nop":
			i++
		}
		instruction.Executed = true
	}
	panic("ouch .. this should not happen")
}

type Instruction struct {
	Operation string
	Value     int
	Executed  bool
}

func readInstructions(day string) []Instruction {
	lines := util.ReadFile(day)
	instructionExpression, _ := regexp.Compile("([a-z]{3}) ([+\\-][0-9]+)")

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		match := instructionExpression.FindStringSubmatch(line)
		value, _ := strconv.Atoi(match[2])
		instructions[i] = Instruction{Operation: match[1], Value: value, Executed: false}
	}
	return instructions
}
