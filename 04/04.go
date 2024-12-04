package main

import (
	"os"
	"bufio"
	"fmt"
)

func getChar(input []string, x int, y int) rune {
	if x >= len(input) || x < 0 {
		return '0'		// to lazy to implement error checking and stuff
	}
	if y >= len(input[x]) || y < 0 {
		return '0'
	}
	return rune(input[x][y])
}


func searchXmas(input []string, x int, y int, dir []int) bool {
	search := "XMAS"

	for _, char := range search {
		if char != getChar(input, x, y) {
			return false
		}
		x += dir[0]
		y += dir[1]
	} 
	return true
}

func searchAllDirs(input []string, x int, y int) int {
	vals := []int{-1, 0, 1}
	res := 0

	for _, dir_x := range vals {
		for _, dir_y := range vals {
			dir := []int{dir_x, dir_y}
			if searchXmas(input, x, y, dir) {
				res++
			}
		}
	}
	return res
}

func isX(input []string, x int, y int) bool {
	diagonals := [][][]int{
		{{-1, -1}, {0, 0}, {1, 1}},
		{{-1, 1}, {0, 0}, {1, -1}},
	}

	for _, diagonal := range diagonals {
		var res string
		for _, dir := range diagonal {
			res = res + string(getChar(input, x+dir[0], y+dir[1]))
		}
		if res != "MAS" && res != "SAM" {
			return false
		}
	}
	return true
}

func part1(input []string) int {
	res := 0

	for x, line := range input {
		for y, char := range line {
			if char != 'X' {
				continue
			}
			res += searchAllDirs(input, x, y)
		}
	}
	return res
}

func part2(input []string) int {
	res := 0

	for x, line := range input {
		for y, char := range line {
			if char == 'A' && isX(input, x, y) {
				res++
			}
		}
	}
	return res
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    var input []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		input = append(input, scanner.Text())
    }

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}