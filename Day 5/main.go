package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var boardingPass = []string{}

func main() {
	loadData()
	part1()
	part2()
}

func loadData() {
	file, _ := os.Open("input")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		boardingPass = append(boardingPass, line)
	}
}

func part1() {
	highestSeat := 0
	for _, line := range boardingPass {
		minRow := 0
		maxRow := 128
		minCol := 0
		maxCol := 8
		for _, v := range line {
			if v == 'F' {
				maxRow -= (maxRow - minRow) / 2
			} else if v == 'B' {
				minRow += (maxRow - minRow) / 2
			} else if v == 'R' {
				minCol += (maxCol - minCol) / 2
			} else {
				maxCol -= (maxCol - minCol) / 2
			}
		}
		seat := minRow*8 + minCol

		if seat > highestSeat {
			highestSeat = seat
		}
	}
	fmt.Println("Part 1 : ", highestSeat)
}

func part2() {
	occupiedSeats := []int{}

	for _, line := range boardingPass {
		minRow := 0
		maxRow := 128
		minCol := 0
		maxCol := 8
		for _, v := range line {
			if v == 'F' {
				maxRow -= (maxRow - minRow) / 2
			} else if v == 'B' {
				minRow += (maxRow - minRow) / 2
			} else if v == 'R' {
				minCol += (maxCol - minCol) / 2
			} else {
				maxCol -= (maxCol - minCol) / 2
			}
		}
		seat := minRow*8 + minCol
		occupiedSeats = append(occupiedSeats, seat)
	}
	sort.Ints(occupiedSeats)
	previous := occupiedSeats[0] - 1
	for _, seat := range occupiedSeats {
		if previous+1 != seat {
			fmt.Println("Part 2 : ", previous+1)
			break
		}
		previous = seat
	}
}
