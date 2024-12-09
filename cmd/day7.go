package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Equation struct {
	result  int
	numbers []int
}

type Data []Equation

func parse_data(str string) Data {
	var data Data

	equations := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

	for _, equation_raw := range equations {
		split := strings.Split(equation_raw, ":")
		var equation Equation

		if len(split) != 2 {
			continue
		}
		var err error
		equation.result, err = strconv.Atoi(split[0])
		if err != nil {
			continue
		}
		numbers := strings.Split(split[1], " ")
		for _, n := range numbers {
			var number int
			number, err = strconv.Atoi(n)
			if err != nil {
				continue
			}
			equation.numbers = append(equation.numbers, number)
		}
		data = append(data, equation)
	}

	return data
}

func generate_possibilities(number int, operators []rune) []string {
	var recurse func(int, []rune, []string) []string
	recurse = func(number int, operators []rune, prev []string) []string {
		if number == 0 {
			return prev
		}

		var possibilities []string

		for _, eq := range prev {
			for _, op := range operators {
				var builder strings.Builder
				builder.WriteString(eq)
				builder.WriteRune(op)
				possibilities = append(possibilities, builder.String())
			}
		}
		return recurse(number-1, operators, possibilities)
	}
	clone := make([]string, len(operators))
	for i, op := range operators {
		clone[i] = string(op)
	}
	return recurse(number-1, operators, clone)
}

func puzzle1(data Data) int {
	calibration_sum := 0
	for _, eq := range data {
		possibilities := generate_possibilities(len(eq.numbers)-1, []rune{'+', '*'})
		for _, possibility := range possibilities {
			number_index := 0
			first := true
			res := 0
			for _, op := range possibility {
				if first {
					res = eq.numbers[number_index]
					number_index++
					first = false
				}
				if op == '+' {
					res += eq.numbers[number_index]
				} else if op == '*' {
					res *= eq.numbers[number_index]
				}
				number_index++
			}
			if res == eq.result {
				calibration_sum += res
				break
			}
		}
	}
	return calibration_sum
}
func puzzle2(data Data) int {
	calibration_sum := 0
	for _, eq := range data {
		//Concatenation operator becomes | instead of || for simplicity
		possibilities := generate_possibilities(len(eq.numbers)-1, []rune{'+', '*', '|'})
		for _, possibility := range possibilities {
			number_index := 0
			first := true
			res := 0
			for _, op := range possibility {
				if first {
					res = eq.numbers[number_index]
					number_index++
					first = false
				}
				if op == '+' {
					res += eq.numbers[number_index]
				} else if op == '*' {
					res *= eq.numbers[number_index]
				} else if op == '|' {
					builder := strings.Builder{}
					builder.WriteString(strconv.Itoa(res))
					builder.WriteString(strconv.Itoa(eq.numbers[number_index]))

					res, _ = strconv.Atoi(builder.String()) //won't fail
				}
				number_index++
			}
			if res == eq.result {
				calibration_sum += res
				break
			}
		}
	}
	return calibration_sum
}

func main() {
	str, error := utils.ReadFileStr("./data/day7.txt")
	if error == nil {
		data := parse_data(str)
		fmt.Printf("Day 7 - Puzzle 1 answer is %v\n", puzzle1(data))
		fmt.Printf("Day 7 - Puzzle 2 answer is %v\n", puzzle2(data))
	}
}
