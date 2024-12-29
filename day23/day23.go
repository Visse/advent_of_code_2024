package main

import (
	"bufio"
	"os"
	"sort"
)

type Edge struct {
	a string
	b string
}

func main() {
	input := func() map[Edge]bool {
		file, err := os.Open("day23.in")
		if err != nil {
			panic(err)
		}

		input := make(map[Edge]bool)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			l := scanner.Text()
			a, b := l[:2], l[3:]
			if l[2] != '-' {
				panic(l)
			}

			input[Edge{a, b}] = true
		}

		return input
	}()

	part1(input)
	part2(input)
}

func part1(input map[Edge]bool) {
	vertices := make(map[string]bool)
	for k := range input {
		vertices[k.a] = true
		vertices[k.b] = true
	}

	ans := 0
	for v1 := range vertices {
		for v2 := range vertices {
			if v1 == v2 {
				continue
			}
			if !(input[Edge{v1, v2}] || input[Edge{v2, v1}]) {
				continue
			}
			for v3 := range vertices {
				if v3 == v1 || v3 == v2 {
					continue
				}
				if !(v1 < v2 && v2 < v3) {
					continue
				}
				if v1[0] != 't' && v2[0] != 't' && v3[0] != 't' {
					continue
				}
				if !(input[Edge{v1, v3}] || input[Edge{v3, v1}]) {
					continue
				}
				if !(input[Edge{v2, v3}] || input[Edge{v3, v2}]) {
					continue
				}
				ans += 1
			}
		}
	}
	println("Part 1:", ans)
}

func part2(input map[Edge]bool) {
	vertices := make([]string, 0)
	seen := make(map[string]bool)

	for e := range input {
		if !seen[e.a] {
			vertices = append(vertices, e.a)
			seen[e.a] = true
		}
		if !seen[e.b] {
			vertices = append(vertices, e.b)
			seen[e.b] = true
		}
	}

	m := make([]string, 0)
	for i := range vertices {
		c := make([]string, 0)
		c = append(c, vertices[i])

		for _, v := range vertices[i+1:] {
			found := true
			for _, o := range c {
				if input[Edge{v, o}] || input[Edge{o, v}] {
					continue
				}
				found = false
				break
			}
			if found {
				c = append(c, v)
			}
		}

		if len(c) > len(m) {
			m = c
		}
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})

	print("Part 2: ")
	for k, c := range m {
		if k != 0 {
			print(",")
		}
		print(c)
	}
	println()
}
