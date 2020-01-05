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

func copy_array(input []int) []int {
    output := []int{}
    for _, i := range input {
        output = append(output, i)
    }
    return output
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
    found := false
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
	    temp_input := copy_array(input)
            temp_input[1] = i
	    temp_input[2] = j
	    output := run_intcode(temp_input)
	    if output[0] == 19690720 {
		fmt.Println(i, j)
		found = true
                break
	    }
	}
	if found {
	    break
	}
    }
}
