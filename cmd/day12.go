package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Pos struct {
	x int
	y int
}

func check_pos(x int, y int, lines []string) byte {
	if y >= len(lines) || y < 0 {
		return 0
	}
	if x >= len(lines[y]) || x < 0 {
		return 0
	}
	return lines[y][x]
}

func get_area_perimeter(x int, y int, plant byte, lines []string, visited_positions map[Pos]int) (int, int) {
	if _, found := visited_positions[Pos{x, y}]; found {
		return 0, 0
	}

	visited_positions[Pos{x, y}]++
	perimeter := 4
	area := 1
	if check_pos(x-1, y, lines) == plant {
		perimeter--
	}
	if check_pos(x, y-1, lines) == plant {
		perimeter--
	}
	if check_pos(x-1, y, lines) == plant {
		perimeter--
	}
	if check_pos(x, y-1, lines) == plant {
		perimeter--
	}

	search := func(x int, y int) {
		if check_pos(x, y, lines) == plant {
			p, a := get_area_perimeter(x, y, plant, lines, visited_positions)
			perimeter += p
			area += a
		}
	}

	search(x+1, y)
	search(x-1, y)
	search(x, y+1)
	search(x, y-1)

	return perimeter, area
}
func get_normals(x int, y int, plant byte, lines []string, normal_map map[Pos][]Pos, visited_positions map[Pos]int) {
	if _, found := visited_positions[Pos{x, y}]; found {
		return
	}
	visited_positions[Pos{x, y}]++
	if check_pos(x, y-1, lines) != plant {
		normal_map[Pos{0, -1}] = append(normal_map[Pos{0, -1}], Pos{x, y})
	}
	if check_pos(x, y+1, lines) != plant {
		normal_map[Pos{0, 1}] = append(normal_map[Pos{0, 1}], Pos{x, y})
	}
	if check_pos(x-1, y, lines) != plant {
		normal_map[Pos{-1, 0}] = append(normal_map[Pos{-1, 0}], Pos{x, y})
	}
	if check_pos(x+1, y, lines) != plant {
		normal_map[Pos{1, 0}] = append(normal_map[Pos{1, 0}], Pos{x, y})
	}

	search := func(x int, y int) {
		if check_pos(x, y, lines) == plant {
			get_normals(x, y, plant, lines, normal_map, visited_positions)
		}
	}

	search(x+1, y)
	search(x-1, y)
	search(x, y+1)
	search(x, y-1)
}
func get_sides(x int, y int, plant byte, lines []string, visited_positions map[Pos]int) int {
	normal_map := map[Pos][]Pos{}
	get_normals(x, y, plant, lines, normal_map, visited_positions)

	edges := [][]Pos{}
	for _, positions := range normal_map {
		sort.Slice(positions, func(i, j int) bool {
			if positions[i].x == positions[j].x {
				return positions[i].y < positions[j].y
			}
			return positions[i].x < positions[j].x
		})
		used_map := map[Pos]int{}
		for _, pos := range positions {
			if _, found := used_map[pos]; found {
				continue
			}
			edge := []Pos{pos}
			used_map[pos]++
			x_pos := 1
			y_pos := 1
			for _, pos2 := range positions {
				if _, found := used_map[pos2]; found {
					continue
				}
				if pos.x == pos2.x && pos.y+y_pos == pos2.y {
					edge = append(edge, pos2)
					used_map[pos2]++
					y_pos++
				} else if pos.y == pos2.y && pos.x+x_pos == pos2.x {
					edge = append(edge, pos2)
					used_map[pos2]++
					x_pos++
				}
			}
			edges = append(edges, edge)
		}
	}
	return len(edges)
}
func puzzle1(data string) int {
	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	visited_positions := map[Pos]int{}
	price := 0
	for y, line := range lines {
		for x, plant := range line {
			if _, found := visited_positions[Pos{x, y}]; found {
				continue
			}
			perimeter, area := get_area_perimeter(x, y, byte(plant), lines, visited_positions)
			price += perimeter * area
		}
	}
	return price
}
func puzzle2(data string) int {
	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")
	visited_positions_area := map[Pos]int{}
	visited_positions_sides := map[Pos]int{}
	price := 0
	for y, line := range lines {
		for x, plant := range line {
			_, area := get_area_perimeter(x, y, byte(plant), lines, visited_positions_area)
			sides := get_sides(x, y, byte(plant), lines, visited_positions_sides)
			if sides != 0 {
				price += area * sides
			}
		}
	}
	return price
}
func main() {
	str, error := utils.ReadFileStr("./data/day12.txt")
	if error == nil {
		fmt.Printf("Day 12 - Puzzle 1 answer is %v\n", puzzle1(str))
		fmt.Printf("Day 12 - Puzzle 2 answer is %v\n", puzzle2(str))
	}
}
