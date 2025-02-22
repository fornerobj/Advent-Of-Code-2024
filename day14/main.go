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

type Robot struct {
    px int
    py int
    vx int
    vy int
}

type Pos struct {
    X int
    Y int
}

func move(bots []Robot, t, w, h int) []Robot {
    for i := range bots {
	nx := (bots[i].px + (bots[i].vx*t)) % w
	if nx < 0 {
	    nx = w+nx
	}
	ny := (bots[i].py + (bots[i].vy*t)) % h
	if ny < 0 {
	    ny = h+ny
	}
	bots[i].px = nx
	bots[i].py = ny
    }
    return bots
}

func printMap(botMap map[Pos]bool, w, h int) {
    for row := 0; row < h; row ++ {
	for col := 0; col < w; col ++ {
	    if(botMap[Pos{col,row}]) {
		fmt.Print("#")
	    }else {
		fmt.Print(".")
	    }
	}
	fmt.Print("\n")
    }
}

func checkForTree(bots []Robot, w, h int) bool {
    _map := make(map[Pos]bool) 

    for _, b := range bots {
	_map[Pos{b.px, b.py}] = true
    }

    //My idea is to create a 3x3 kernel to sweep through the coordinate grid
    //similar to what is done in CNNs. I suspect that any shape created would 
    //need at least 1 3x3 block to be filled in.

    for row := 1; row < h-1; row++ {
	for col := 1; col < w-1; col++ {
	    if _map[Pos{col-1,row-1}] && _map[Pos{col,row-1}] && _map[Pos{col+1,row-1}] &&
		_map[Pos{col-1,row}] && _map[Pos{col,row}] && _map[Pos{col+1,row}] &&
		_map[Pos{col-1,row+1}] && _map[Pos{col,row+1}] && _map[Pos{col+1,row+1}] {
		    printMap(_map, w, h) 
		    return true
		}
	}
    }
    return false
}

func part1(bots []Robot) int {
    q1, q2, q3, q4 := 0, 0, 0, 0
    w, h := 101, 103

    bots = move(bots, 100, w, h)
    for _, b := range bots {
	x, y := b.px, b.py 
	if x < w/2 && y < h/2 {
	    q1++
	}else if x > w/2 && y < h/2 {
	    q2 ++
	}else if x > w/2 && y > h/2 {
	    q3 ++
	}else if x < w/2 && y > h/2 {
	    q4 ++
	}
    }
    return q1*q2*q3*q4
}

func part2(bots []Robot) int {
    w, h := 101, 103
    for i := 0; i < 100000; i++ {
	if checkForTree(bots, w, h) {
	    return i
	}
	bots = move(bots, 1, w, h)
    }
    return -1
}

func parse (line string) (int, int, int, int) {
	pandv := strings.Split(line, " ")
	pxandpy := strings.Split(pandv[0], ",")
	px, err := strconv.Atoi(strings.TrimPrefix(pxandpy[0], "p="))
	check(err)
	py, err := strconv.Atoi(pxandpy[1])
	check(err)

	vxandvy := strings.Split(pandv[1], ",")
	vx, err := strconv.Atoi(strings.TrimPrefix(vxandvy[0], "v="))
	check(err)
	vy, err := strconv.Atoi(vxandvy[1])
	check(err)

	return px, py, vx, vy
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]
    bots := make([]Robot, 0)
    for _, l := range lines {
	px, py, vx, vy := parse(l)
	bots = append(bots, Robot{px,py,vx,vy})
    }

    fmt.Println("part 1:", part1(bots))
    bots = make([]Robot, 0)
    for _, l := range lines {
	px, py, vx, vy := parse(l)
	bots = append(bots, Robot{px,py,vx,vy})
    }
    fmt.Println("part 2:", part2(bots))

}
