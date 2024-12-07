package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Rule struct {
	result int
	values []int
}

func main() {
	input := func() []Rule {
		file, err := os.Open("day7.in")
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(file)

		input := make([]Rule, 0, 100)

		for {
			rule := Rule{
				result: 0,
				values: make([]int, 0),
			}

			_, err = fmt.Fscanf(reader, "%d:", &rule.result)
			if err == io.EOF {
				return input
			}

			for {
				var value int
				fmt.Fscanf(reader, "%d", &value)
				rule.values = append(rule.values, value)

				b, err := reader.ReadByte()
				if err != nil {
					panic(err)
				}
				if b == '\n' {
					break
				}
				reader.UnreadByte()
			}
			input = append(input, rule)
		}
	}()

	part1(input)
	part2(input)
}

func can_solve(rule Rule, tot int, i int, part2 bool) bool {
	if tot > rule.result {
		return false
	}
	i += 1
	if i >= len(rule.values) {
		return tot == rule.result
	}
	if can_solve(rule, tot*rule.values[i], i, part2) {
		return true
	}
	if can_solve(rule, tot+rule.values[i], i, part2) {
		return true
	}
	if part2 {
		v := rule.values[i]
		for v > 0 {
			tot *= 10
			v /= 10
		}
		tot += rule.values[i]
		return can_solve(rule, tot, i, part2)
	}
	return false
}

func part1(input []Rule) {
	ans := 0
	for _, r := range input {
		if can_solve(r, r.values[0], 0, false) {
			ans += r.result
		}
	}
	println("Part 1:", ans)
}

func part2(input []Rule) {
	ans := 0
	for _, r := range input {
		if can_solve(r, r.values[0], 0, true) {
			ans += r.result
		}
	}
	println("Part 2:", ans)

}
