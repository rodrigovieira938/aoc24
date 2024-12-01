package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func parse_day1_data(str string) [2][]int {
	lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
	var data [2][]int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		numbers := strings.Split(lines[i], "   ")
		number1, err := strconv.Atoi(numbers[0])
		number2, err2 := strconv.Atoi(numbers[1])
		if err == nil && err2 == nil {
			data[0] = append(data[0], number1)
			data[1] = append(data[1], number2)
		} else {
			fmt.Println(err, " ", err2)
		}
	}

	return data
}

func puzzle1(data *[2][]int) int {
	// Sorts the arrays from small to bigger
	sort.Ints(data[0])
	sort.Ints(data[1])

	fmt.Println()

	var distance int = 0

	for i := 0; i < len(data[0]); i++ {
		distance += utils.AbsDiffInt(data[1][i], data[0][i])
	}

	return distance
}

func puzzle2(data *[2][]int) int {
	var similarity_score int = 0
	similarity_map := map[int]int{}
	min := data[1][0]
	max := data[1][0]
	for i := 0; i < len(data[0]); i++ {
		number := data[1][i]
		similarity_map[number]++
		if number < min {
			min = number
		}
		if number > max {
			max = number
		}
	}

	for i := 0; i < len(data[0]); i++ {
		number := data[0][i]
		if number < min {
			continue
		} else if number > max { // Data is ordered so if this number is bigger than the max, every other element in front is bigger
			break
		}

		similarity_score += number * similarity_map[number]
	}

	return similarity_score
}

func main() {

	str, error := utils.ReadFileStr("./data/day1.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		data := parse_day1_data(str) //sorted by puzzle1
		fmt.Printf("Day 1 - Puzzle 1 answer is %v\n", puzzle1(&data))
		fmt.Printf("Day 1 - Puzzle 2 answer is %v\n", puzzle2(&data))
	}
}
