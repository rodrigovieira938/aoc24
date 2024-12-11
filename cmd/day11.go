package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func parse_data(str string) []int {
	numbers := []int{}
	for _, n := range strings.Split(str, " ") {
		number, error := strconv.Atoi(n)
		if error == nil {
			numbers = append(numbers, number)
		}
	}
	return numbers
}

type number_gen struct {
	number int
	gen    int
}

var memoization map[number_gen]int = map[number_gen]int{}

func blink_number(number int, generation int) int {
	if generation == 0 {
		return 1
	}
	if stone_count, found := memoization[number_gen{number, generation}]; found {
		return stone_count
	}

	stone_count := 0
	if number == 0 {
		stone_count = blink_number(1, generation-1)
	} else {
		number_str := strconv.Itoa(number)
		if len(number_str)%2 == 0 {
			first_half, _ := strconv.Atoi(number_str[0 : len(number_str)/2])
			second_half, _ := strconv.Atoi(number_str[len(number_str)/2:])
			stone_count = blink_number(first_half, generation-1) + blink_number(second_half, generation-1)
		} else {
			stone_count = blink_number(number*2024, generation-1)
		}
	}
	memoization[number_gen{number, generation}] = stone_count
	return stone_count
}

func blink(input []int) []int {
	output := []int{}
	for _, number := range input {
		if number == 0 {
			output = append(output, 1)
		} else {
			number_str := strconv.Itoa(number)
			if len(number_str)%2 == 0 {
				first_half, _ := strconv.Atoi(number_str[0 : len(number_str)/2])
				second_half, _ := strconv.Atoi(number_str[len(number_str)/2:])
				output = append(output, first_half, second_half)
			} else {
				output = append(output, number*2024)
			}
		}
	}
	return output
}
func puzzle1(data []int) int {
	for i := 0; i < 25; i++ {
		data = blink(data)
	}
	return len(data)
}
func puzzle2(data []int) int {
	count := 0
	for _, number := range data {
		count += blink_number(number, 75)
	}
	return count
}

func main() {
	str, error := utils.ReadFileStr("./data/day11.txt")
	if error == nil {
		numbers := parse_data(strings.ReplaceAll(strings.ReplaceAll(str, "\r\n", ""), "\n", ""))
		fmt.Printf("Day 11 - Puzzle 1 answer is %v\n", puzzle1(numbers))
		fmt.Printf("Day 11 - Puzzle 2 answer is %v\n", puzzle2(numbers))
	}
}
