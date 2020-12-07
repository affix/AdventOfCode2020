package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	ContainedBy []string
	Contains    map[string]int
}

var input = []string{}

func main() {
	loadData()
	rules := parseRules(input)
	fmt.Println("Part 1:", part1(rules))
	fmt.Println("Part 2:", part2(rules))
}

func part1(bagContainsInfo map[string]Bag) int {
	bagToBagsContaining := make(map[string][]string)
	for key, value := range bagContainsInfo {
		for bagContainedBy := range value.Contains {
			bagToBagsContaining[bagContainedBy] = append(bagToBagsContaining[bagContainedBy], key)
		}
	}

	return len(recursiveSearchBagsEventuallyContaining(bagToBagsContaining, "shiny gold"))
}

func part2(bagContainsInfo map[string]Bag) (numBags int) {
	numBags = recursiveSearchHowManyBagsContained(bagContainsInfo, "shiny gold")

	return
}

func recursiveSearchHowManyBagsContained(bagContainsInfo map[string]Bag, color string) (count int) {
	for bagType, num := range bagContainsInfo[color].Contains {
		count += num + (num * recursiveSearchHowManyBagsContained(bagContainsInfo, bagType))
	}

	return
}

func recursiveSearchBagsEventuallyContaining(bagToBagsContaining map[string][]string, color string) (items map[string]bool) {
	items = make(map[string]bool)
	if _, present := bagToBagsContaining[color]; present {
		for _, containingColor := range bagToBagsContaining[color] {
			items[containingColor] = true
			for item := range recursiveSearchBagsEventuallyContaining(bagToBagsContaining, containingColor) {
				items[item] = true
			}
		}
	}

	return
}

func parseRules(lines []string) (bags map[string]Bag) {
	bags = make(map[string]Bag)
	for _, line := range lines {
		newBag := Bag{}
		newBag.Contains = make(map[string]int)

		lineMatches := regexp.MustCompile("^(.+) bags contain (.+)$").FindStringSubmatch(line)
		containsList := strings.Split(lineMatches[2], ",")
		for _, item := range containsList {
			if strings.Contains(item, "no other bags") {
				continue
			}
			containsMatches := regexp.MustCompile("(\\d+) (.+) bag").FindStringSubmatch(item)

			count, _ := strconv.Atoi(containsMatches[1])
			newBag.Contains[containsMatches[2]] = count
		}

		if lineMatches != nil {
			bags[lineMatches[1]] = newBag
		}
	}

	return
}

func loadData() {
	file, _ := os.Open("input")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
}
