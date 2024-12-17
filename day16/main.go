package main

import (
	"aoc2024/utils"
	"container/heap"
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

func printMap(grid [][]string, seats map[Robot]bool) {
    for seat := range seats {
	if seats[seat] {
	    grid[seat.pos.R][seat.pos.C] = "O"
	}
    }
    for _, row := range grid {
	fmt.Println(row)
    }
}

var dirToDirection map[int]Pos = map[int]Pos{
    0: {-1, 0},
    1: {0, 1},
    2: {1, 0},
    3: {0, -1},
}

func findPath(grid [][]string) (int, int) {
    start := Robot{getStart(grid), 1}
    finish := getFinish(grid)

    pq := &utils.PriorityQueue{}

    heap.Init(pq)
    heap.Push(pq, &utils.Item{
	Value: start,
	Priority: 0,
    })

    prev := make(map[Robot][]Robot)
    costs := make(map[Robot]int)
    costs[start] = 0

    for pq.Len() > 0 {
	curItem := heap.Pop(pq).(*utils.Item)
	cur := curItem.Value.(Robot)
	curCost := curItem.Priority

	if curCost > costs[cur] {
	    continue
	}

	for i := -1; i <= 1; i++ {
	    newDir := (cur.Dir + i + 4) % 4
	    newPos := cur.pos
	    moveCost := 0
	    
	    if i == 0 {
		newPos = Pos{cur.pos.R + dirToDirection[newDir].R, cur.pos.C + dirToDirection[newDir].C}
		moveCost = 1
	    } else {
		moveCost = 1000
	    }

	    if grid[newPos.R][newPos.C] == "#" {
		continue
	    }

	    next := Robot{newPos, newDir}
	    newCost := curCost + moveCost

	    if oldCost, exists := costs[next]; !exists || newCost < oldCost {
		costs[next] = newCost
		prev[next] = []Robot{cur}
		heap.Push(pq, &utils.Item {
		    Value: next, 
		    Priority: newCost,
		})
	    }else if oldCost == newCost {
		prev[next] = append(prev[next], cur)
	    }
	}
    }

    minCost := math.MaxInt64
    var finalRobot Robot
    for d := 0; d < 4; d++ {
	r := Robot{finish, d}
	if cost, exists := costs[r]; exists && cost < minCost {
	    minCost = cost
	    finalRobot = r
	}
    }


    stack := []Robot{finalRobot}
    pathSet := map[Robot]bool {finalRobot:true}

    for len(stack) > 0 {
	temp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	for _, other := range prev[temp] {
	    if !pathSet[other] {
		pathSet[other] = true
		stack = append(stack, other)
	    }
	}
    }

    uniquePositions := make(map[Pos]bool)
    for r := range pathSet {
	if !uniquePositions[r.pos] {
	    uniquePositions[r.pos] = true
	}
    }
    
    return minCost, len(uniquePositions)
}

func part1(grid [][]string) int {
    minCost,_ := findPath(grid)
    return minCost
}

func part2(grid [][]string) int {
    _,numSeats := findPath(grid)
    return numSeats
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
