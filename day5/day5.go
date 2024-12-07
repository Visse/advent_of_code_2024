package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	rules, updates := func() ([][2]int, [][]int) {
		file, err := os.Open("day5.in")
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(file)

		rules := make([][2]int, 0)

		for {
			var a, b int
			n, err := fmt.Fscanf(reader, "%d|%d\n", &a, &b)
			if err != nil {
				break
			}
			if n == 0 {
				break
			}

			rules = append(rules, [2]int{a, b})
		}

		updates := make([][]int, 0)
	outer:
		for {
			update := make([]int, 0)
			for {
				var a int
				n, err := fmt.Fscanf(reader, "%d", &a)
				if err == io.EOF {
					break outer
				}
				if err != nil {
					panic(err)
				}
				if n == 0 {
					break
				}
				update = append(update, a)
				b, err := reader.ReadByte()
				if err != nil {
					panic(err)
				}
				if b == '\n' {
					break
				}
				if b != ',' {
					panic("Unexpected byte")
				}

			}
			updates = append(updates, update)
		}

		return rules, updates
	}()

	ans := 0

	incorect_updates := make([][]int, 0)
	for _, update := range updates {
		correct := true
	outer:
		for i, d := range update {
			for _, c := range update[i+1:] {
				for _, r := range rules {
					if r[1] == d && r[0] == c {
						correct = false
						incorect_updates = append(incorect_updates, update)

						break outer
					}
				}
			}
		}

		if correct {
			ans += update[len(update)/2]
		}
	}

	println("Part 1:", ans)

	ans = 0
	for _, update := range incorect_updates {

		correct := false
		for correct == false {
			correct = true
		outer_2:
			for di, d := range update {
				// fmt.Printf("%v %v\n", update, update[di+1:])
				for ci, c := range update[di+1:] {
					for _, r := range rules {
						if r[1] == d && r[0] == c {
							// fmt.Printf("Swap %d %d: %v - %v\n", d, c, r, update)
							// println("Swap", d, c)
							update[di] = c
							update[ci+di+1] = d
							correct = false
							break outer_2
						}
					}
				}
			}
		}

		ans += update[len(update)/2]
	}
	println("Part 2:", ans)

}
