package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"aoc/utils"
)

func getDistances(report []int) []int {
	var res []int
	var prev = report[0]

	for _, num := range report[1:] {
		res = append(res, num - prev)
		prev = num
	}
	return res
}

func isSafe(report []int) bool {
	var distances = getDistances(report)

	if !utils.All(distances, func(x int) bool {return x > 0}) && !utils.All(distances, func(x int) bool {return x < 0}) {
		return false
	}
	return utils.All(distances, func(x int) bool { return utils.Abs(x) <= 3 && utils.Abs(x) >= 1 })
}

func part1(reports [][]int) int {
	var res = 0
	for _, report := range reports {
		if isSafe(report) {
			res++
		}
	}
	return res
}

func part2(reports [][]int) int {
	var res = 0
	for _, report := range reports {
		if isSafe(report) {
			res++
			continue
		}
		// brutforce it...
		for i, _ := range report {
			if isSafe(utils.Remove(report, i)) {
				res++
				break
			}

		}
	}
	return res
}

func main() {

    file, _ := os.Open("input.txt")
    defer file.Close()

    var reports [][]int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        numbers := strings.Fields(scanner.Text())
		var items []int
		for _, item := range numbers {
			num, _ := strconv.Atoi(item)
			items = append(items, num)	
		} 
		reports = append(reports, items)
    }

    fmt.Println("Part 1:", part1(reports))
    fmt.Println("Part 2:", part2(reports))
}