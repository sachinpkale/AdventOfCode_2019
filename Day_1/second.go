package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func fuel_required(weight int) int {
    fuel := 0
    for weight > 0 {
	weight = int(weight / 3) - 2
	if(weight >= 0) {
	    fuel += weight
	} else {
	    break;
	}
    }
    return fuel
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
