package main

import (
	"fmt"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func puzzle1(board []string, movements string) int {
	x := -1
	y := -1

	can_move := func(movement rune) bool {
		if movement == '<' || movement == '>' {
			line := board[y]
			new_x := x
			if movement == '<' {
				new_x--
			} else {
				new_x++
			}
			c := line[new_x]
			if new_x < 0 || new_x >= len(line) {
				return false
			}
			if c == '#' {
				return false
			} else if c == '.' {
				return true
			} else if c == 'O' {
				for new_x >= 0 && new_x < len(line) {
					c := line[new_x]
					if c == '#' {
						return false
					}
					if c == '.' {
						return true
					}
					if movement == '<' {
						new_x--
					} else {
						new_x++
					}
				}
			}
		} else if movement == '^' || movement == 'v' {
			new_y := y
			if movement == '^' {
				new_y--
			} else {
				new_y++
			}
			if new_y < 0 || new_y >= len(board) {
				return false
			}
			c := board[new_y][x]
			if c == '#' {
				return false
			} else if c == '.' {
				return true
			} else if c == 'O' {
				for new_y >= 0 && new_y < len(board) {
					c := board[new_y][x]
					if c == '#' {
						return false
					}
					if c == '.' {
						return true
					}
					if movement == '^' {
						new_y--
					} else {
						new_y++
					}
				}
			}
		}
		return false
	}
	move := func(movement rune) {
		if movement == '<' {
			new_x := x - 1
			if board[y][new_x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y] = utils.StringReplaceAtIndex(board[y], '@', new_x)
				x--
			} else if board[y][new_x] == 'O' {
				for board[y][new_x] != '.' {
					new_x--
				}
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y] = utils.StringReplaceAtIndex(board[y], 'O', new_x)
				x--
				board[y] = utils.StringReplaceAtIndex(board[y], '@', x)
			}
		} else if movement == '>' {
			new_x := x + 1
			if board[y][new_x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y] = utils.StringReplaceAtIndex(board[y], '@', new_x)
				x++
			} else if board[y][new_x] == 'O' {
				for board[y][new_x] != '.' {
					new_x++
				}
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y] = utils.StringReplaceAtIndex(board[y], 'O', new_x)
				x++
				board[y] = utils.StringReplaceAtIndex(board[y], '@', x)
			}
		} else if movement == '^' {
			new_y := y - 1
			if board[new_y][x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], '@', x)
				y--
			} else if board[new_y][x] == 'O' {
				for board[new_y][x] != '.' {
					new_y--
				}
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], 'O', x)
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				y--
				board[y] = utils.StringReplaceAtIndex(board[y], '@', x)
			}
		} else if movement == 'v' {
			new_y := y + 1
			if board[new_y][x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], '@', x)
				y++
			} else if board[new_y][x] == 'O' {
				for board[new_y][x] != '.' {
					new_y++
				}
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], 'O', x)
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				y++
				board[y] = utils.StringReplaceAtIndex(board[y], '@', x)
			}
		}
	}

	for _y, line := range board {
		for _x, c := range line {
			if c == '@' {
				x = _x
				y = _y
				break
			}
		}
		if x != -1 && y != -1 {
			break
		}
	}
	for _, movement := range movements {
		if can_move(movement) {
			move(movement)
		}
	}
	sum := 0
	for y, line := range board {
		for x, c := range line {
			if c == 'O' {
				sum += y*100 + x
			}
		}
	}
	return sum
}

func main() {
	str, error := utils.ReadFileStr("./data/day15.txt")
	if error == nil {
		sep := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n\n")
		if len(sep) == 2 {
			board := strings.Split(sep[0], "\n")
			movements := strings.ReplaceAll(sep[1], "\n", "")
			fmt.Printf("Day 15 - Puzzle 1 answer is %v\n", puzzle1(board, movements))
		}
	}
}
