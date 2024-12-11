package main

import (
	"fmt"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Pos struct {
	x int
	y int
}

func puzzle1(data []string) int {
	starting_points := []Pos{}

	for y, line := range data {
		for x, c := range line {
			if c == '0' {
				starting_points = append(starting_points, Pos{x, y})
			}
		}
	}
	var walkthrough func(pos Pos) map[Pos]int
	walkthrough = func(pos Pos) map[Pos]int {
		exit_positions := map[Pos]int{}

		number := data[pos.y][pos.x]
		checkpos := func(pos Pos) bool {
			if pos.y >= 0 && pos.y < len(data) {
				line := data[pos.y]
				if pos.x >= 0 && pos.x < len(line) {
					c := line[pos.x]
					if c == number+1 {
						return true
					}
				}
			}
			return false
		}
		changed_pos := true
		var pos_to_change Pos

		if number == '9' {
			exit_positions[pos]++
		}

		possible_positions := []Pos{{pos.x, pos.y + 1}, {pos.x, pos.y - 1}, {pos.x + 1, pos.y}, {pos.x - 1, pos.y}}

		for _, p := range possible_positions {
			if checkpos(p) {
				if data[p.y][p.x] == number+1 {
					if changed_pos {
						for k, v := range walkthrough(p) {
							exit_positions[k] += v
						}
					} else {
						changed_pos = true
						pos_to_change = p
					}
				}
			}
		}
		pos = pos_to_change
		return exit_positions
	}
	score := 0
	for _, start_pos := range starting_points {
		exit_positions := walkthrough(start_pos)
		score += len(exit_positions)
	}
	return score
}
func puzzle2(data []string) int {
	starting_points := []Pos{}

	for y, line := range data {
		for x, c := range line {
			if c == '0' {
				starting_points = append(starting_points, Pos{x, y})
			}
		}
	}
	var walkthrough func(pos Pos) map[Pos]int
	walkthrough = func(pos Pos) map[Pos]int {
		exit_positions := map[Pos]int{}

		number := data[pos.y][pos.x]
		checkpos := func(pos Pos) bool {
			if pos.y >= 0 && pos.y < len(data) {
				line := data[pos.y]
				if pos.x >= 0 && pos.x < len(line) {
					c := line[pos.x]
					if c == number+1 {
						return true
					}
				}
			}
			return false
		}
		changed_pos := true
		var pos_to_change Pos

		if number == '9' {
			exit_positions[pos]++
		}

		possible_positions := []Pos{{pos.x, pos.y + 1}, {pos.x, pos.y - 1}, {pos.x + 1, pos.y}, {pos.x - 1, pos.y}}

		for _, p := range possible_positions {
			if checkpos(p) {
				if data[p.y][p.x] == number+1 {
					if changed_pos {
						for k, v := range walkthrough(p) {
							exit_positions[k] += v
						}
					} else {
						changed_pos = true
						pos_to_change = p
					}
				}
			}
		}
		pos = pos_to_change
		return exit_positions
	}
	score := 0
	for _, start_pos := range starting_points {
		exit_positions := walkthrough(start_pos)
		for _, v := range exit_positions {
			score += v
		}
	}
	return score
}
func main() {
	str, error := utils.ReadFileStr("./data/day10.txt")
	if error == nil {
		data := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
		fmt.Printf("Day 9 - Puzzle 1 answer is %v\n", puzzle1(data))
		fmt.Printf("Day 9 - Puzzle 2 answer is %v\n", puzzle2(data))
	}
}
