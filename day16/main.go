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

type Robot struct {
    pos Pos
    // 0up 1right 2down 3left
    Dir int
}

func getStart(grid [][]string) Pos {
    for i, r := range grid {
	for j, c := range r {
	    if c == "S" {
		return Pos{i,j}
	    }
	}
    }
    return Pos{-1,-1}
}

func getFinish(grid [][]string) Pos {
    for i, r := range grid {
	for j, c := range r {
	    if c == "E" {
		return Pos{i,j}
	    }
	}
    }
    return Pos{-1,-1}
}

var dirs []Pos = []Pos {
    {0,1},
    {0,-1},
    {1,0},
    {-1,0},
}

var dirToDirection map[int]Pos = map[int]Pos{
    0: {-1, 0},
    1: {0, 1},
    2: {1, 0},
    3: {0, -1},
}

// func findPath(grid [][]string, cur, goal Pos, seen map[Pos]bool) int {
//     if cur == goal {
// 	return 0
//     }
//     cost := 100000
//     for _, dir := range dirs {
// 	newPos := Pos{cur.R+dir.R, cur.C+dir.C}
// 	if grid[newPos.R][newPos.C] != "#" && !seen[newPos] {
// 	    seen[newPos] = true
// 	    cost = min(cost, 1+findPath(grid, newPos, goal, seen)) 
// 	}
//     } 
//     return cost
// }


func findPath(grid [][]string, cur Robot, goal Pos, seen map[Robot]bool) (int, []Pos) {
    if cur.pos == goal {
        return 0, []Pos{goal}
    }
    
    cost := 100000
    var bestPath []Pos

    newPos := Pos{cur.pos.R + dirToDirection[cur.Dir].R, cur.pos.C + dirToDirection[cur.Dir].C}
    forward := Robot{newPos, cur.Dir}

    if grid[newPos.R][newPos.C] != "#" && !seen[forward] {
	seen[forward] = true
	subCost, subPath := findPath(grid, forward, goal, seen)
	totalCost := 1 + subCost
	if totalCost < cost {
	    cost = totalCost
	    bestPath = append([]Pos{newPos}, subPath...)
	}
	seen[forward] = false
    }
    
    rotatedC := Robot{cur.pos, (cur.Dir + 1) % 4}
    if !seen[rotatedC] {
	seen[rotatedC] = true
	subCost, subPath := findPath(grid, rotatedC, goal, seen)
	totalCost := 1000+subCost
	if totalCost < cost {
	    cost = totalCost
	    bestPath = append([]Pos{cur.pos}, subPath...)
	}
	seen[rotatedC] = false
    }

    rotateCC := cur.Dir - 1
    if rotateCC == -1 {
	rotateCC = 3
    }
    rotatedCC := Robot{cur.pos, rotateCC}
    if !seen[rotatedCC] {
	seen[rotatedCC] = true
	subCost, subPath := findPath(grid, rotatedCC, goal, seen)
	totalCost := 1000+subCost
	if totalCost < cost {
	    cost = totalCost
	    bestPath = append([]Pos{cur.pos}, subPath...)
	}
	seen[rotatedCC] = false
    }

    return cost, bestPath
}

func part1(grid [][]string) int {
    start := getStart(grid) 
    goal := getFinish(grid)
    minCost, bestPath := findPath(grid, Robot{start, 0}, goal, make(map[Robot]bool))
    fmt.Println(bestPath)
    return minCost
}

func part2() int {
    return 1
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    grid := make([][]string, 0)

    for _, line := range lines {
	grid = append(grid, strings.Split(line, ""))
    }

    fmt.Println("part 1:", part1(grid))
    // fmt.Println("part 2:", part2(grid))

}
