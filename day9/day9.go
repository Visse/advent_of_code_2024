package main

import (
	"os"
)

func main() {
	input := func() []int {
		raw_input, err := os.ReadFile("day9.in")
		if err != nil {
			panic(err)
		}

		input := make([]int, 0, len(raw_input)-1)
		for _, c := range raw_input[:len(raw_input)-1] {
			input = append(input, int(c-'0'))
		}

		return input
	}()

	part1(input)
	part2(input)
}

func part1(input []int) {
	size := 0
	for _, v := range input {
		size += v
	}
	disk := make([]int, size)

	p := 0
	for id, v := range input {
		if id%2 == 1 {
			id = -1
		} else {
			id /= 2
		}

		for range v {
			disk[p] = id
			p += 1
		}
	}

	s := 0
	e := len(disk) - 1

	for {
		for disk[s] != -1 {
			s += 1
		}
		for disk[e] == -1 {
			e -= 1
		}
		if e <= s {
			break
		}
		disk[s] = disk[e]
		disk[e] = -1
	}

	ans := 0
	for i, v := range disk[:e+1] {
		ans += i * v
	}
	println("Part 1:", ans)
}

func part2(input []int) {
	type Slot struct {
		loc  int
		size int
	}

	files := make([]Slot, 0)
	free := make([]Slot, 0)
	/*
		print_disk := func() {
			size := 0
			for _, v := range input {
				size += v
			}
			disk := make([]int, size)
			for i := range disk {
				disk[i] = -1
			}
			for id, file := range files {
				for b := range file.size {
					loc := file.loc + b
					if disk[loc] != -1 {
						panic("overlap!")
					}
					disk[loc] = id
				}
			}
			for _, v := range disk {
				if v == -1 {
					print(".")
				} else {
					print("X")
				}
			}
			println()
		}
	*/
	p := 0
	for id, v := range input {
		if id%2 == 0 {
			files = append(files, Slot{loc: p, size: v})
		} else {
			free = append(free, Slot{loc: p, size: v})
		}
		p += v
	}

	for i := range files {
		i = len(files) - 1 - i

		for p := range free[:i] {
			if files[i].size <= free[p].size {
				files[i].loc = free[p].loc
				free[p].loc += files[i].size
				free[p].size -= files[i].size
				break
			}
		}
	}
	//print_disk()
	ans := 0
	for id, file := range files {
		for b := range file.size {
			ans += (file.loc + b) * id
		}
	}
	println("Part 2:", ans)

	//fmt.Printf("%v", free)

}
