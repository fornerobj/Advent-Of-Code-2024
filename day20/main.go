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
    X int
    Y int
}

type Node struct {
    pos Pos
    cost int
}

var directions []Pos = []Pos {
    {0,1},
    {0,-1},
    {1,0},
    {-1,0},
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

func BFS(grid [][]string, start, finish Pos) (int, map[Pos]int) {
    w, h := len(grid[0]), len(grid)

    q := []Node {Node{start,0}}
    distances := make(map[Pos]int)
    seen := make(map[Pos]bool)
    
    for len(q) > 0 {
	cur := q[0]
	q = q[1:]
	cx, cy := cur.pos.X, cur.pos.Y
	cc := cur.cost
	distances[cur.pos] = cur.cost

	if seen[cur.pos] {
	    continue
	}
	seen[cur.pos] = true

	if cur.pos == finish {
	    return cur.cost, distances
	}

	for _, dir := range directions {
	    dx, dy := dir.X, dir.Y
	    nx, ny := cx+dx, cy+dy

	    if nx < 0 || nx >= w || ny < 0 || ny >= h {
		continue
	    }

	    if grid[nx][ny] == "#" {
		continue
	    }

	    newNode := Node{Pos{nx, ny}, cc+1}
	    if !seen[newNode.pos] {
		q = append(q, newNode)
	    }

	}
    }
    return -1, nil
}

func part1(grid [][]string) int{
    start, finish := getStart(grid), getFinish(grid)
    noSkip, dFromStart := BFS(grid, start, finish)
    _, dFromFinish := BFS(grid, finish, start)
    count := 0
    
    for p1 := range dFromStart {
	x, y := p1.X, p1.Y
	for p2 := range dFromFinish {
	    nx, ny := p2.X, p2.Y
	    if int(math.Abs(float64(nx-x))) + int(math.Abs(float64(ny-y))) <= 2 {
		if dFromStart[Pos{x,y}] + int(math.Abs(float64(nx-x))) + int(math.Abs(float64(ny-y))) + dFromFinish[Pos{nx,ny}] <= noSkip - 100 {
		    count += 1
		}
	    }
	    
	}
    }
    return count
}

func part2(grid [][]string) int {
    start, finish := getStart(grid), getFinish(grid)
    noSkip, dFromStart := BFS(grid, start, finish)
    _, dFromFinish := BFS(grid, finish, start)
    count := 0
    
    for p1 := range dFromStart {
	x, y := p1.X, p1.Y
	for p2 := range dFromFinish {
	    nx, ny := p2.X, p2.Y
	    if int(math.Abs(float64(nx-x))) + int(math.Abs(float64(ny-y))) <= 20 {
		if dFromStart[Pos{x,y}] + int(math.Abs(float64(nx-x))) + int(math.Abs(float64(ny-y))) + dFromFinish[Pos{nx,ny}] <= noSkip - 100 {
		    count += 1
		}
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

    grid := make([][]string, 0)

    for _, line := range lines {
	grid = append(grid, strings.Split(line, ""))
    }


    fmt.Println("part 1:", part1(grid))
    fmt.Println("part 2:", part2(grid))

}
