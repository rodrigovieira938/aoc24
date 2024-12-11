package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func parse_data(str string) []string {
	numbers := []string{}
	for _, n := range strings.Split(str, " ") {
		if utils.IsStrInt(n) {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func blink(input []string) []string {
	output := []string{}
	for _, number := range input {
		if number == "0" {
			output = append(output, "1")
		} else if len(number)%2 == 0 {
			first_half := number[0 : len(number)/2]
			n, _ := strconv.Atoi(number[len(number)/2:])
			second_half := strconv.Itoa(n) //easy way to remove trailing 0
			output = append(output, first_half, second_half)
		} else {
			n, _ := strconv.Atoi(number)
			n *= 2024
			output = append(output, strconv.Itoa(n))
		}
	}
	return output
}
func puzzle1(data []string) int {
	for i := 0; i < 25; i++ {
		data = blink(data)
	}
	return len(data)
}

func main() {
	str, error := utils.ReadFileStr("./data/day11.txt")
	if error == nil {
		numbers := parse_data(strings.ReplaceAll(strings.ReplaceAll(str, "\r\n", ""), "\n", ""))
		fmt.Printf("Day 11 - Puzzle 1 answer is %v\n", puzzle1(numbers))
	}
}
