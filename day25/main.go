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

func part1(locks, keys [][]int) int {
    n := len(locks[0])
    count := 0
    for _, l := range locks {
	for _, k := range keys {
	    fit := true
	    for i := 0; i < n; i++ {
		if l[i]+k[i] > 5 {
		    fit = false
		    break
		}
	    }
	    if fit {
		count++
	    }
	}
    }
    return count
}

func part2() int {
    return 0
}

func convert(lockOrKey []string) []int {
    pinHeights := []int{0,0,0,0,0}
    for r, row := range lockOrKey {
	if r == 0 || r == len(lockOrKey) -1 {
	    continue
	}
	for i, c := range strings.Split(row, "") {
	    if c == "#" {
		pinHeights[i]++
	    }
	}
    }
    return pinHeights
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := strings.TrimSpace(string(data))
    locksAndKeys := strings.Split(input, "\n\n")
    n := len(locksAndKeys)

    locks := make([][]int, 0)
    keys := make([][]int, 0)

    for i:=0; i < n; i++ {
	lockOrKey := strings.Split(locksAndKeys[i], "\n")
	lock := true
	for _, c := range strings.Split(lockOrKey[0], "") {
	    if c != "#" {
		lock = false
	    }
	}
	pinHeights := convert(lockOrKey)
	if lock {
	    locks = append(locks, pinHeights)
	} else {
	    keys = append(keys, pinHeights)
	}
    }

    fmt.Println("part 1:", part1(locks, keys))
    fmt.Println("part 2:", part2())

}
