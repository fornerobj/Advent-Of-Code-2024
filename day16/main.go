package main

import (
	"fmt"
	"math"
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

func findLowestNeighbor(visited map[Robot]bool, costs map[Robot]int) Robot {
    minCost := math.MaxInt64
    minRobot := Robot{Pos{-1,-1}, -1}
    for robot, seen := range visited {
	if !seen {
	    if costs[robot] < minCost {
		minRobot = robot
		minCost = costs[robot]
	    }
	}
    }
    return minRobot
}

func listEmpty(visited map[Robot]bool) bool {
    for _, seen := range visited {
	if !seen {
	    return false
	}
    } 
    return true
}

func findPath(grid [][]string) int {
    visited := make(map[Robot]bool)
    costs := make(map[Robot]int)
    for r, row := range grid {
	for c, _ := range row {
	    for d := 0; d < 4; d++ {
		pos := Pos{r,c}
		tempRobot := Robot{pos, d}
		visited[tempRobot] = false
		costs[tempRobot] = math.MaxInt64
	    }
	}
    }

    cur := Robot{getStart(grid), 1}
    finish := getFinish(grid)

    costs[cur] = 0

    for !listEmpty(visited) {
	visited[cur] = true
	if cur.Dir == -1 {
	    break
	}
	if grid[cur.pos.R][cur.pos.C] == "#" {
	    cur = findLowestNeighbor(visited, costs)
	    continue
	}
	left := Robot{cur.pos, cur.Dir-1} 
	if left.Dir == -1 {
	    left.Dir = 3
	}
	if 1000+costs[cur] < costs[left] {
	    costs[left] = 1000 + costs[cur]
	}

	right := Robot{cur.pos, (cur.Dir + 1) % 4}
	if 1000 + costs[cur] < costs[right] {
	    costs[right] = 1000 + costs[cur]
	}

	dPos := dirToDirection[cur.Dir]
	forward := Robot{Pos{cur.pos.R+dPos.R, cur.pos.C+dPos.C}, cur.Dir}
	if 1 + costs[cur] < costs[forward] {
	    costs[forward] = 1+costs[cur]
	}

	cur = findLowestNeighbor(visited, costs)
    }
    
    minCost := math.MaxInt64
    for d := 0; d < 4; d++ {
	r := Robot{finish, d}
	if costs[r] < minCost {
	    minCost = costs[r]
	}
    }

    return minCost
}

func part1(grid [][]string) int {
    minCost := findPath(grid)
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
