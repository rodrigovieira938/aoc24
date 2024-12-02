package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func parse_day2_data(str string) [][]int {
	levels := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	var data [][]int
	for _, level := range levels {
		if level == "" {
			continue
		}

		var numbers []int
		for _, number_str := range strings.Split(level, " ") {
			number, err := strconv.Atoi(number_str)
			if err == nil {
				numbers = append(numbers, number)
			}
		}
		data = append(data, numbers)
	}
	return data
}

func puzzle1(data [][]int) int {
	var safe int = 0

	for _, level := range data {
		level_len := len(level)
		is_safe := true
		is_increasing := level[0] < level[1]

		for i, number := range level {
			if i+1 < level_len {
				number2 := level[i+1]

				diff := utils.AbsDiffInt(number, number2)

				if diff <= 0 || diff > 3 {
					is_safe = false
					break
				}
				if is_increasing && level[i] > level[i+1] ||
					!is_increasing && level[i] < level[i+1] {
					is_safe = false
					break
				}
			}
		}
		if is_safe {
			safe++
		}
	}
	return safe
}

func main() {
	str, error := utils.ReadFileStr("./data/day2.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		data := parse_day2_data(str)
		safe_count := puzzle1(data)
		fmt.Printf("Day 1 - Puzzle 1 answer is %v\n", safe_count)
	}
}
