package main

import (
	"os"
)

func main() {
	input, err := os.ReadFile("day3.in")
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func part1(input []byte) {
	cur := 0

	take := func() (int, bool) {
		if cur < len(input) {
			c := int(input[cur])
			cur += 1
			return c, false
		}
		return 0, true

	}

	take_num := func() (int, bool) {
		num, err := take()
		if err {
			return 0, true
		}
		num -= int('0')
		if num < 0 || num > 9 {
			return 0, true
		}
		for {
			d, err := take()
			if err {
				return num, false
			}
			d -= '0'
			if d < 0 || d > 9 {
				cur -= 1
				return num, false
			}
			num = num*10 + d
		}
	}
	expect_byte := func(b byte) bool {
		n, err := take()
		if err {
			return false
		}
		return n == int(b)
	}
	expect_str := func(s string) bool {
		for _, c := range []byte(s) {
			if !expect_byte(c) {
				return false
			}
		}
		return true
	}

	ans := 0
	for cur < len(input) {
		if expect_str("mul(") {
			num1, err := take_num()
			if err {
				println("Failed to parse 1st")
				continue
			}
			if !expect_byte(',') {
				continue
			}
			num2, err := take_num()
			if err {
				println("Failed to parse 2nd")
				continue
			}
			if !expect_byte(')') {
				continue
			}
			//println(num1, "*", num2)
			ans += num1 * num2
		}
	}
	println("Part 1:", ans)
}

func part2(input []byte) {
	cur := 0

	take := func() (int, bool) {
		if cur < len(input) {
			c := int(input[cur])
			cur += 1
			return c, false
		}
		return 0, true
	}

	take_num := func() (int, bool) {
		num, err := take()
		if err {
			return 0, true
		}
		num -= int('0')
		if num < 0 || num > 9 {
			return 0, true
		}
		for {
			d, err := take()
			if err {
				return num, false
			}
			d -= '0'
			if d < 0 || d > 9 {
				cur -= 1
				return num, false
			}
			num = num*10 + d
		}
	}
	expect_byte := func(b byte) bool {
		n, err := take()
		if err {
			return false
		}
		return n == int(b)
	}
	expect_str := func(s string) bool {
		for _, c := range []byte(s) {
			if !expect_byte(c) {
				return false
			}
		}
		return true
	}

	ans := 0
	for cur < len(input) {
		save := cur
		if !expect_str("don't()") {
			cur = save
		} else {
			// Disabled, scan to next do()
			for cur < len(input) && !expect_str("do()") {
			}
		}
		if expect_str("mul(") {
			num1, err := take_num()
			if err {
				println("Failed to parse 1st")
				continue
			}
			if !expect_byte(',') {
				continue
			}
			num2, err := take_num()
			if err {
				println("Failed to parse 2nd")
				continue
			}
			if !expect_byte(')') {
				continue
			}
			//println(num1, "*", num2)
			ans += num1 * num2
		}
	}
	println("Part 2:", ans)
}
