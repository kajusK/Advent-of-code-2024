package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

func getResult(input string) int {
	re, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(string(input), -1)

	res := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		res += x*y
	}
	return res
}

func part2(input string) int {
	remaining := "do()"+input
	res := 0

	for {
		start := 0
		stop := len(remaining) 
		if strings.Index(remaining, "don't()") != -1 {
			stop = strings.Index(remaining, "don't()") + len("don't()")
		} 
		if strings.Index(remaining, "do()") != -1 {
			start = strings.Index(remaining, "do()")
		}
		if start < stop {
			res += getResult(remaining[start:stop])
		}

		remaining = remaining[stop:]
		if len(remaining) == 0 {
			break
		}
	}
	return res
}

func main() {
	input, _ := os.ReadFile("input.txt")

	fmt.Println("Part 1: ", getResult(string(input)))
	fmt.Println("Part 2: ", part2(string(input)))
}