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

func up(lines []string, i int, j int) bool {
    if i < 3 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j])
    }
    return word == "XMAS"
}

func down(lines []string, i int, j int) bool {
    if i >= len(lines)-4 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i+k][j])
    }
    return word == "XMAS"
}

func left(lines []string, i int, j int) bool {
    if j < 3 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i][j-k])
    }
    return word == "XMAS"
}

func right(lines []string, i int, j int) bool {
    if j >= len(lines[i])-3 {
	return false 
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i][j+k])
    }
    return word == "XMAS"
}

func up_left(lines []string, i int, j int) bool {
    if i < 3 || j < 3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j-k])
    }
    return word == "XMAS"
}

func up_right(lines []string, i int, j int) bool {
    if i < 3 || j >= len(lines[i])-3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i-k][j+k])
    }
    return word == "XMAS"
}

func down_left(lines []string, i int, j int) bool {
    if i >= len(lines)-4 || j < 3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i+k][j-k])
    }
    return word == "XMAS"
}

func down_right(lines []string, i int, j int) bool {
    if i >= len(lines)-4 || j >= len(lines[i])-3 {
	return false
    }
    word := ""
    for k := 0; k < 4; k++ {
	word = word + string(lines[i+k][j+k])
    }
    return word == "XMAS"
}

func part1(lines []string) int {
    count := 0
    for i, line := range lines {
	for j := 0; j < len(line); j++ {
	   if(up(lines, i, j)) {
		count++
	    } 
	   if(down(lines, i, j)) {
		count++
	    } 
	   if(left(lines, i, j)) {
		count++
	    } 
	   if(right(lines, i, j)) {
		count++
	    } 
	   if(up_left(lines, i, j)) {
		count++
	    } 
	   if(up_right(lines, i, j)) {
		count++
	    } 
	   if(down_left(lines, i, j)) {
		count++
	    } 
	   if(down_right(lines, i, j)) {
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
