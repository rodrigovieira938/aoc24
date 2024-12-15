package main

import (
	"fmt"
	"slices"
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
func puzzle2(board []string, movements string) int {
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
			} else if c == '[' || c == ']' {
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
			d := 0
			if movement == '^' {
				d--
			} else {
				d++
			}
			var check func(x, y int) bool
			check = func(x, y int) bool {
				x_open := 0
				x_close := 0
				if board[y][x] == '#' {
					return false
				} else if board[y][x] == '.' {
					return true
				}
				if board[y][x] == '[' {
					x_open = x
					x_close = x + 1
				} else {
					x_open = x - 1
					x_close = x
				}
				if board[y+d][x_open] != '.' {
					if !check(x_open, y+d) {
						return false
					}
				}
				if board[y+d][x_close] != '.' {
					if !check(x_close, y+d) {
						return false
					}
				}
				return true
			}
			return check(x, y+d)
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
			} else if board[y][new_x] == ']' {
				for board[y][new_x] != '.' {
					new_x--
				}
				builder := strings.Builder{}
				builder.WriteString(board[y][0:new_x])
				builder.WriteString(board[y][new_x+1 : x+1])
				builder.WriteRune('.')
				builder.WriteString(board[y][x+1:])
				board[y] = builder.String()
				x--
			}
		} else if movement == '>' {
			new_x := x + 1
			if board[y][new_x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y] = utils.StringReplaceAtIndex(board[y], '@', new_x)
				x++
			} else if board[y][new_x] == '[' {
				for board[y][new_x] != '.' {
					new_x++
				}
				builder := strings.Builder{}
				builder.WriteString(board[y][0:x])
				builder.WriteRune('.')
				builder.WriteString(board[y][x:new_x])
				builder.WriteString(board[y][new_x+1:])
				board[y] = builder.String()
				x++
			}
		} else if movement == '^' {
			new_y := y - 1
			if board[new_y][x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], '@', x)
				y--
			} else if board[new_y][x] == '[' || board[new_y][x] == ']' {
				var move func(x, y int)
				move = func(x, y int) {
					x_open := 0
					x_close := 0
					if board[y][x] == '[' {
						x_open = x
						x_close = x + 1
					} else {
						x_open = x - 1
						x_close = x
					}
					if board[y-1][x_open] != '.' {
						move(x_open, y-1)
					}
					if board[y-1][x_close] != '.' {
						move(x_close, y-1)
					}

					board[y] = utils.StringReplaceAtIndex(board[y], '.', x_open)
					board[y] = utils.StringReplaceAtIndex(board[y], '.', x_close)
					board[y-1] = utils.StringReplaceAtIndex(board[y-1], '[', x_open)
					board[y-1] = utils.StringReplaceAtIndex(board[y-1], ']', x_close)
				}
				move(x, new_y)
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y-1] = utils.StringReplaceAtIndex(board[y-1], '@', x)
				y--
			}
		} else if movement == 'v' {
			new_y := y + 1
			if board[new_y][x] == '.' {
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[new_y] = utils.StringReplaceAtIndex(board[new_y], '@', x)
				y++
			} else if board[new_y][x] == '[' || board[new_y][x] == ']' {
				var move func(x, y int)
				move = func(x, y int) {
					x_open := 0
					x_close := 0
					if board[y][x] == '[' {
						x_open = x
						x_close = x + 1
					} else {
						x_open = x - 1
						x_close = x
					}

					if board[y+1][x_open] != '.' {
						move(x_open, y+1)
					}
					if board[y+1][x_close] != '.' {
						move(x_close, y+1)
					}

					board[y] = utils.StringReplaceAtIndex(board[y], '.', x_open)
					board[y] = utils.StringReplaceAtIndex(board[y], '.', x_close)
					board[y+1] = utils.StringReplaceAtIndex(board[y+1], '[', x_open)
					board[y+1] = utils.StringReplaceAtIndex(board[y+1], ']', x_close)
				}
				move(x, new_y)
				board[y] = utils.StringReplaceAtIndex(board[y], '.', x)
				board[y+1] = utils.StringReplaceAtIndex(board[y+1], '@', x)
				y++
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
			if c == '[' {
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
			fmt.Printf("Day 15 - Puzzle 1 answer is %v\n", puzzle1(slices.Clone(board), movements))
			builder := strings.Builder{}
			for _, line := range board {
				for _, c := range line {
					if c == '@' {
						builder.WriteString("@.")
					} else if c == 'O' {
						builder.WriteString("[]")
					} else if c == '#' {
						builder.WriteString("##")
					} else {
						builder.WriteString("..")
					}
				}
				builder.WriteRune('\n')
			}
			newboard := strings.Split(builder.String(), "\n")
			fmt.Printf("Day 15 - Puzzle 1 answer is %v\n", puzzle2(newboard, movements))
		}
	}
}
