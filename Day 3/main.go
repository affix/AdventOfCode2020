package main

import (
	"bufio"
	"fmt"
	"os"
)

var data = []string{}
var trees = 0
var x = 0
var width = 0

type slope struct {
	x, y int
}

func main() {
	data = readFile("input", func(i int, text string) string {
		return text
	})
	width = len(data[0])
	part1()
	part2()

}

func part1() {
	for y := range data {
		if isTree(data, x, y, width) {
			trees++
		}
		x += 3
	}

	fmt.Printf("Part 1: %d\n", trees)
}

func part2() {
	slopes := []slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	allSlopes := []int{}

	for _, slope := range slopes {
		trees = 0
		x = 0
		y := 0
		for y < len(data) {
			if isTree(data, x, y, width) {
				trees++
			}
			x += slope.x
			y += slope.y
		}

		allSlopes = append(allSlopes, trees)
	}

	product := 1
	for _, num := range allSlopes {
		product *= num
	}
	fmt.Printf("Part 2: %d\n", product)
}

func readFile(path string, process func(int, string) string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, process(0, scanner.Text()))
	}

	return result
}

func isTree(data []string, x int, y int, width int) bool {
	return data[y][x%width] == '#'
}
