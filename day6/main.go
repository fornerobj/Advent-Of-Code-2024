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

type character struct {
    x_pos int
    y_pos int
    // 1=N, 2=E, 3=S, 4=W
    orientation int
}

type Pos struct {
    X int
    Y int
}

func find_starting_pos(_map []string) character {
    var c1 character
    for i, row := range _map {
	for j, c := range row {
	    if string(c) == "<" || string(c) == ">" || string(c) == "^" || string(c) == "v" {
		c1.x_pos = j
		c1.y_pos = i
		if string(c) == "<" {
		    c1.orientation = 4
		}
		if string(c) == "^" {
		    c1.orientation = 1
		}
		if string(c) == ">" {
		    c1.orientation = 2
		}
		if string(c) == "v" {
		    c1.orientation = 3
		}
	    }
	}
    }
    return c1
}

func isClear(_map []string, x, y int) bool{
    if string(_map[y][x]) == "#" {
	return false
    }  
    return true
}

func getNewPos(c1 character) (int, int) {
    switch c1.orientation {
    case 1:
	return c1.x_pos, c1.y_pos-1
    case 2:
	return c1.x_pos+1, c1.y_pos
    case 3:
	return c1.x_pos, c1.y_pos+1
    case 4:
	return c1.x_pos-1, c1.y_pos
    default:
	return -1, -1
    }
}

func rotate90(c1 character) character {
    switch c1.orientation {
    case 1:
	c1.orientation = 2
    case 2:
	c1.orientation = 3
    case 3:
	c1.orientation = 4
    case 4:
	c1.orientation = 1
    }
    return c1
}

func part1(_map []string) int {
    count := 1
    guard := find_starting_pos(_map)
    width := len(_map[0])
    height := len(_map) - 1

    visited := make(map[Pos]bool)
    visited[Pos{guard.x_pos, guard.y_pos}] = true
    
    for ;; {
	tempX, tempY := getNewPos(guard)
	if tempX < 0 || tempX >= width || tempY < 0 || tempY >= height {
	    break
	}
	if(isClear(_map, tempX, tempY)) {
	    guard.x_pos = tempX
	    guard.y_pos = tempY
	    if(!visited[Pos{tempX, tempY}]) {
		count ++
		visited[Pos{tempX, tempY}] = true
	    }
	}else {
	    guard = rotate90(guard)
	}
    }
    return count
}


func part2(lines []string) int {
    count := 0
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
 //    for _, line := range lines {
	// fmt.Println(line)
 //    }

    fmt.Println("part 1:", part1(lines))
    fmt.Println("part 2:", part2(lines))

}
