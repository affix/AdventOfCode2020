package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/set"
)

func main() {

	part1()
	numberOfValidPassports := 0

}

func part2() {
	requiredElements := set.New(set.NonThreadSafe)
	requiredElements.Add("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")

	elementsInThisPassport := set.New(set.NonThreadSafe)

	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
			if elementsMissingFromPassport.IsEmpty() {
				numberOfValidPassports++
			}

			elementsInThisPassport.Clear()
		}

		splitOnSpace := strings.Split(line, " ")
		for _, s := range splitOnSpace {
			if len(s) > 1 {
				kv := strings.Split(s, ":")
				if validateValue(kv[0], kv[1]) {
					fmt.Println(kv[0], kv[1], "VALID")
					elementsInThisPassport.Add(kv[0])
				}
			}
		}
	}

	elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
	if elementsMissingFromPassport.IsEmpty() {
		numberOfValidPassports++
	}

	println(numberOfValidPassports)
}
func part1() {

	numberOfValidPassports := 0

	requiredElements := set.New(set.NonThreadSafe)
	requiredElements.Add("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")

	elementsInThisPassport := set.New(set.NonThreadSafe)

	input, _ := os.Open("input")
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
			if elementsMissingFromPassport.IsEmpty() {
				numberOfValidPassports++
			}

			elementsInThisPassport.Clear()
		}

		splitOnSpace := strings.Split(line, " ")
		for _, s := range splitOnSpace {
			if len(s) > 1 {
				elementsInThisPassport.Add(strings.Split(s, ":")[0])
			}
		}
	}

	elementsMissingFromPassport := set.Difference(requiredElements, elementsInThisPassport)
	if elementsMissingFromPassport.IsEmpty() {
		numberOfValidPassports++
	}

	println(numberOfValidPassports)
}

func validateValue(key string, v string) bool {
	value, _ := strconv.Atoi(v)
	switch key {
	case "byr":
		return value >= 1920 && value <= 2002
	case "iyr":
		return value >= 2010 && value <= 2020
	case "eyr":
		return value >= 2020 && value <= 2030
	case "hgt":
		regex := regexp.MustCompile(`(1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in`)
		return regex.Match([]byte(v))
	case "hcl":
		regex := regexp.MustCompile(`#[0-9a-fA-F]{6}`)
		return regex.Match([]byte(v))
	case "ecl":
		regex := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
		return regex.Match([]byte(v))
	case "pid":
		regex := regexp.MustCompile(`[0-9]{9}`)
		return regex.Match([]byte(v))
	case "cid":
		return true
	}
	return false
}
