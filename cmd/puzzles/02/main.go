package main

import (
	"advent-of-code-2020/pkg/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type passwordPolicy struct {
	number1, number2    int
	character, password string
}

func (p passwordPolicy) validPart1() bool {
	count := strings.Count(p.password, p.character)
	return count >= p.number1 && count <= p.number2
}

func (p passwordPolicy) validPart2() bool {
	charPos1 := string(p.password[p.number1-1])
	charPos2 := string(p.password[p.number2-1])
	return (charPos1 == p.character) != (charPos2 == p.character)
}

func main() {
	policies := getPasswordPolicies()

	result1, duration1 := part1(policies)
	fmt.Printf("Part 1: %10d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(policies)
	fmt.Printf("Part 2: %10d (duration: %s)\n", result2, duration2)
}

func part1(policies []passwordPolicy) (int, time.Duration) {
	start := time.Now()

	valid := 0
	for _, policy := range policies {
		if policy.validPart1() {
			valid++
		}
	}
	return valid, time.Since(start)
}

func part2(policies []passwordPolicy) (int, time.Duration) {
	start := time.Now()

	valid := 0
	for _, policy := range policies {
		if policy.validPart2() {
			valid++
		}
	}
	return valid, time.Since(start)
}

func getPasswordPolicies() []passwordPolicy {
	list := util.ReadFile("02")
	pattern := regexp.MustCompile("([0-9]+)\\-([0-9]+) ([a-z]): ([a-z]+)")

	passwordPolicies := make([]passwordPolicy, 0, len(list))
	for _, line := range list {
		match := pattern.FindStringSubmatch(line)
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		policy := passwordPolicy{number1: n1, number2: n2, character: match[3], password: match[4]}
		passwordPolicies = append(passwordPolicies, policy)
	}

	return passwordPolicies
}
