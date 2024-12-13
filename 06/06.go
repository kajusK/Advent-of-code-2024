package main

import (
	"os"
	"bufio"
	"fmt"
)

func getDirection(char rune) [2]int {
	dirs := map[rune][2]int {
		'>': [2]int{1, 0},
		'<': [2]int{-1, 0},
		'^': [2]int{0, -1},
		'v': [2]int{0, 1},
	}
	return dirs[char]
}

func move(pos [2]int, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func isNextObstacle(obstacles map[[2]int]bool, pos [2]int, dir [2]int) bool {
	return obstacles[move(pos, dir)]
}

func turnRight(dir [2]int) [2]int {
	return [2]int{-dir[1], dir[0]}
}

func isInArea(pos [2]int, size [2]int) bool {
	if pos[0] < 0 || pos[0] >= size[0] || pos[1] < 0 || pos[1] >= size[1] {
		return false
	}
	return true
}

func part1(obstacles map[[2]int]bool, pos [2]int, dir [2]int, size [2]int) int {
    visited := make(map[[2]int]bool)

	for {
		if !isInArea(pos, size) {
			break
		}
		visited[pos] = true
		if isNextObstacle(obstacles, pos, dir) {
			dir = turnRight(dir)
		} else {
			pos = move(pos, dir)
		}
	}
	return len(visited)
}

func isOnKnownPath(visited map[[2]int][]int, pos [2]int, dir [2]int) bool {
	for i := 0; i < len(visited[pos]); i += 2 {
		pos_dir := [2]int{visited[pos][i], visited[pos][i+1]}
		if dir == pos_dir {
			return true
		}
	}
	return false
}

func mapCopy(input map[[2]int][]int) map[[2]int][]int {
	res := make(map[[2]int][]int)
	for key, value := range input {
		for _, inner := range value {
			res[key] = append(res[key], inner)
		}
    }
	return res
}

func addVisited(visited map[[2]int][]int, pos [2]int, dir [2]int) {
	visited[pos] = append(visited[pos], dir[0])
	visited[pos] = append(visited[pos], dir[1])
}

func isLoop(visited map[[2]int][]int, obstacles map[[2]int]bool, pos [2]int, dir [2]int, size [2]int) bool {
	visited = mapCopy(visited)
	dir = turnRight(dir)

	for {
		if !isInArea(pos, size) {
			return false
		}
		if isOnKnownPath(visited, pos, dir) {
			return true
		}
		addVisited(visited, pos, dir)
		if (isNextObstacle(obstacles, pos, dir)) {
			dir = turnRight(dir)
		} else {
			pos = move(pos, dir)
		} 
	}
}

func isVisited(visited map[[2]int][]int, pos [2]int) bool {
	return len(visited[pos]) != 0
}

func part2(obstacles map[[2]int]bool, pos [2]int, dir [2]int, size [2]int) int {
    visited := make(map[[2]int][]int)
	newObs := make(map[[2]int]bool)

	for {
		if !isInArea(pos, size) {
			break
		}
		next := move(pos, dir)
		if !obstacles[next] && isInArea(next, size) && !isVisited(visited, next) && isLoop(visited, obstacles, pos, dir, size) { 
			newObs[next] = true
		}
		addVisited(visited, pos, dir)
		if (isNextObstacle(obstacles, pos, dir)) {
			dir = turnRight(dir)
		} else {
			pos = move(pos, dir)
		}
	}
	return len(newObs)
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    obstacles := make(map[[2]int]bool)
	var guard [2]int
	var dir [2]int

    scanner := bufio.NewScanner(file)
	y := 0
	width := 0
    for scanner.Scan() {
		line := scanner.Text()
		width = len(line)
		for x, c := range line {
			if c == '.' {
				continue
			} else if c == '#' {
				obstacles[[2]int{x, y}] = true
			} else {
				guard = [2]int{x, y}
				dir = getDirection(c)
			}
		}
		y++
    }
	fmt.Println("Part 1: ", part1(obstacles, guard, dir, [2]int{width, y}))
	fmt.Println("Part 2: ", part2(obstacles, guard, dir, [2]int{width, y}))
}