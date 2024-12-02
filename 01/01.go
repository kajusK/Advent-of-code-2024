package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "sort"
    "aoc/utils"
)

func totalCount(list []int, search int) int {
    count := 0

    for _, num := range list {
        if num == search {
            count++
        }
    }
    return count
}

func part1(left []int, right []int) int {
    var result []int

    sort.Ints(left)
    sort.Ints(right)
    for i := 0; i < len(left); i++ {
        result = append(result, utils.Abs(left[i] - right[i]))
    }

    return utils.Sum(result)
}

func part2(left []int, right []int) int {
    total := 0
    for _, num := range left {
        total += num * totalCount(right, num)
    }
    return total
}

func main() {

    file, _ := os.Open("input.txt")
    defer file.Close()

    var left, right []int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        parts := strings.Fields(scanner.Text())
        leftnum, _ := strconv.Atoi(parts[0])
        rightnum, _ := strconv.Atoi(parts[1])
        left = append(left, leftnum)
        right = append(right, rightnum)
    }


    fmt.Println("Part 1:", part1(left, right))
    fmt.Println("Part 2:", part2(left, right))
}