package main

import (
	"fmt"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func diagonal_rightdown(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}

	if col+3 >= len(lines[0]) || line+3 >= len(lines) {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line+1][col+1])
	strbuilder.WriteByte(lines[line+2][col+2])
	strbuilder.WriteByte(lines[line+3][col+3])

	return strbuilder.String()
}
func diagonal_leftdown(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}

	if col-3 >= len(lines[0]) || line+3 >= len(lines) || col-3 < 0 {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line+1][col-1])
	strbuilder.WriteByte(lines[line+2][col-2])
	strbuilder.WriteByte(lines[line+3][col-3])

	return strbuilder.String()
}

func diagonal_rightdown2(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}

	if col+2 >= len(lines[0]) || line+2 >= len(lines) {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line+1][col+1])
	strbuilder.WriteByte(lines[line+2][col+2])

	return strbuilder.String()
}
func diagonal_leftdown2(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}

	if col-2 >= len(lines[0]) || line+2 >= len(lines) || col-2 < 0 {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line+1][col-1])
	strbuilder.WriteByte(lines[line+2][col-2])

	return strbuilder.String()
}

func vertical(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}
	if line+3 >= len(lines) {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line+1][col])
	strbuilder.WriteByte(lines[line+2][col])
	strbuilder.WriteByte(lines[line+3][col])

	return strbuilder.String()
}

func horizontal(col int, line int, lines []string) string {
	strbuilder := strings.Builder{}

	if col+3 >= len(lines[line]) {
		return ""
	}

	strbuilder.WriteByte(lines[line][col])
	strbuilder.WriteByte(lines[line][col+1])
	strbuilder.WriteByte(lines[line][col+2])
	strbuilder.WriteByte(lines[line][col+3])

	return strbuilder.String()
}

func puzzle1(lines []string) int {
	count := 0
	columns := len(lines[0])

	for y := 0; y < len(lines); y++ {
		for x := 0; x < columns; x++ {
			c := lines[y][x]
			if c == 'X' {
				count += utils.BoolToInt(horizontal(x, y, lines) == "XMAS")
				count += utils.BoolToInt(vertical(x, y, lines) == "XMAS")
				count += utils.BoolToInt(diagonal_rightdown(x, y, lines) == "XMAS")
				count += utils.BoolToInt(diagonal_leftdown(x, y, lines) == "XMAS")
			}
			if c == 'S' {
				count += utils.BoolToInt(horizontal(x, y, lines) == "SAMX")
				count += utils.BoolToInt(vertical(x, y, lines) == "SAMX")
				count += utils.BoolToInt(diagonal_rightdown(x, y, lines) == "SAMX")
				count += utils.BoolToInt(diagonal_leftdown(x, y, lines) == "SAMX")
			}
		}
	}
	return count
}

func puzzle2(lines []string) int {
	count := 0
	columns := len(lines[0])

	for y := 0; y < len(lines); y++ {
		for x := 0; x < columns; x++ {
			line := lines[y]
			if x+2 < len(line) {
				if line[x] == 'M' || line[x] == 'S' {
					fmt.Println(x+1, y+1, diagonal_rightdown2(x, y, lines), diagonal_leftdown2(x+2, y, lines))
					if (diagonal_rightdown2(x, y, lines) == "MAS" || diagonal_rightdown2(x, y, lines) == "SAM") && (diagonal_leftdown2(x+2, y, lines) == "MAS" || diagonal_leftdown2(x+2, y, lines) == "SAM") {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	str, error := utils.ReadFileStr("./data/day4.txt")
	if error == nil {
		lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

		columns := len(lines[0])

		for y := 0; y < len(lines); y++ { //remove wrong lines
			if len(lines[y]) != columns {
				lines = append(lines[:y], lines[y+1:]...)
			}
		}

		puzzle1_res := puzzle1(lines)
		puzzle2_res := puzzle2(lines)

		fmt.Printf("Day 4 - Puzzle 1 answer is %v\n", puzzle1_res)
		fmt.Printf("Day 4 - Puzzle 2 answer is %v\n", puzzle2_res)
	}
}
