package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var input = []int{}
var inv = 0
var preamble = 25

func main() {
	loadData()
	part1()
	part2()
}

func part1() {
loop:
	for i := preamble; i < len(input); i++ {
		for j := i - preamble; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if input[j]+input[k] == input[i] {
					continue loop
				}
			}
		}
		inv = input[i]
		break
	}

	fmt.Println("Part 1: ", inv)
}

func part2() {
	for i := preamble; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			sum := 0
			for _, v := range input[i : j+1] {
				sum += v
			}
			if sum == inv {
				sort.Ints(input[i : j+1])
				fmt.Println("Part 2: ", input[i]+input[j])
				break
			}
		}
	}

}

func loadData() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		input = append(input, line)
	}
}
