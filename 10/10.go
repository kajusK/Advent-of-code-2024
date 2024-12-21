package main

import (
	"os"
	"bufio"
	"fmt"
)


func get_tops(terrain [][]int, x int, y int, tops map[[2]int]int) {
	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, -1},
		{0, 1},
	}

	cur := terrain[x][y]
	if cur == 9 {
		tops[[2]int{x, y}] += 1
		return
	}
	for _, dir := range dirs {
		new_x := x + dir[0]
		new_y := y + dir[1]
		if new_x >= len(terrain) || new_x < 0 || new_y >= len(terrain[0]) || new_y < 0 {
			continue
		}
		if terrain[new_x][new_y] == cur + 1 {
			get_tops(terrain, new_x, new_y, tops)
		}
	}
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)

	var terrain [][]int

	for scanner.Scan() {
		var line []int 
		for _, c := range scanner.Text() {
			line = append(line, int(c - '0'))
		}
		terrain = append(terrain, line)
	}

	part1 := 0
	part2 := 0
	for x, line := range terrain {
		for y, height := range line {
			if height == 0 {
				tops := make(map[[2]int]int)
				get_tops(terrain, x, y, tops)
				part1 += len(tops)
				for _, v := range tops {
					part2 += v
				}
			}
		}
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}