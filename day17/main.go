package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func combo(A, B, C, operand int) int {
    if operand < 4 {
	return operand
    }
    if operand == 4 {
	return A
    }
    if operand == 5 {
	return B
    }
    if operand == 6 {
	return C
    }
    return -1
}

func runProgram(A, B, C int, instructions []int) ([]int, bool) {
    out := make([]int, 0)
    for i:= 0; i < len(instructions); i += 2 {
	if instructions[i] == 0 {
	    operand := combo(A, B, C, instructions[i+1]) 
	    A = int(float64(A)/math.Pow(2, float64(operand)))
	} else if instructions[i] == 1 {
	    B = B^instructions[i+1]
	} else if instructions[i] == 2 {
	    B = combo(A, B, C, instructions[i+1]) % 8
	} else if instructions[i] == 3 {
	    if A == 0 {
		continue
	    }
	    i = instructions[i+1]-2
	} else if instructions[i] == 4 {
	    B = B^C
	} else if instructions[i] == 5 {
	    out = append(out, combo(A,B,C,instructions[i+1])%8)
	} else if instructions[i] == 6 {
	    operand := combo(A, B, C, instructions[i+1]) 
	    B = int(float64(A)/math.Pow(2, float64(operand)))
	} else if instructions[i] == 7 {
	    operand := combo(A, B, C, instructions[i+1]) 
	    C = int(float64(A)/math.Pow(2, float64(operand)))
	}

    }

    if len(out) != len(instructions) {
	return out, false
    }
    for i, num := range out {
	if num != instructions[i] {
	    return out, false
	}
    }
    return out, true
}

func part1(A, B, C int, instructions []int) string {
    out, _ := runProgram(A,B,C,instructions)
    out_string := ""
    for i, num := range out {
	out_string += strconv.Itoa(num)
	if i != len(out)-1 {
	    out_string += ","
	}
    }
    return out_string
}

func part2(B, C int, instructions []int) int {
    for A := 0;;A++ {
	if A % 1000000 == 0 {
	    fmt.Println(A)
	}
	_, isMatching := runProgram(A,B,C,instructions)
	if isMatching {
	    return A
	}
    } 
}

func parse (input string) (int, int, int, []int) {
    split := strings.Split(input, "\n\n")
    registers := strings.Split(split[0], "\n")
    instructionString := split[1]

    regAString := strings.TrimPrefix(registers[0], "Register A: ")
    A, err := strconv.Atoi(regAString)
    check(err)

    regBString := strings.TrimPrefix(registers[1], "Register B: ")
    B, err := strconv.Atoi(regBString)
    check(err)

    regCString := strings.TrimPrefix(registers[2], "Register C: ")
    C, err := strconv.Atoi(regCString)
    check(err)

    listString :=strings.TrimSpace(strings.TrimPrefix(instructionString, "Program: "))
    nums := make([]int, 0)
    for _, n := range strings.Split(listString, ",") {
	ns := string(n)
	ni, err := strconv.Atoi(ns)
	check(err)
	nums = append(nums, ni)
    }

    return A, B, C, nums
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)

    A, B, C, instructions := parse(input)

    fmt.Println("part 1:", part1(A, B, C, instructions))
    fmt.Println("part 2:", part2(B, C, instructions))

}
