package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/affix/AdventOfCOde2020/Day_11/seats"
)

func main() {
	lines := loadData()

	fmt.Println("Part 1:", seats.New(lines, 1, 4).Evolve().CountOccupied())

	fmt.Println("Part 2:", seats.New(lines, -1, 5).Evolve().CountOccupied())
}

func loadData() (input []string) {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return
}
