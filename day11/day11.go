package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input := func() []int {
		file, err := os.Open("day11.in")
		if err != nil {
			panic(err)
		}
		input := make([]int, 0)
		for {
			var val int
			_, err = fmt.Fscan(file, &val)
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			input = append(input, val)
		}

		return input
	}()

	part1(input)
	part2(input)
}

func part1(input []int) {

	var count_stones func(int, int) int
	count_stones = func(steps, n int) int {
		steps -= 1
		if steps < 0 {
			return 1
		}

		if n == 0 {
			return count_stones(steps, 1)
		}

		d, t := 0, n
		for t != 0 {
			t /= 10
			d += 1
		}
		if d%2 == 0 {
			m := 1
			for range d / 2 {
				m *= 10
			}
			return count_stones(steps, n/m) + count_stones(steps, n%m)
		}

		return count_stones(steps, n*2024)
	}

	ans := 0
	for _, n := range input {
		ans += count_stones(25, n)
	}

	println("Part 1:", ans)
}

func part2(input []int) {
	type Key struct {
		step int
		num  int
	}

	cache := make(map[Key]int)

	var count_stones func(int, int) int
	count_stones = func(steps, n int) int {
		key := Key{
			step: steps,
			num:  n,
		}
		val, ok := cache[key]
		if ok {
			return val
		}

		steps -= 1
		if steps < 0 {
			return 1
		}
		var ans int

		if n == 0 {
			ans = count_stones(steps, 1)
		} else {
			d, t := 0, n
			for t != 0 {
				t /= 10
				d += 1
			}
			if d%2 == 0 {
				m := 1
				for range d / 2 {
					m *= 10
				}
				ans = count_stones(steps, n/m) + count_stones(steps, n%m)
			} else {
				ans = count_stones(steps, n*2024)
			}
		}
		cache[key] = ans
		return ans
	}

	ans := 0
	for _, n := range input {
		ans += count_stones(75, n)
	}

	println("Part 2:", ans)
}
