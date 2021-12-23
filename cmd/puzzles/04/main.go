package main

import (
	"advent-of-code-2020/pkg/util"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	result1, duration1 := part1(getPassports("04/input", NewPassport1))
	fmt.Printf("Part 1: %4d (duration: %s)\n", result1, duration1)

	result2, duration2 := part2(getPassports("04/input", NewPassport2))
	fmt.Printf("Part 2: %4d (duration: %s)\n", result2, duration2)
}

func part1(passports []passport) (int, time.Duration) {
	start := time.Now()

	valid := 0
	for _, p := range passports {
		if p.valid() {
			valid++
		}
	}
	return valid, time.Since(start)
}

func part2(passports []passport) (int, time.Duration) {
	start := time.Now()

	valid := 0
	for _, p := range passports {
		if p.valid() {
			valid++
		}
	}
	return valid, time.Since(start)
}

type passport struct {
	birthYear      int
	issueYear      int
	expirationYear int
	height         int
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

var nineDigitPattern = regexp.MustCompile("^[0-9]{9}$")
var hairColorPattern = regexp.MustCompile("^#[0-9a-f]{6}$")

func NewPassport1(pairs []string) passport {
	passportMap := make(map[string]string, len(pairs))
	for _, field := range pairs {
		pair := strings.Split(field, ":")
		passportMap[pair[0]] = pair[1]
	}

	return passport{
		birthYear:      len(passportMap["byr"]),
		issueYear:      len(passportMap["iyr"]),
		expirationYear: len(passportMap["eyr"]),
		height:         len(passportMap["hgt"]),
		hairColor:      passportMap["hcl"],
		eyeColor:       passportMap["ecl"],
		passportID:     passportMap["pid"],
		countryID:      passportMap["cid"]}
}

func NewPassport2(pairs []string) passport {
	passportMap := make(map[string]string, len(pairs))
	for _, field := range pairs {
		pair := strings.Split(field, ":")
		passportMap[pair[0]] = pair[1]
	}
	p := passport{}

	v := passportMap["byr"]
	if len(v) == 4 {
		byr, err := strconv.Atoi(v)
		if err == nil && byr >= 1920 && byr <= 2002 {
			p.birthYear = byr
		}
	}
	v = passportMap["iyr"]
	if len(v) == 4 {
		iyr, err := strconv.Atoi(v)
		if err == nil && iyr >= 2010 && iyr <= 2020 {
			p.issueYear = iyr
		}
	}
	v = passportMap["eyr"]
	if len(v) == 4 {
		eyr, err := strconv.Atoi(v)
		if err == nil && eyr >= 2020 && eyr <= 2030 {
			p.expirationYear = eyr
		}
	}
	v = passportMap["hgt"]
	if strings.HasSuffix(v, "cm") && len(v) == 5 {
		hgt, err := strconv.Atoi(v[:len(v)-2])
		if err == nil && hgt >= 150 && hgt <= 193 {
			p.height = hgt
		}
	} else if strings.HasSuffix(v, "in") && len(v) == 4 {
		hgt, err := strconv.Atoi(v[:len(v)-2])
		if err == nil && hgt >= 59 && hgt <= 76 {
			p.height = hgt
		}
	}
	v = passportMap["hcl"]
	if hairColorPattern.MatchString(v) {
		p.hairColor = v
	}
	v = passportMap["ecl"]
	if v == "amb" || v == "blu" || v == "brn" || v == "gry" || v == "grn" || v == "hzl" || v == "oth" {
		p.eyeColor = v
	}
	v = passportMap["pid"]
	if nineDigitPattern.MatchString(v) {
		p.passportID = v
	}

	return p
}

func (p passport) valid() bool {
	return p.birthYear > 0 &&
		p.issueYear > 0 &&
		p.expirationYear > 0 &&
		p.height > 0 &&
		len(p.hairColor) > 0 &&
		len(p.eyeColor) > 0 &&
		len(p.passportID) > 0
}

func getPassports(day string, newPassport func([]string) passport) []passport {
	file := util.OpenFile(day)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	passports := make([]passport, 0, 10)

	var passportLine string
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			passportLine += " " + line
		} else {
			fields := strings.Fields(passportLine)
			passports = append(passports, newPassport(fields))
			passportLine = ""
		}

	}
	fields := strings.Fields(passportLine)
	passports = append(passports, newPassport(fields))
	return passports
}
