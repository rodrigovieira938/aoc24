package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Vec2 struct {
	x, y int
}

type Robot struct {
	pos      Vec2
	velocity Vec2
}

func parse_data(str string) []Robot {
	robots := []Robot{}
	lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

	for _, line := range lines {
		sep := strings.Split(line, "p=")
		if len(sep) == 2 {
			sep = strings.Split(sep[1], " v=")
			if len(sep) == 2 {
				pos := strings.Split(sep[0], ",")
				if len(pos) != 2 {
					continue
				}
				vel := strings.Split(sep[1], ",")
				if len(vel) != 2 {
					continue
				}
				robot := Robot{}
				var error error
				robot.pos.x, error = strconv.Atoi(pos[0])
				if error != nil {
					continue
				}
				robot.pos.y, error = strconv.Atoi(pos[1])
				if error != nil {
					continue
				}
				robot.velocity.x, error = strconv.Atoi(vel[0])
				if error != nil {
					continue
				}
				robot.velocity.y, error = strconv.Atoi(vel[1])
				if error != nil {
					continue
				}
				robots = append(robots, robot)
			}
		}
	}

	return robots
}
func puzzle1(robots []Robot) int {
	quadrants := [4]int{0, 0, 0, 0}
	map_size_x := 101
	map_size_y := 103
	lines := make([][]byte, map_size_y)
	for y := range lines {
		for x := 0; x < map_size_x; x++ {
			lines[y] = append(lines[y], '.')
		}
	}

	vertical_split := (map_size_x - 1) / 2
	horizontal_split := (map_size_y - 1) / 2

	get_pos_after_seconds := func(x int, y int, vx int, vy int, seconds int) (int, int) {
		x_pos := (x + vx*seconds) % map_size_x
		if x_pos < 0 {
			x_pos = map_size_x + x_pos
		}
		y_pos := (y + vy*seconds) % map_size_y
		if y_pos < 0 {
			y_pos = map_size_y + y_pos
		}
		return x_pos, y_pos
	}

	for _, robot := range robots {
		x, y := get_pos_after_seconds(robot.pos.x, robot.pos.y, robot.velocity.x, robot.velocity.y, 200)
		if y == horizontal_split || x == vertical_split {
			continue
		}

		if y > horizontal_split {
			if x > vertical_split {
				quadrants[3]++
			} else {
				quadrants[2]++
			}
		} else {
			if x > vertical_split {
				quadrants[1]++
			} else {
				quadrants[0]++
			}
		}

		if lines[y][x] != '.' {
			lines[y][x] += 1
		} else {
			lines[y][x] = '1'
		}
	}
	for y, line := range lines {
		for x, char := range line {
			if y == horizontal_split || x == vertical_split {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func puzzle2(robots []Robot) int {
	map_size_x := 101
	map_size_y := 103

	vertical_split := (map_size_x - 1) / 2
	horizontal_split := (map_size_y - 1) / 2

	get_pos_after_seconds := func(x int, y int, vx int, vy int, seconds int) (int, int) {
		x_pos := (x + vx*seconds) % map_size_x
		if x_pos < 0 {
			x_pos = map_size_x + x_pos
		}
		y_pos := (y + vy*seconds) % map_size_y
		if y_pos < 0 {
			y_pos = map_size_y + y_pos
		}
		return x_pos, y_pos
	}

	seconds := 0
	for {
		lines := make([][]byte, map_size_y)

		for y := range lines {
			for x := 0; x < map_size_x; x++ {
				lines[y] = append(lines[y], '.')
			}
		}
		for _, robot := range robots {
			x, y := get_pos_after_seconds(robot.pos.x, robot.pos.y, robot.velocity.x, robot.velocity.y, seconds)
			if y == horizontal_split || x == vertical_split {
				continue
			}

			if lines[y][x] != '.' {
				lines[y][x] += 1
			} else {
				lines[y][x] = '1'
			}
		}
		builder := strings.Builder{}
		for y, line := range lines {
			for x, char := range line {
				if y == horizontal_split || x == vertical_split {
					builder.WriteRune(' ')
				} else {
					builder.WriteByte(char)
				}
			}
			builder.WriteRune('\n')
		}
		str := builder.String()
		//Find for christmas tree
		if strings.Contains(str, "1111111111111111111111111111111") {
			fmt.Println(str)
			return seconds
		}

		seconds++
	}

	return -1
}

func main() {
	str, error := utils.ReadFileStr("./data/day14.txt")
	if error == nil {
		data := parse_data(str)
		fmt.Printf("Day 14 - Puzzle 1 answer is %v\n", puzzle1(data))
		fmt.Printf("Day 14 - Puzzle 2 answer is %v\n", puzzle2(data))
	}
}
