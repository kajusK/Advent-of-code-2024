package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func blink(data map[int]int) map[int]int {
	res := make(map[int]int)

	for stone, count := range data {
		textual := strconv.Itoa(stone)
		if stone == 0 {
			res[1] += count
		} else if len(textual) % 2 == 0 {
			a, _ := strconv.Atoi(textual[:len(textual)/2])
			b, _ := strconv.Atoi(textual[len(textual)/2:])
			res[a] += count
			res[b] += count
		} else {
			res[stone*2024] += count
		}
	}
	return res	
}

func calculate(data []int, loops int) int {
	state := make(map[int]int)
	for _, num := range data {
		state[num] += 1
	}

	for i := 0; i < loops; i++ {
		state = blink(state)
	}

	res := 0
	for _, count := range state {
		res += count
	}
	return res
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
	scanner.Scan()

	var data []int	
	for _, num := range strings.Fields(scanner.Text()) {
		value, _ := strconv.Atoi(num)
		data = append(data, value)
	}

	fmt.Println("Part 1: ", calculate(data, 25))
	fmt.Println("Part 2: ", calculate(data, 75))
}