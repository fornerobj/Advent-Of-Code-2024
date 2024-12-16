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

type Pos struct {
    R int
    C int
}

var arrowToMap = map[string]Pos{
    ">": {0, 1},
    "<": {0, -1},
    "^": {-1, 0},
    "v": {1, 0},
}

func printMap(grid[][]string) {
    for _, line := range grid {
	for _, c := range line {
	    fmt.Print(c)
	}
	fmt.Print("\n")
    }
}

func expandMap(grid [][]string) [][]string {
    newMap := make([][]string, 0)

    for _, line := range grid {
	newLine := make([]string, 0)
	for _, c := range line {
	    if c == "#" {
		newLine = append(newLine, "#")
		newLine = append(newLine, "#")
	    } else if c == "O" {
		newLine = append(newLine, "[")
		newLine = append(newLine, "]")
	    } else if c == "@" {
		newLine = append(newLine, "@")
		newLine = append(newLine, ".")
	    } else {
		newLine = append(newLine, ".")
		newLine = append(newLine, ".")
	    }
	}
	newMap = append(newMap, newLine)
    }
    return newMap
}

func findRobot(grid [][]string) Pos {
    height := len(grid)
    width := len(grid[0])
    for row := 0; row < height; row ++ {
	for col := 0; col < width; col++ {
	    if grid[row][col] == "@" {
		return Pos{row, col}
	    }
	}
    }
    return Pos{-1,-1}
}

func moveRobot(grid [][]string, instruction string) [][]string {
    dpos := arrowToMap[instruction]
    robot := findRobot(grid)
    pos := robot
    if grid[robot.R + dpos.R][robot.C + dpos.C] == "#" {
	return grid
    }
    if grid[robot.R + dpos.R][robot.C + dpos.C] == "." {
	grid[robot.R+dpos.R][robot.C+dpos.C] = "@"
	grid[robot.R][robot.C] = "."
	return grid
    }
    for ;; {
	pos = Pos{pos.R+dpos.R, pos.C + dpos.C}
	if grid[pos.R][pos.C] == "#" {
	    break
	}
	if grid[pos.R][pos.C] == "." {
	    for ;; {
		if grid[pos.R][pos.C] == "@" {
		    break
		}
		grid[pos.R][pos.C] = "O"
		pos = Pos{pos.R-dpos.R, pos.C-dpos.C}
	    }
	    grid[robot.R+dpos.R][robot.C+dpos.C] = "@"
	    grid[robot.R][robot.C] = "."
	    break
	}
    }
    return grid
}

func moveRobotBig(grid [][]string, instruction string) [][]string {
    dpos := arrowToMap[instruction]
    robot := findRobot(grid)
    pos := robot
    if grid[robot.R + dpos.R][robot.C + dpos.C] == "#" {
	return grid
    }
    if grid[robot.R + dpos.R][robot.C + dpos.C] == "." {
	grid[robot.R+dpos.R][robot.C+dpos.C] = "@"
	grid[robot.R][robot.C] = "."
	return grid
    }
    for ;; {
	pos = Pos{pos.R+dpos.R, pos.C + dpos.C}
	if grid[pos.R][pos.C] == "#" {
	    break
	}
	if grid[pos.R][pos.C] == "." {
	    for i:=0;;i++ {
		if grid[pos.R][pos.C] == "@" {
		    break
		}
		if i%2!=0 {
		    grid[pos.R][pos.C] = "]"
		}else {
		    grid[pos.R][pos.C] = "["
		}
		pos = Pos{pos.R-dpos.R, pos.C-dpos.C}
	    }
	    grid[robot.R+dpos.R][robot.C+dpos.C] = "@"
	    grid[robot.R][robot.C] = "."
	    break
	}
    }
    return grid
}

func part1(grid [][]string, instructions []string) int {
    count := 0
    for _, ins := range instructions {
	grid = moveRobot(grid, ins)
    }

    for i, line := range grid {
	for j, c := range line {
	    if c == "O" {
		count += 100*i + j
	    }
	}
    }
    return count
}

func part2(grid [][]string, instructions []string) int {
    count := 0
    printMap(grid)
    for _, ins := range instructions {
	grid = moveRobotBig(grid, ins)
	printMap(grid)
    }

    for i, line := range grid {
	for j, c := range line {
	    if c == "[" {
		count += 100*i + j
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
    split := strings.Split(input, "\n\n")
    gridString := split[0]
    instructionsString := strings.ReplaceAll(split[1], "\n", "")
    grid := make([][]string, 0)
    for _, line := range strings.Split(gridString, "\n") {
	grid = append(grid, strings.Split(line, ""))
    }
    
    bigGrid := expandMap(grid)
    instructions := strings.Split(instructionsString, "")

    fmt.Println("part 1:", part1(grid, instructions))
    fmt.Println("part 2:", part2(bigGrid, instructions))

}
