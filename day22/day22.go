package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := func() []int {
		file, err := os.Open("day22.in")
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(file)

		input := make([]int, 0)
		for {
			var i int
			_, err := fmt.Fscanf(reader, "%v\n", &i)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			input = append(input, i)
		}

		return input
	}()

	part1(input)
	part2(input)
}

func part1(input []int) {
	next := func(n int) int {
		n = ((n * 64) ^ n) % 16777216
		n = ((n / 32) ^ n) % 16777216
		n = ((n * 2048) ^ n) % 16777216
		return n
	}

	ans := 0
	for _, n := range input {
		for range 2000 {
			n = next(n)
		}
		ans += n
	}

	println("Part 1:", ans)
}

func part2(input []int) {
	next := func(n int) int {
		n = ((n * 64) ^ n) % 16777216
		n = ((n / 32) ^ n) % 16777216
		n = ((n * 2048) ^ n) % 16777216
		return n
	}

	count := make(map[int]int)
	for _, n := range input {
		advance := func() int {
			n = next(n)
			return n % 10
		}
		a := n % 10
		b := advance()
		c := advance()
		d := advance()

		f := make(map[int]bool)

		for range 2000 - 4 {
			e := advance()

			k1 := b - a
			k2 := c - b
			k3 := d - c
			k4 := c - e
			k := (k1+10)*100 + (k2+10)*10000 + (k3+10)*1000000 + (k4+10)*100000000
			if !f[k] {
				count[k] += e
				f[k] = true
			}
			a, b, c, d = b, c, d, e
		}
	}

	ans := 0
	for _, v := range count {
		ans = max(ans, v)
	}
	println(len(count))

	println("Part 2:", ans)
}
