package main

import (
	"fmt"
	"os"
)

type Input struct {
	a, b, c int
	ins     []int
}

func main() {
	input := func() Input {
		file, err := os.Open("day17.in")
		if err != nil {
			panic(err)
		}
		var a, b, c int
		var program string
		_, err = fmt.Fscanf(file, "Register A: %v\nRegister B: %v\nRegister C: %v\n\nProgram: %s\n", &a, &b, &c, &program)
		if err != nil {
			panic(err)
		}

		ins := make([]int, 0)
		for _, c := range program {
			if c >= '0' && c < '9' {
				ins = append(ins, int(c-'0'))
			}
		}
		return Input{a, b, c, ins}
	}()

	part1(input)
	part2(input)
}

func part1(input Input) {
	pc := 0
	out := make([]int, 0)
	for {
		//fmt.Printf("%v %v %v %v %v\n", pc, input.a, input.b, input.c, out)
		if pc < 0 || pc+1 >= len(input.ins) {
			break
		}
		op := func() int {
			return input.ins[pc+1]
		}
		combo_op := func() int {
			switch input.ins[pc+1] {
			case 0:
				return 0
			case 1:
				return 1
			case 2:
				return 2
			case 3:
				return 3
			case 4:
				return input.a % 8
			case 5:
				return input.b % 8
			case 6:
				return input.c % 8
			}
			println(input.ins[pc+1])
			panic("Invalid combo op")
		}

		switch input.ins[pc] {
		case 0: // adv
			input.a = input.a / (1 << combo_op())
			break
		case 1: // bxl
			input.b = input.b ^ op()
			break
		case 2: // bst
			input.b = combo_op() % 8
			break
		case 3: // jnz
			if input.a != 0 {
				pc = op() - 2
			}
			break
		case 4: // bxc
			input.b = input.b ^ input.c
			break
		case 5: // out
			out = append(out, combo_op())
			break
		case 6: // bdv
			input.b = input.a / (1 << combo_op())
			break
		case 7: // cdv
			input.c = input.a / (1 << combo_op())
		default:
			panic("Invalid ins")
		}

		pc += 2
	}

	print("Part 1: ")
	for i, v := range out {
		if i != 0 {
			print(",")
		}
		print(v)
	}
	println()
}

func part2(input Input) {
	values := make([]int, 1)

	for v := range input.ins {
		v = input.ins[len(input.ins)-1-v]

		s := len(values)

		for val := range s {
			val = values[val] << 3
			for i := range 8 {
				var a, b, c int
				a = val + i
				b = a % 8
				b = b ^ 1
				c = a >> b
				b = b ^ 5
				b = b ^ c
				if b%8 == v {
					values = append(values, val+i)
				}
			}
		}

		values = values[s:]
	}

	ans := values[0]
	for _, v := range values {
		ans = min(ans, v)
	}

	println("Part 2:", ans)
}
