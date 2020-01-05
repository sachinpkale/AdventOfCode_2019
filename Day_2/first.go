package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func run_intcode(commands []int) []int {
	for i := 0; commands[i] != 99; i+= 4 {
	    if commands[i] == 1 {
		commands[commands[i + 3]] = commands[commands[i + 1]] +  commands[commands[i + 2]]
	    }
	    if commands[i] == 2 {
		commands[commands[i + 3]] = commands[commands[i + 1]] *  commands[commands[i + 2]]
	    }
	}
    return commands
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    input_s := []string{}
    for scanner.Scan() {
	input_s = strings.Split(scanner.Text(), ",")
    }
    input := []int{}
    for _, i := range input_s {
        j, _ := strconv.Atoi(i)
	input = append(input, j) 
    }
    input[1] = 12
    input[2] = 2

    fmt.Println(run_intcode(input))
}
