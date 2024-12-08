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


func good(s []int, cur, target int, ops string, part2 bool) bool {
    if len(s) == 0 {
	if cur == target {
	    // fmt.Printf("%d = %s\n",target, ops)
	    return true
	}else {
	    return false
	}
    }

    plus := good(s[1:], cur + s[0], target, ops + " + " + strconv.Itoa(s[0]), part2)
    if plus {
	return true
    }
    times := good(s[1:], cur * s[0], target, ops + " * " + strconv.Itoa(s[0]), part2)
    if times {
	return times
    }
    if part2 {
	cc, err := strconv.Atoi(strconv.Itoa(cur)+strconv.Itoa(s[0]))
	check(err)
	concat := good(s[1:], cc, target, ops + " || " + strconv.Itoa(s[0]), part2)
	if concat {
	    return concat
	}
    }
    return false
}

func part1(lines []string) int {
    count := 0
    for _,line := range lines {
	split := strings.Split(line, ":")
	target, err := strconv.Atoi(split[0])
	check(err)
	nums_strings := split[1]
	nums := make([]int, 0)
	for _, n := range strings.Split(nums_strings, " ") {
	    if n =="" {
		continue
	    }
	    ni, err := strconv.Atoi(n)
	    check(err)
	    nums = append(nums, ni)
	}
	if good(nums[1:], nums[0], target, strconv.Itoa(nums[0]), false) {
            count += target
        }
    }
    return count
}



func part2(lines []string) int {
    count := 0
    for _,line := range lines {
	split := strings.Split(line, ":")
	target, err := strconv.Atoi(split[0])
	check(err)
	nums_strings := split[1]
	nums := make([]int, 0)
	for _, n := range strings.Split(nums_strings, " ") {
	    if n =="" {
		continue
	    }
	    ni, err := strconv.Atoi(n)
	    check(err)
	    nums = append(nums, ni)
	}
	if good(nums[1:], nums[0], target, strconv.Itoa(nums[0]), true) {
            count += target
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

    fmt.Println("part 1:", part1(lines))
    fmt.Println("part 2:", part2(lines))

}
