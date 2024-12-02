package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.in")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	input := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		num := []int{}
		for _, n := range fields {
			i, _ := strconv.Atoi(n)
			num = append(num, i)
		}

		input = append(input, num)
	}

	part1(input)
	part2(input)
}

func part1(input [][]int) {
	pass := 0

report:
	for _, r := range input {
		prev := 0
		d := false
		for i, val := range r {
			if i == 0 {
				prev = val
				continue
			} else if i == 1 {
				d = val < prev
			}

			if d != (val < prev) {
				//fmt.Printf("Unsafe d: %v\n", r)
				continue report
			}
			diff := val - prev
			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				//fmt.Printf("Unsafe: %v\n", r)
				continue report
			}
			prev = val
		}
		//fmt.Printf("Safe: %v\n", r)
		pass += 1
	}
	println("Part 1:", pass)
}

func part2(input [][]int) {
	pass := 0

	check_report := func(skip int, r []int) bool {
		prev := 0
		d := false
		for i, val := range r {
			if i == skip {
				continue
			}
			if skip < i {
				i -= 1
			}

			if i == 0 {
				prev = val
				continue
			} else if i == 1 {
				d = val < prev
			}

			if d != (val < prev) {
				//fmt.Printf("Unsafe d: %v\n", r)
				return false
			}
			diff := val - prev
			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				//fmt.Printf("Unsafe: %v\n", r)
				return false
			}
			prev = val
		}
		//fmt.Printf("Safe: %v\n", r)
		return true
	}

report:
	for _, r := range input {
		if check_report(-1, r) {
			pass += 1
			continue report
		}

		for skip := range r {
			if check_report(skip, r) {
				pass += 1
				continue report
			}
		}
	}
	println("Part 1:", pass)

}
