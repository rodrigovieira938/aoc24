package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Pos struct {
	x int
	y int
}

func is_frequency_r(c rune) bool {
	return unicode.IsLower(c) || unicode.IsUpper(c) || unicode.IsDigit(c)
}

func insert_antinode(pos Pos, data []string) {
	if pos.y < len(data) && pos.y >= 0 {
		if pos.x < len(data[pos.y]) && pos.x >= 0 {
			data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '#', pos.x)
		}
	}
}

func puzzle1(data []string) int {
	frequencies := map[rune][]Pos{}

	for fy, line := range data {
		for fx, char := range line {
			if is_frequency_r(char) {
				frequencies[char] = append(frequencies[char], Pos{fx, fy})
			}
		}
	}

	for _, positions := range frequencies {
		for x, pos1 := range positions {
			for y, pos2 := range positions {
				if x != y {
					insert_antinode(Pos{pos1.x + (pos1.x - pos2.x), pos1.y + (pos1.y - pos2.y)}, data)
					insert_antinode(Pos{pos2.x + (pos2.x - pos1.x), pos2.y + (pos2.y - pos1.y)}, data)
				}
			}
		}
	}
	unique_positions := 0
	for _, line := range data {
		fmt.Println(line)
		for _, c := range line {
			if c == '#' {
				unique_positions++
			}
		}
	}
	return unique_positions
}

func main() {
	str, error := utils.ReadFileStr("./data/day8.txt")
	if error == nil {
		data := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
		fmt.Printf("Day 8 - Puzzle 1 answer is %v\n", puzzle1(data))
	}
}
