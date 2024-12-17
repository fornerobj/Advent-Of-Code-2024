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

func copyGrid(grid [][]string) [][]string {
    newGrid := make([][]string, len(grid))

    for i := 0; i < len(grid); i++ {
	line := grid[i]
	newLine := make([]string, len(line))

	for j := 0; j < len(line); j++ {
	    newLine[j] = line[j]
	}

	newGrid[i] = newLine
    }

    return newGrid
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

func inSlice(slice []Pos, pos Pos) bool {
    for i:=0; i < len(slice); i++ {
	if slice[i] == pos {
	    return true
	}
    }
    return false
}

func part2(grid [][]string, instructions []string) int {
    count := 0
    cPos := findRobot(grid)
    for _, instruction := range instructions {
	dr, dc := arrowToMap[instruction].R, arrowToMap[instruction].C
	positionsToMove := make([]Pos, 0)
	positionsToMove = append(positionsToMove, cPos)
	i := 0
	cantMove := false

	for i < len(positionsToMove) {
	    r, c := positionsToMove[i].R, positionsToMove[i].C
	    nr, nc := r+dr, c+dc
	    if grid[nr][nc] == "[" || grid[nr][nc] == "]" {
		if !inSlice(positionsToMove, Pos{nr,nc}) {
		    positionsToMove = append(positionsToMove, Pos{nr, nc})
		}
		if grid[nr][nc] == "[" {
		    if !inSlice(positionsToMove, Pos{nr,nc+1}) {
			positionsToMove = append(positionsToMove, Pos{nr, nc+1})
		    }
		}
		if grid[nr][nc] == "]" {
		    if !inSlice(positionsToMove, Pos{nr,nc-1}) {
			positionsToMove = append(positionsToMove, Pos{nr, nc-1})
		    }
		}
	    }else if grid[nr][nc] == "#" {
		cantMove = true
		break
	    }
	    i++
	}
	if cantMove {
	    continue
	}
	newGrid := copyGrid(grid)

	for _, pos := range positionsToMove {
	    newGrid[pos.R][pos.C] = "."
	}
	for _, pos := range positionsToMove {
	    newGrid[pos.R+dr][pos.C+dc] = grid[pos.R][pos.C]
	}
	grid = newGrid

	cPos.R += dr
	cPos.C += dc
	 
    }
    for r, row := range grid {
	for c, col := range row {
	    if col == "[" {
		count += 100*r + c
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
