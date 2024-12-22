package main

import (
	"fmt"
	"os"
)

type Input struct {
	t, p []string
}

func main() {
	input := func() Input {
		file, err := os.ReadFile("day19.in")
		if err != nil {
			panic(err)
		}

		t := make([]string, 0)
		for {
			if file[0] == '\n' {
				break
			}
			i := 0
			for ; file[i] >= 'a' && file[i] <= 'z'; i++ {
			}

			t = append(t, string(file[:i]))
			file = file[i:]

			for file[0] == ' ' || file[0] == ',' {
				file = file[1:]
			}
		}

		p := make([]string, 0)
		for len(file) != 0 {
			i := 0
			for ; file[i] != '\n'; i++ {
			}
			if i != 0 {
				p = append(p, string(file[:i]))
			}
			file = file[i+1:]
		}

		fmt.Printf("%v\n", string(file))
		return Input{t, p}
	}()

	part1(input)
	part2(input)
}

func part1(input Input) {
	var is_possible func(s string) bool
	is_possible = func(s string) bool {
		if len(s) == 0 {
			return true
		}

		for _, t := range input.t {
			if len(s) < len(t) {
				continue
			}
			if t == s[:len(t)] {
				if is_possible(s[len(t):]) {
					return true
				}
			}
		}
		return false
	}

	ans := 0
	for _, p := range input.p {
		if is_possible(p) {
			ans++
		}
	}
	println("Part 1:", ans)
}

func part2(input Input) {
	var is_possible func(s string) int
	ways := make(map[string]int)
	is_possible = func(s string) int {
		if len(s) == 0 {
			return 1
		}
		w, e := ways[s]
		if e {
			return w
		}

		i := 0
		for _, t := range input.t {
			if len(s) < len(t) {
				continue
			}
			if t == s[:len(t)] {
				i += is_possible(s[len(t):])
			}
		}
		ways[s] = i
		return i
	}

	ans := 0
	for _, p := range input.p {
		ans += is_possible(p)
	}
	println("Part 2:", ans)

}
