package main

import (
	"fmt"
	"os"
	"strconv"
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

func getScore(grid [][]int, pos Pos, reached map[Pos]bool) int {
    num := grid[pos.R][pos.C]
    reached[pos] = true
    if num == 9 {
	return 1
    }
    width := len(grid[0])
    height := len(grid)

    score := 0
    directions := []Pos{Pos{pos.R+1, pos.C}, Pos{pos.R-1, pos.C}, Pos{pos.R, pos.C-1}, Pos{pos.R, pos.C+1}}
    noNeighbors := true
    for _, nPos := range directions {
	if !(nPos.R >= 0 && nPos.R < height && nPos.C >= 0 && nPos.C < width) {
	    continue
	}
	if grid[nPos.R][nPos.C] != num + 1 {
	    continue
	}
	if reached[nPos] {
	    continue
	}
	noNeighbors = false
	score += getScore(grid, nPos, reached)
    }

    if noNeighbors {
	return 0
    }
    return score
}

func getScore2(grid [][]int, pos Pos) int {
    num := grid[pos.R][pos.C]
    if num == 9 {
	return 1
    }
    width := len(grid[0])
    height := len(grid)

    score := 0
    directions := []Pos{Pos{pos.R+1, pos.C}, Pos{pos.R-1, pos.C}, Pos{pos.R, pos.C-1}, Pos{pos.R, pos.C+1}}
    noNeighbors := true
    for _, nPos := range directions {
	if !(nPos.R >= 0 && nPos.R < height && nPos.C >= 0 && nPos.C < width) {
	    continue
	}
	if grid[nPos.R][nPos.C] != num + 1 {
	    continue
	}
	noNeighbors = false
	score += getScore2(grid, nPos)
    }

    if noNeighbors {
	return 0
    }
    return score
}
func part1(grid [][]int) int {
    count := 0
    for r, row := range grid {
	for c, num := range row {
	    if num == 0 {
		count += getScore(grid, Pos{r, c}, make(map[Pos]bool))
	    }
	}
    }
    return count
}


func part2(grid [][]int) int {
    count := 0
    for r, row := range grid {
	for c, num := range row {
	    if num == 0 {
		count += getScore2(grid, Pos{r, c})
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
    lines = lines[:len(lines)-1]
    grid := make([][]int, 0)
    for _, line := range lines {
	n := len(line)
	row := make([]int, n)	
	for i, c := range line  {
	    m, err := strconv.Atoi(string(c)) 
	    check(err)
	    row[i] = m
	}
	grid = append(grid, row)
    }

    fmt.Println("part 1:", part1(grid))
    fmt.Println("part 2:", part2(grid))

}
