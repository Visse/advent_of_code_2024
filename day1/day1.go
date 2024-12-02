package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("day1.in")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	input := make([][2]int, 0, 2)

	for {
		var num1, num2 int
		_, err = fmt.Fscanln(file, &num1, &num2)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		input = append(input, [2]int{num1, num2})
	}

	part1(input)
	part2(input)
}

func part1(input [][2]int) {
	var list1, list2 []int

	for _, p := range input {
		list1 = append(list1, p[0])
		list2 = append(list2, p[1])
	}

	slices.Sort(list1)
	slices.Sort(list2)

	ans := 0
	for i := range input {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		ans += diff
	}

	println("Part 1: ", ans)
}

func part2(input [][2]int) {
	ans := 0

	for _, n := range input {
		for _, o := range input {
			if n[0] == o[1] {
				ans += n[0]
			}
		}
	}

	println("Part 2: ", ans)
}
