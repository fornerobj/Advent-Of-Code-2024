package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check (e error) {
    if e != nil {
        panic(e)
    }
}

func isCorrect(adjList map[int][]int, nums []int) bool {
    goodLine := true
    for j, n := range(nums) {
	badNums := adjList[n]
	for k:=j+1; k < len(nums); k++ {
	    if slices.Contains(badNums, nums[k]) {
		goodLine = false	
		break
	    }
	}
    }
    return goodLine
}

func createAdjList(lines []string) (map[int][]int, int) {
    adjList := make(map[int][]int)
    i := 0

    for ; lines[i] != ""; i++ {
	nums := strings.Split(lines[i], "|")	
	num1, err := strconv.Atoi(nums[0])
	check(err)
	num2, err := strconv.Atoi(nums[1])
	check(err)

	if _, ok := adjList[num2]; !ok {
	    adjList[num2] = []int{}
	}
	adjList[num2] = append(adjList[num2], num1)
    }

    return adjList, i
}

func part1(lines []string) int {
    count := 0
    adjList, i := createAdjList(lines)

    for ; i < len(lines); i++ {
	line := lines[i]

	if(line == "") {
	    continue
	}

	var nums []int
	for _, numStr := range strings.Split(line, ",") {
	    num, err := strconv.Atoi(numStr)
	    check(err)
	    nums = append(nums, num)
	}

	goodLine := isCorrect(adjList, nums)
	if(goodLine) {
	    count += nums[len(nums)/2]
	}
    }

    return count
}

func topological_sort(graph map[int][]int) []int {
    visited := make(map[int]bool)
    stack := []int{}
    var dfs func(node int)

    dfs = func(node int) {
	visited[node] = true

	for _, neighbor := range graph[node] {
	    if !visited[neighbor] {
		dfs(neighbor)
	    }
	}

	stack = append(stack, node)
    }

    //have to do for all nodes, becuase graph may be disconnected
    for node := range graph {
	if !visited[node] {
	    dfs(node)
	}
    }

    //usually have to reverse the stack after DFS for topological sort, 
    //but the edges in my graph are already backwards
    return stack
}

func part2(lines []string) int {
    count := 0
    adjList, i := createAdjList(lines)

    for ; i < len(lines); i++ {
	line := lines[i]

	if(line == "") {
	    continue
	}

	var nums []int
	for _, numStr := range strings.Split(line, ",") {
	    num, err := strconv.Atoi(numStr)
	    check(err)
	    nums = append(nums, num)
	}

	goodLine := isCorrect(adjList, nums)

	if !goodLine {
	    newAdjList := make(map[int][]int)
	    for _, num := range nums {
		if _, ok := newAdjList[num]; !ok {
		    newAdjList[num] = []int{}
		}
		if _, ok := adjList[num]; !ok {
		    continue
		}
		for _, neighbor := range adjList[num] {
		    if slices.Contains(nums, neighbor) {
			newAdjList[num] = append(newAdjList[num], neighbor)
		    }
		}
	    }
	    newLine := topological_sort(newAdjList)
	    count += newLine[len(newLine)/2]
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
 //    for _, line := range lines {
	// fmt.Println(line)
 //    }

    fmt.Println("part 1:", part1(lines))
    fmt.Println("part 2:", part2(lines))

}
