package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Pos struct {
	x, y int
}
type Machine struct {
	a, b, reward Pos
}
type Data []Machine

func parse_data(lines []string) Data {
	i := 0
	data := Data{}
	for {
		if i+2 >= len(lines) {
			break
		}
		if len(lines[i]) == 0 {
			i++
			continue
		}

		machine := Machine{}

		sep_x_y := strings.Split(lines[i], "Button A: X+")
		if len(sep_x_y) == 2 {
			sep := strings.Split(sep_x_y[1], ", Y+")
			if len(sep) == 2 {
				var x, y int
				var error error
				x, error = strconv.Atoi(sep[0])
				if error != nil {
					i += 3
					continue
				}
				y, error = strconv.Atoi(sep[1])
				if error != nil {
					i += 3
					continue
				}
				machine.a.x = x
				machine.a.y = y
			} else {
				i += 3
				continue
			}
		} else {
			continue
		}
		sep_x_y = strings.Split(lines[i+1], "Button B: X+")
		if len(sep_x_y) == 2 {
			sep := strings.Split(sep_x_y[1], ", Y+")
			if len(sep) == 2 {
				var x, y int
				var error error
				x, error = strconv.Atoi(sep[0])
				if error != nil {
					i += 3
					continue
				}
				y, error = strconv.Atoi(sep[1])
				if error != nil {
					i += 3
					continue
				}
				machine.b.x = x
				machine.b.y = y
			} else {
				i += 3
				continue
			}
		} else {
			i += 3
			continue
		}
		sep_x_y = strings.Split(lines[i+2], "Prize: X=")
		if len(sep_x_y) == 2 {
			sep := strings.Split(sep_x_y[1], ", Y=")
			if len(sep) == 2 {
				var x, y int
				var error error
				x, error = strconv.Atoi(sep[0])
				if error != nil {
					i += 3
					continue
				}
				y, error = strconv.Atoi(sep[1])
				if error != nil {
					i += 3
					continue
				}
				machine.reward.x = x
				machine.reward.y = y
			} else {
				i += 3
				continue
			}
		} else {
			i += 3
			continue
		}
		data = append(data, machine)
		i += 3
	}
	return data
}

func puzzle1(data Data) int {
	get_tokens := func(machine Machine) int {
		// 94A + 22B = 8400
		// 34A + 67B = 5400
		B_float := (float64(machine.reward.y)*float64(machine.a.x) - float64(machine.a.y)*float64(machine.reward.x)) / (float64(machine.b.y)*float64(machine.a.x) - float64(machine.a.y)*float64(machine.b.x))
		A_float := (float64(machine.reward.x) - float64(machine.b.x)*B_float) / float64(machine.a.x)

		A := int(A_float)
		B := int(B_float)

		if (A*machine.a.x+B*machine.b.x) != machine.reward.x || (A*machine.a.y+B*machine.b.y) != machine.reward.y {
			return 0
		}
		return A*3 + B
	}
	sum := 0
	for _, machine := range data {
		tokens := get_tokens(machine)
		sum += tokens
	}
	return sum
}

func main() {
	str, error := utils.ReadFileStr("./data/day13.txt")
	if error == nil {
		data := parse_data(strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n"))
		fmt.Printf("Day 13 - Puzzle 1 answer is %v\n", puzzle1(data))
		for i := range data {
			data[i].reward.x += 10000000000000
			data[i].reward.y += 10000000000000
		}
		fmt.Printf("Day 13 - Puzzle 2 answer is %v\n", puzzle1(data))
	}
}
