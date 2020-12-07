package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var input = []string{}

func main() {
	loadData()
	part1()
	part2()
}

func part2() {
	count := 0
	groupSize := 0
	set := make(map[rune]int)
	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {

			for _, v := range set {
				if v == groupSize {
					count++
				}
			}

			groupSize = 0
			set = make(map[rune]int)
			continue
		}
		for _, char := range line {
			set[char] = set[char] + 1
		}
		groupSize++
	}
	for _, v := range set {
		if v == groupSize {
			count++
		}
	}
	fmt.Println("Part 2 : ", count)
}

func part1() {
	count := 0
	set := make(map[rune]int)
	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			count += len(set)

			set = make(map[rune]int)
			continue
		}
		for _, char := range line {
			set[char] = set[char] + 1
		}
	}
	count += len(set)
	fmt.Println("Part 1 : ", count)
}

func loadData() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
}
