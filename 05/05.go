package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"aoc/utils"
)

func isRuleOk(update []int, rule []int) bool {
	left := utils.IndexOf(update, rule[0])
	right := utils.IndexOf(update, rule[1])
	if left == -1 || right == -1 {
		return true
	}
	if left < right {
		return true
	}
	return false
}

func isValid(update []int, rules [][]int) bool {
	for _, rule := range rules {
		if !isRuleOk(update, rule) {
			return false
		}
	}
	return true
}

func getMiddle(update []int) int {
	return update[len(update)/2]
}

func insertInfront(list []int, from int, infront int) []int {
	value := list[from]
	list = utils.Remove(list, from)
	if from < infront {
		infront--
	}
	return utils.Insert(list, infront, value)
}

func fixOrder(update []int, rules [][]int) []int {
	for {
		if isValid(update, rules) {
			return update
		}

		for _, rule := range rules {
			if isRuleOk(update, rule) {
				continue
			}
			left := utils.IndexOf(update, rule[0])
			right := utils.IndexOf(update, rule[1])
			update = insertInfront(update, left, right)
		}
	}
}

func part1(updates [][]int, rules [][]int) int {
	res := 0
	for _, update := range updates {
		if !isValid(update, rules) {
			continue
		}
		res += getMiddle(update)
	}
	return res
}

func part2(updates [][]int, rules [][]int) int {
	res := 0
	for _, update := range updates {
		if isValid(update, rules) {
			continue
		}
		new := fixOrder(update, rules)
		res += getMiddle(new)
	}
	return res
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    var rules [][]int
	var updates [][]int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		spl := strings.Split(line, "|")
		first, _ := strconv.Atoi(spl[0])
		second, _ := strconv.Atoi(spl[1])
		rules = append(rules, []int{first, second})
    }

    for scanner.Scan() {
		spl := strings.Split(scanner.Text(), ",")
		var res []int
		for _, item := range spl {
			num, _ := strconv.Atoi(item)
			res = append(res, num)
		}
		updates = append(updates, res)
	}

	fmt.Println("Part 1: ", part1(updates, rules))
	fmt.Println("Part 2: ", part2(updates, rules))
}