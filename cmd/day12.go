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

func get_area_perimeter(x int, y int, plant byte, lines []string, visited_positions map[Pos]int) (int, int) {
	if _, found := visited_positions[Pos{x, y}]; found {
		return 0, 0
	}

	check_pos := func(x int, y int) byte {
		if y >= len(lines) || y < 0 {
			return 0
		}
		if x >= len(lines[y]) || x < 0 {
			return 0
		}
		return lines[y][x]
	}

	visited_positions[Pos{x, y}]++
	perimeter := 4
	area := 1
	if check_pos(x-1, y) == plant {
		perimeter--
	}
	if check_pos(x, y-1) == plant {
		perimeter--
	}
	if check_pos(x-1, y) == plant {
		perimeter--
	}
	if check_pos(x, y-1) == plant {
		perimeter--
	}

	search := func(x int, y int) {
		if check_pos(x, y) == plant {
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
			fmt.Printf("Perimeter for %c(%v, %v) is %v and area is %v\n", plant, x, y, perimeter, area)
		}
	}
	return price
}
func main() {
	str, error := utils.ReadFileStr("./data/day12.txt")
	if error == nil {
		fmt.Printf("Day 12 - Puzzle 1 answer is %v\n", puzzle1(str))
	}
}
