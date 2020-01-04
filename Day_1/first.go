package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func fuel_required(weight int) int {
    return int(weight / 3) - 2
}

func main() {
    sum := 0
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	weight, _ := strconv.Atoi(scanner.Text())
        sum += fuel_required(weight)
    }
    fmt.Println(sum)
}
