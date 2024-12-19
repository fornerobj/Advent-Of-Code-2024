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

var memoNums map[string]int = make(map[string]int)

func goodCombos(towels []string, target string) int {
    if numWays, exists := memoNums[target]; exists {
	return numWays
    }
    if len(target) == 0 {
	memoNums[target] = 1
	return 1
    }

    numWays := 0
    for _, t := range towels {
	n := len(t)
	if n > len(target) {
	    continue
	}
	if t == target[:n] {
	    numWays += goodCombos(towels, target[n:]) 
	}
    }

    memoNums[target] = numWays
    return numWays
}

func part1(towels, targets []string) int{
    count := 0 
    for _, t := range targets {
	if goodCombos(towels, t) > 0 {
	    count ++
	}
    }
    return count
}

func part2(towels, targets []string) int {
    count := 0 
    for _, t := range targets {
	count += goodCombos(towels, t)
    }
    return count

}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n\n")

    towelsString := lines[0]
    towels := make([]string, 0)
    for _, t := range strings.Split(towelsString, ",") {
	towels = append(towels, strings.TrimSpace(t))
    }

    targets := strings.Split(lines[1], "\n")
    targets = targets[:len(targets)-1]

    fmt.Println("part 1:", part1(towels, targets))
    fmt.Println("part 2:", part2(towels, targets))

}
