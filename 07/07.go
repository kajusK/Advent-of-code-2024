package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func isValid(result int, cur int, data []int) bool {
	if cur > result {
		return false
	}
	if len(data) == 0 {
		return result == cur
	}

	return isValid(result, cur*data[0], data[1:]) || isValid(result, cur + data[0], data[1:])
}

func mergeNum(a int, b int) int {
	num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return num
}

func isValid2(result int, cur int, data []int) bool {
	if cur > result {
		return false
	}
	if len(data) == 0 {
		return result == cur
	}

	if isValid2(result, cur*data[0], data[1:]) || isValid2(result, cur + data[0], data[1:]) {
		return true
	}
	return isValid2(result, mergeNum(cur, data[0]), data[1:]) || isValid2(result, mergeNum(cur, data[0]), data[1:])
}

func sumValid(input [][]int, fn func(int, int, []int)bool) int {
	res := 0

	for _, line := range input {
		if fn(line[0], 0, line[1:]) {
			res += line[0]
		}
	}
	return res
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    var data [][]int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		input := strings.ReplaceAll(scanner.Text(), ":", "")
        parts := strings.Fields(input)

		var line []int
		for _, part := range parts {
        	num, _ := strconv.Atoi(part)
			line = append(line, num)
		}
		data = append(data, line)
    }

	fmt.Println("Part 1: ", sumValid(data, isValid))
	fmt.Println("Part 2: ", sumValid(data, isValid2))
}