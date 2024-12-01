package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

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

func main() {

	str, error := utils.ReadFileStr("./data/day1_1.txt")
	if error != nil {
		fmt.Println(error)
	} else {
		lines := strings.Split(str, "\n")

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
			}
		}

		fmt.Printf("The difference is %v\n", puzzle1(&data))
	}
}
