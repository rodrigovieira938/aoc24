package main

import (
	"fmt"
	"strconv"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func puzzle1(str string) int {
	var sum int = 0
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		for str[i] != 'm' {
			if i+1 < strlen {
				i++
			} else {
				break
			}
		}
		if i+3 < strlen {
			if str[i:i+4] == "mul(" {
				i += 4
				start := i
				for utils.IsCharNum(str[i]) {
					i++
				}
				end := i
				if str[i] == ',' {
					i++
				}
				start2 := i
				for utils.IsCharNum(str[i]) {
					i++
				}
				end2 := i
				if str[i] == ')' {
					num1, _ := strconv.Atoi(str[start:end])
					num2, _ := strconv.Atoi(str[start2:end2])

					sum += num1 * num2
				}
			} else {
				i++
			}
		}
	}

	return sum
}

func main() {
	str, err := utils.ReadFileStr("./data/day3.txt")
	if err == nil {
		sum := puzzle1(str)
		fmt.Printf("Day 3 - Puzzle 1 answer is %v\n", sum)
	}
}
