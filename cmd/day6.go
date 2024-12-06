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
type TurningPoint struct {
	char byte
	x    int
	y    int
}

func find_starting_pos(data []string) Pos {
	for y, line := range data {
		for x, char := range line {
			if char == '^' || char == '<' || char == '>' || char == 'V' {
				return Pos{x, y}
			}
		}
	}
	return Pos{-1, -1}
}
func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
func print_map(data []string) {
	//It takes a lot of time to print out
	/*for _, line := range data {
		fmt.Println(line)
	}
	fmt.Println()*/
}

func puzzle1(data []string) (int, []TurningPoint) {
	starting_pos := find_starting_pos(data)
	pos := starting_pos
	var c byte = ' '

	lines := len(data)
	cols := len(data[0])
	print_map(data)

	distinct_pos := map[Pos]bool{}
	distinct_pos[pos] = true

	turning_points := []TurningPoint{}

	for true {
		c = data[pos.y][pos.x]

		_, found := distinct_pos[pos]
		if !found {
			distinct_pos[pos] = true
		}

		if c == '^' {
			if pos.y-1 < 0 {
				break
			} else {
				if data[pos.y-1][pos.x] == '#' {
					data[pos.y] = replaceAtIndex(data[pos.y], '>', pos.x)
					turning_points = append(turning_points, TurningPoint{'~', pos.x, pos.y})
					print_map(data)
				} else {
					data[pos.y] = replaceAtIndex(data[pos.y], 'X', pos.x)
					pos.y -= 1
					data[pos.y] = replaceAtIndex(data[pos.y], '^', pos.x)
					print_map(data)
				}
			}
		} else if c == 'V' {
			if pos.y+1 >= lines {
				break
			} else {
				if data[pos.y+1][pos.x] == '#' {
					data[pos.y] = replaceAtIndex(data[pos.y], '<', pos.x)
					turning_points = append(turning_points, TurningPoint{'V', pos.x, pos.y})
					print_map(data)
				} else {
					data[pos.y] = replaceAtIndex(data[pos.y], 'X', pos.x)
					pos.y += 1
					data[pos.y] = replaceAtIndex(data[pos.y], 'V', pos.x)
					print_map(data)
				}
			}
		} else if c == '<' {
			if pos.x-1 < 0 {
				break
			} else {
				if data[pos.y][pos.x-1] == '#' {
					data[pos.y] = replaceAtIndex(data[pos.y], '^', pos.x)
					turning_points = append(turning_points, TurningPoint{'<', pos.x, pos.y})
					print_map(data)
				} else {
					data[pos.y] = replaceAtIndex(data[pos.y], 'X', pos.x)
					pos.x -= 1
					data[pos.y] = replaceAtIndex(data[pos.y], '<', pos.x)
					print_map(data)
				}
			}
		} else if c == '>' {
			if pos.x+1 >= cols {
				break
			} else {
				if data[pos.y][pos.x+1] == '#' {
					data[pos.y] = replaceAtIndex(data[pos.y], 'V', pos.x)
					turning_points = append(turning_points, TurningPoint{'>', pos.x, pos.y})
					print_map(data)
				} else {
					data[pos.y] = replaceAtIndex(data[pos.y], 'X', pos.x)
					pos.x += 1
					data[pos.y] = replaceAtIndex(data[pos.y], '>', pos.x)
					print_map(data)
				}
			}
		}
	}
	distinct_pos_sum := 0

	for range distinct_pos {
		distinct_pos_sum++
	}
	return distinct_pos_sum, turning_points
}

func main() {
	str, error := utils.ReadFileStr("./data/day6.txt")
	if error == nil {
		data := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
		distinct_pos_sum, turning_points := puzzle1(data)
		fmt.Printf("Day 5 - Puzzle 1 answer is %v\n", distinct_pos_sum)
	}
}
