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
