package main

import (
	"os"
	"bufio"
	"fmt"
)

type File struct {
	id int
	size int
	orig_pos int
}

func find_space(data []int, size int) int {
	space_size := 0
	for i := 0; i < len(data); i++ {
		if space_size >= size {
			return i - space_size
		}
		if data[i] != -1 {
			space_size = 0
			continue
		}
		space_size++
	}
	return -1
}

func move_file(data []int, file File, pos int) []int {
	for i := 0; i < file.size; i++ {
		data[pos+i] = file.id	
		data[file.orig_pos+i] = -1
	}
	return data
}

func checksum(data []int) int {
	res := 0
	for i, val := range data {
		if val != -1 {
			res += i*val
		}
	}
	return res
}

func copy_arr(in []int) []int {
	res := make([]int, len(in))
	copy(res, in)
	return res
}

func part2(data []int, files []File) int {
	for i := len(files) - 1; i > 0; i-- {
		file := files[i]
		space := find_space(data, file.size)
		if space == -1 || space > file.orig_pos {
			continue
		}
		data = move_file(data, file, space)
	}
	return checksum(data)
}

func part1(data []int) int {
	for pos := len(data) - 1; pos > 0; pos-- {
		space := find_space(data, 1)
		if space == -1 || space > pos {
			break
		}
		data[space] = data[pos]
		data[pos] = -1 
	}
	return checksum(data)
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text() + "0"

	var data []int
	var files []File

	pos := 0
	for id := 0; id < len(input)/2; id++ {
		size := int(input[id*2] - '0')
		space := int(input[id*2+1] - '0')
		files = append(files, File{size: size, id: id, orig_pos: pos})

		for i := 0; i < size; i++ {
			data = append(data, id)
		}
		for i := 0; i < space; i++ {
			data = append(data, -1)
		}
		pos += size + space
	}

	fmt.Println("Part 1: ", part1(copy_arr(data)))
	fmt.Println("Part 2: ", part2(copy_arr(data), files))
}