package main

import (
	"fmt"
	"slices"
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

	for {
		c = data[pos.y][pos.x]

		_, found := distinct_pos[pos]
		if !found {
			distinct_pos[pos] = true
		}

		if c == '^' {
			if pos.y-1 < 0 {
				break
			} else {
				if data[pos.y-1][pos.x] == '#' || data[pos.y-1][pos.x] == 'O' {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '>', pos.x)
					turning_point := TurningPoint{'^', pos.x, pos.y}
					for _, point := range turning_points {
						if point == turning_point {
							return 0, []TurningPoint{}
						}
					}
					turning_points = append(turning_points, turning_point)
					print_map(data)
				} else {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'X', pos.x)
					pos.y -= 1
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '^', pos.x)
					print_map(data)
				}
			}
		} else if c == 'V' {
			if pos.y+1 >= lines {
				break
			} else {
				if data[pos.y+1][pos.x] == '#' || data[pos.y+1][pos.x] == 'O' {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '<', pos.x)
					turning_point := TurningPoint{'V', pos.x, pos.y}
					for _, point := range turning_points {
						if point == turning_point {
							return 0, []TurningPoint{}
						}
					}
					turning_points = append(turning_points, turning_point)
					print_map(data)
				} else {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'X', pos.x)
					pos.y += 1
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'V', pos.x)
					print_map(data)
				}
			}
		} else if c == '<' {
			if pos.x-1 < 0 {
				break
			} else {
				if data[pos.y][pos.x-1] == '#' || data[pos.y][pos.x-1] == 'O' {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '^', pos.x)
					turning_point := TurningPoint{'<', pos.x, pos.y}
					for _, point := range turning_points {
						if point == turning_point {
							return 0, []TurningPoint{}
						}
					}
					turning_points = append(turning_points, turning_point)
					print_map(data)
				} else {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'X', pos.x)
					pos.x -= 1
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '<', pos.x)
					print_map(data)
				}
			}
		} else if c == '>' {
			if pos.x+1 >= cols {
				break
			} else {
				if data[pos.y][pos.x+1] == '#' || data[pos.y][pos.x+1] == 'O' {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'V', pos.x)
					turning_point := TurningPoint{'>', pos.x, pos.y}
					for _, point := range turning_points {
						if point == turning_point {
							return 0, []TurningPoint{}
						}
					}
					turning_points = append(turning_points, turning_point)
					print_map(data)
				} else {
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], 'X', pos.x)
					pos.x += 1
					data[pos.y] = utils.StringReplaceAtIndex(data[pos.y], '>', pos.x)
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

func puzzle2(data []string) int {
	count := 0
	for y, line := range data {
		if len(line) == 0 {
			continue
		}
		for x := range line {
			if line[x] != '#' && line[x] != 'O' && line[x] != '<' && line[x] != '>' && line[x] != '^' && line[x] != 'V' {
				clone := make([]string, len(data))
				copy(clone, data)
				clone[y] = utils.StringReplaceAtIndex(clone[y], 'O', x)
				distinct_pos_sum, _ := puzzle1(clone)
				is_loop := distinct_pos_sum == 0
				if is_loop {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	str, error := utils.ReadFileStr("./data/day6.txt")
	if error == nil {
		data := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")
		cols := len(data[0])

		// Remove empty lines, or lines with less characters than the first line
		for i, line := range data {
			if len(line) != cols {
				data = append(data[:i], data[i+1:]...)
			}
		}
		distinct_pos_sum, _ := puzzle1(slices.Clone(data))
		fmt.Printf("Day 6 - Puzzle 1 answer is %v\n", distinct_pos_sum)
		fmt.Printf("Day 6 - Puzzle 2 answer is %v\n", puzzle2(data))
	}
}
