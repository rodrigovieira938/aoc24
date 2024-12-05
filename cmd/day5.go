package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RodrigoVieira938/aoc24/utils"
)

type Dependency struct {
	first  int
	second int
}

type Update []int

type Data struct {
	Dependencies []Dependency
	Updates      []Update
}

func parse_data(str string) Data {
	lines := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n")

	data := Data{}

	for _, line := range lines {
		update_dependencies := strings.Split(line, "|")
		if len(update_dependencies) > 1 {
			num1, error1 := strconv.Atoi(update_dependencies[0])
			num2, error2 := strconv.Atoi(update_dependencies[1])
			if error1 == nil && error2 == nil {
				dep := Dependency{num1, num2}
				data.Dependencies = append(data.Dependencies, dep)
			}
		} else {
			update_str := strings.Split(line, ",")

			update := Update{}
			all_good := true

			for _, number_str := range update_str {
				n, error := strconv.Atoi(number_str)
				if error != nil {
					all_good = false
					break
				} else {
					update = append(update, n)
				}
			}
			if len(update) > 0 && all_good {
				data.Updates = append(data.Updates, update)
			}
		}
	}
	return data
}

// Searchs in map or insert, if not found in map nor dependencies returns -1
func search_in_map_or_insert_second(dependency int, dep_map map[int][]int, dependencies []Dependency) []int {
	val, ok := dep_map[dependency]
	if ok {
		return val
	}
	for _, dep := range dependencies {
		if dep.second == dependency {
			dep_map[dependency] = append(dep_map[dependency], dep.first)
		}
	}
	return dep_map[dependency]
}

func search_in_map_or_insert_first(dependency int, dep_map map[int][]int, dependencies []Dependency) []int {
	val, ok := dep_map[dependency]
	if ok {
		return val
	}
	for _, dep := range dependencies {
		if dep.first == dependency {
			dep_map[dependency] = append(dep_map[dependency], dep.second)
		}
	}
	return dep_map[dependency]
}

func is_update_correct(update Update, dependencies []Dependency, dep_map map[int][]int) bool {
	ok := true

	for i := len(update) - 1; i >= 0; i-- {
		num := update[i]
		dependencies := search_in_map_or_insert_second(num, dep_map, dependencies)

		utils.SliceContains(dependencies, 0)

		for _, dep := range dependencies {
			index, found := utils.SliceContains(update, dep)
			if found && index > i {
				ok = false
				break
			}
		}
		if !ok {
			break
		}
	}
	return ok
}

func puzzle1(data Data) (int, []Update) {
	sum := 0

	dep_map := make(map[int][]int)
	incorrectly_updates := []Update{}

	for _, update := range data.Updates {

		if is_update_correct(update, data.Dependencies, dep_map) {
			sum += update[len(update)/2]
		} else {
			incorrectly_updates = append(incorrectly_updates, update)
		}
	}
	return sum, incorrectly_updates
}

func puzzle2(data Data) int {
	sum := 0

	dep_map := make(map[int][]int)

	for _, update := range data.Updates {
		for !is_update_correct(update, data.Dependencies, dep_map) {
			for i := len(update) - 1; i >= 0; i-- {
				num := update[i]
				dependencies := search_in_map_or_insert_second(num, dep_map, data.Dependencies)
				move := 0
				for _, val := range dependencies {
					index, found := utils.SliceContains(update, val)
					if index > move && found {
						move = index
					}
				}
				n := update[i]
				update[i] = update[move]
				update[move] = n
			}
		}
		sum += update[len(update)/2]
	}
	return sum
}

func main() {
	str, error := utils.ReadFileStr("./data/day5.txt")
	if error == nil {
		data := parse_data(str)
		sum, incorrectly_updates := puzzle1(data)
		fmt.Printf("Day 5 - Puzzle 1 answer is %v\n", sum)
		data.Updates = incorrectly_updates
		fmt.Printf("Day 5 - Puzzle 2 answer is %v\n", puzzle2(data))
	}
}
