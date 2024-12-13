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

func getAreaAndPerimeter(pos Pos, plant string,  _map [][]string, seen map[Pos]bool) (int, int) {
    w := len(_map[0])
    h := len(_map)

    if pos.R < 0 || pos.R >= h || pos.C < 0 || pos.C >= w {
	return 0, 1
    }

    if _map[pos.R][pos.C] != plant {
	return 0, 1
    }

    if seen[pos] {
	return 0, 0
    }

    seen[pos] = true

    directions := []Pos{
	{1,0},
	{-1,0},
	{0,1},
	{0,-1},
    }

    res := 0
    res2 := 0
    
    for _, p := range directions {
	dr, dc := p.R, p.C
	a, p := getAreaAndPerimeter(Pos{pos.R+dr,pos.C+dc}, plant, _map, seen)
	res += a
	res2 += p
    }

    return 1 + res, res2
}

func part1(_map [][]string) int {
    seen := make(map[Pos]bool)
    count := 0

    for r, row := range _map {
	for c, plant := range row {
	   if !seen[Pos{r, c}] {
		area, perimeter := getAreaAndPerimeter(Pos{r,c}, plant, _map, seen)
		count += area*perimeter
	    }  
	}
    }
    return count
}

func part2(_map [][]string) int {
    count := 0
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]
    _map := make([][]string, len(lines))
    for i, line := range lines {
	_map[i] = strings.Split(line, "")
    }

    fmt.Println("part 1:", part1(_map))
    fmt.Println("part 2:", part2(_map))

}
