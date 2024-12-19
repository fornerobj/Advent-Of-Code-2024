package main

import (
	"aoc2024/utils"
	"container/heap"
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
    X int
    Y int
}

var directions []Pos = []Pos {
    {0,1},
    {0,-1},
    {1,0},
    {-1,0},
}

func sim(bytes []Pos, numBytes int) bool {
    start := Pos{0,0}
    finish := Pos{70,70}
    w, h := 71, 71

    corrupted := make(map[Pos]bool)

    for i, byte := range bytes {
	if i >= numBytes {
	    break
	}
	corrupted[byte] = true
    }

    pq := &utils.PriorityQueue{}

    heap.Init(pq)
    heap.Push(pq, &utils.Item{
	Value: start,
	Priority: 0,
    })

    prev := make(map[Pos]Pos)
    costs := make(map[Pos]int)
    costs[start] = 0

    for pq.Len() > 0 {
	curItem := heap.Pop(pq).(*utils.Item)
	cur := curItem.Value.(Pos)
	curCost := curItem.Priority

	if curCost > costs[cur] {
	    continue
	}
	

	for _, dir := range directions {
	    dx, dy := dir.X, dir.Y
	    nx, ny := cur.X+dx, cur.Y+dy 

	    if nx < 0 || nx >= w || ny < 0 || ny >= h {
		continue
	    }
	    if corrupted[Pos{nx,ny}] {
		continue
	    }

	    next := Pos{nx, ny}
	    newCost := curCost + 1

	    if oldCost, exists := costs[next]; !exists || newCost < oldCost {
		costs[next] = newCost
		prev[next] = cur
		heap.Push(pq, &utils.Item {
		    Value: next, 
		    Priority: newCost,
		})
	    }
	}
    }
    _, exists := costs[finish]
    return exists
}


func part1(bytes []Pos) int{
    start := Pos{0,0}

    // finish := Pos{6,6}
    // w, h := 7, 7
    // numBytes := 12

    finish := Pos{70,70}
    w, h := 71, 71
    numBytes := 1024

    corrupted := make(map[Pos]bool)

    for i, byte := range bytes {
	if i >= numBytes {
	    break
	}
	corrupted[byte] = true
    }

    pq := &utils.PriorityQueue{}

    heap.Init(pq)
    heap.Push(pq, &utils.Item{
	Value: start,
	Priority: 0,
    })

    prev := make(map[Pos]Pos)
    costs := make(map[Pos]int)
    costs[start] = 0

    for pq.Len() > 0 {
	curItem := heap.Pop(pq).(*utils.Item)
	cur := curItem.Value.(Pos)
	curCost := curItem.Priority

	if curCost > costs[cur] {
	    continue
	}
	

	for _, dir := range directions {
	    dx, dy := dir.X, dir.Y
	    nx, ny := cur.X+dx, cur.Y+dy 

	    if nx < 0 || nx >= w || ny < 0 || ny >= h {
		continue
	    }
	    if corrupted[Pos{nx,ny}] {
		continue
	    }

	    next := Pos{nx, ny}
	    newCost := curCost + 1

	    if oldCost, exists := costs[next]; !exists || newCost < oldCost {
		costs[next] = newCost
		prev[next] = cur
		heap.Push(pq, &utils.Item {
		    Value: next, 
		    Priority: newCost,
		})
	    }
	}
    }
    return costs[finish]
}

func part2(bytes []Pos) Pos {
    for i := 0; i < len(bytes); i++ {
	pathExists := sim(bytes, i)
	if !pathExists {
	    fmt.Println("Found at i=", i)
	    return bytes[i-1]
	}
    }
    return Pos{-1,-1}
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    bytes := make([]Pos, 0)
    for _, line := range lines {
	nums := strings.Split(line, ",")
	n1, err := strconv.Atoi(nums[0])
	check(err)
	n2, err := strconv.Atoi(nums[1])
	check(err)
	bytes = append(bytes, Pos{n1, n2})
    }

    fmt.Println("part 1:", part1(bytes))
    pos := part2(bytes)
    fmt.Println("part 2:", pos.X, ",", pos.Y)

}
