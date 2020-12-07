package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input = [][]string{}

func main() {
	readData()
	part1()
	part2()
}

func part2() {
	validcount := 0

	for _, v := range input {
		positions := strings.Split(v[0], "-")
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		pos1 = pos1 - 1
		pos2 = pos2 - 1
		var indices = []int{}
		for i, val := range v[2] {
			if string(val) == v[1] {
				indices = append(indices, i)
			}
		}
		if (itemExists(indices, pos1) && !itemExists(indices, pos2)) || (!itemExists(indices, pos1) && itemExists(indices, pos2)) {
			validcount++
		}
	}
	fmt.Println("Part2 : Found", validcount, "valid passwords")

}

func itemExists(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func part1() {
	var validpass = [][]string{}
	for _, v := range input {
		minmax := strings.Split(v[0], "-")
		count := strings.Count(v[2], v[1])
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		if count >= min && count <= max {
			validpass = append(validpass, v)
		}
	}
	fmt.Println("Part1 : Found", len(validpass), "valid passwords")
}

func readData() {
	path := "input"

	buf, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = buf.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(buf)
	for snl.Scan() {
		i := string(snl.Text())
		reg, err := regexp.Compile("[^a-zA-Z0-9 -]+")
		if err != nil {
			log.Fatal(err)
		}
		processedString := reg.ReplaceAllString(i, "")

		splits := strings.Split(processedString, " ")
		input = append(input, splits)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
}
