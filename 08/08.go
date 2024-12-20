package main

import (
	"os"
	"bufio"
	"fmt"
)

func is_in_map(pos [2]int, size [2]int) bool {
	return pos[0] >= 0 && pos[0] < size[0] && pos[1] >= 0 && pos[1] < size[1]
}

func get_pair_antinode(a [2]int, b [2]int, size [2]int) [][2]int {
	distance := [2]int{a[0] - b[0], a[1] - b[1]}
	var res [][2]int

	res1 := [2]int{a[0] + distance[0], a[1] + distance[1]}
	if is_in_map(res1, size) {
		res = append(res, res1)
	}

	res2 := [2]int{b[0] - distance[0], b[1] - distance[1]}
	if is_in_map(res2, size) {
		res = append(res, res2)
	}
	return res
}

func get_all_antinode(a [2]int, b [2]int, size [2]int) [][2]int {
	distance := [2]int{a[0] - b[0], a[1] - b[1]}
	var res [][2]int

	res = append(res, a)
	pos := a
	for {
		pos = [2]int{pos[0] + distance[0], pos[1] + distance[1]}
		if !is_in_map(pos, size) {
			break
		}
		res = append(res, pos)
	}

	pos = a
	for {
		pos = [2]int{pos[0] - distance[0], pos[1] - distance[1]}
		if !is_in_map(pos, size) {
			break
		}
		res = append(res, pos)
	}
	return res
}

func get_antinotes(antennas map[rune][][2]int, size [2]int, get_nodes func([2]int, [2]int, [2]int)[][2]int) int {
	antinodes := make(map[[2]int]bool)

	for _, positions := range antennas {
		for i, a := range positions {
			for _, b := range positions[i+1:] {
				for _, pos := range get_nodes(a, b, size) {
					antinodes[pos] = true
				}
			}
		}
	}

	return len(antinodes)
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    antennas := make(map[rune][][2]int)

    scanner := bufio.NewScanner(file)
	y := 0
	width := 0
    for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, c := range line {
			if c == '.' {
				continue
			} 
			antennas[c] = append(antennas[c], [2]int{x, y})
		}
		y++
    }
	size := [2]int{width, y}

	fmt.Println("Part 1: ", get_antinotes(antennas, size, get_pair_antinode))
	fmt.Println("Part 2: ", get_antinotes(antennas, size, get_all_antinode))
}