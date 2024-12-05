package utils

func Abs(val int) int {
    if val > 0 {
        return val
    }
    return -val
}

func Sum(list []int) int {
    total := 0
    for _, num := range list {
        total += num
    }
    return total
}

func All(list []int, f func(x int) bool) bool {
    for _, num := range list {
        if !f(num) {
            return false
        }
    }
    return true
}

func Count(list []int, f func(x int) bool) int {
    i := 0
    for _, num := range list {
        if f(num) {
            i++
        }
    }
    return i
}

func Remove(list []int, index int) []int {
    var result []int
    for i, item := range list {
        if i == index {
            continue
        }
        result = append(result, item)
    }
    return result
}

func Insert(list []int, index int, value int) []int {
    var result []int
    for i, item := range list {
        if i == index {
            result = append(result, value)
        }
        result = append(result, item)
    }
    return result
}

func IndexOf(list []int, value int) int {
    for i, item := range list {
        if item == value {
            return i
        }
    }
    return -1
}