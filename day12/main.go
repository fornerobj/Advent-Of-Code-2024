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
var directions []Pos = []Pos{
    {-1,0}, //up
    {0,1}, //right
    {1,0}, //down
    {0,-1}, //left
}

func dfs(pos Pos, _map [][]string, seen map[Pos]bool) []Pos {
    stack := []Pos{pos}
    w, h := len(_map[0]), len(_map)
    localSeen := make(map[Pos]bool)
    localSeen[pos] = true
    seen[pos] = true

    for len(stack) > 0 {
	cur := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	cr, cc := cur.R, cur.C
	curLetter := _map[cr][cc]
	for _, dir := range directions {
	    dr, dc := dir.R, dir.C 
	    nr, nc := cr+dr, cc+dc
	    if nr < 0 || nr >= h || nc < 0 || nc >= w {
		continue
	    }
	    if _map[nr][nc] != curLetter {
		continue
	    }
	    if seen[Pos{nr,nc}] {
		continue
	    }
	    seen[Pos{nr,nc}] = true
	    localSeen[Pos{nr,nc}] = true
	    stack = append(stack, Pos{nr,nc})
	}
    }
    group := make([]Pos, len(localSeen))
    i := 0
    for foundPos := range localSeen {
	group[i] = foundPos
	i++	
    }

    return group
}

func getPerimeter(group []Pos, _map [][]string) int {
    letter := _map[group[0].R][group[0].C]
    w, h := len(_map[0]), len(_map)
    perimeter := 0
    for _, cur := range group {
	cr, cc := cur.R, cur.C
	for _, dir := range directions {
	    dr, dc := dir.R, dir.C
	    nr, nc := cr+dr, cc+dc
	    if nr < 0 || nr >= h || nc < 0 || nc >=w {
		perimeter += 1
		continue
	    }
	    if _map[nr][nc] != letter {
		perimeter += 1
	    }
	}
    } 
    return perimeter
}

func getGroups(_map [][]string) [][]Pos {
    seen := make(map[Pos]bool)
    groups := make([][]Pos, 0)

    for r, row := range _map {
	for c, _ := range row {
	    if !seen[Pos{r, c}] {
		g := dfs(Pos{r,c}, _map, seen)
		groups = append(groups, g)
	    }  
	}
    }
    return groups
}

func part1(_map [][]string) int {
    count := 0
    groups := getGroups(_map)

    for _, group := range groups {
	area := len(group)
	perimeter := getPerimeter(group, _map)
	count += area*perimeter
    }
    return count
}

type Edge struct {
    pos Pos
    //0 up, 1 right, 2 down, 3 left
    orientation int
}

func getEdges (_map [][]string, pos Pos) []Edge{
    w, h := len(_map[0]), len(_map)
    letter := _map[pos.R][pos.C]
    edges := make([]Edge, 0)
    for i, dir := range directions {
	nr, nc := pos.R + dir.R, pos.C + dir.C
	if nr < 0 || nr >= h || nc < 0 || nc >= w {
	    edges = append(edges, Edge{pos, i})
	    continue
	}
	if _map[nr][nc] != letter {
	    edges = append(edges, Edge{pos, i})
	}
    }

    return edges
}

func getSides(_map [][]string, group []Pos) int {
    edges := make(map[Edge]bool)

    for _, pos := range group {
	newEdges := getEdges(_map, pos)
	for _, e := range newEdges {
	    edges[e] = true
	}
    }

    for e := range edges {
	if !edges[e] {
	    continue
	}
	if e.orientation == 0 || e.orientation == 2 {
	    nc := e.pos.C
	    for ;; {
		nc += 1
		newEdge := Edge{Pos{e.pos.R, nc}, e.orientation}
		if _, exists := edges[newEdge]; !exists {
		    break
		} else {
		    edges[newEdge] = false
		}
	    }
	    nc = e.pos.C
	    for ;; {
		nc -= 1
		newEdge := Edge{Pos{e.pos.R, nc}, e.orientation}
		if _, exists := edges[newEdge]; !exists {
		    break
		} else {
		    edges[newEdge] = false
		}
	    }
	} else {
	    nr := e.pos.R
	    for ;; {
		nr += 1
		newEdge := Edge{Pos{nr, e.pos.C}, e.orientation}
		if _, exists := edges[newEdge]; !exists {
		    break
		} else {
		    edges[newEdge] = false
		}
	    }
	    nr = e.pos.R
	    for ;; {
		nr -= 1
		newEdge := Edge{Pos{nr, e.pos.C}, e.orientation}
		if _, exists := edges[newEdge]; !exists {
		    break
		} else {
		    edges[newEdge] = false
		}
	    }
	}
    }

    count := 0
    for e := range edges {
	if edges[e] {
	    count ++
	}
    }

    return count
}

func part2(_map [][]string) int {
    count := 0
    groups := getGroups(_map)
    for _, g := range groups {
	count += len(g)*getSides(_map, g)
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
    _map := make([][]string, len(lines))
    for i, line := range lines {
	_map[i] = strings.Split(line, "")
    }

    fmt.Println("part 1:", part1(_map))
    fmt.Println("part 2:", part2(_map))

}
