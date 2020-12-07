package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var target = 2020
var input = []int{}

func main() {
	readData()
	i, j := sums(input)
	n1 := input[i]
	n2 := input[j]

	fmt.Println("Part 1 : ", n1*n2)

	w, v, y := sums2(input)
	n3 := input[w]
	n4 := input[v]
	n5 := input[y]

	fmt.Println("Part 2 : ", n3*n4*n5)
}

func sums(nums []int) (int, int) {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return j, i
		}
		m[target-v] = i
	}
	return -1, -1
}

func sums2(nums []int) (int, int, int) {
	m := make(map[int]int)
	for i, v := range nums {
		for j, w := range nums {
			if k, ok := m[w]; ok {
				return j, i, k
			}
			m[target-w-v] = j
		}
		m[target-v] = i
	}
	return -1, -1, -1
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
		i, _ := strconv.Atoi(snl.Text())
		input = append(input, i)
	}
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}
}
