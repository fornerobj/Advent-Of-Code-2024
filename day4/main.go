package main

import (
	"fmt"
	"os"
	"strings"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func vertical(lines []string, i int, j int) bool {
    if i < 3 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j])
    }
    return word == "XMAS" || word == "SAMX"
}

func horizontal(lines []string, i int, j int) bool {
    if j < 3 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i][j-k])
    }
    return word == "XMAS" || word == "SAMX"
}

func diagonal_left_to_right(lines []string, i int, j int) bool {
    if i < 3 || j < 3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j-k])
    }
    return word == "XMAS" || word == "SAMX"
}

func diagonal_right_to_left(lines []string, i int, j int) bool {
    if i < 3 || j >= len(lines[i])-3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j+k])
    }
    return word == "XMAS" || word == "SAMX"
}

func part1(lines []string) int {
    count := 0
    for i, line := range lines {
	for j := 0; j < len(line); j++ {
	   if(vertical(lines, i, j)) {
		count++
	    } 
	   if(horizontal(lines, i, j)) {
		count++
	    } 
	   if(diagonal_left_to_right(lines, i, j)) {
		count++
	    } 
	   if(diagonal_right_to_left(lines, i, j)) {
		count++
	    } 
	}
    }
    return count
}

func mas_down_left(lines []string, i int, j int) bool {
    if i >= len(lines)-3 || j < 2 {
	return false
    }
    word := ""
    for k := 0; k < 3; k++ {
	word = word + string(lines[i+k][j-k])
    }
    return word == "MAS" || word == "SAM"
}

func mas_down_right(lines []string, i int, j int) bool {
    if i >= len(lines)-3 || j >= len(lines[i])-2 {
	return false
    }
    word := ""
    for k := 0; k < 3; k++ {
	word = word + string(lines[i+k][j+k])
    }
    return word == "MAS" || word == "SAM"
}

func part2(lines []string) int {
    count := 0
    for i, line := range lines {
	for j := 0; j < len(line); j++ {
	    if mas_down_right(lines, i, j) && mas_down_left(lines, i, j+2) {
		count++
	    }
	}
    }
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")

    fmt.Println("part 1:", part1(lines))
    fmt.Println("part 2:", part2(lines))

}
