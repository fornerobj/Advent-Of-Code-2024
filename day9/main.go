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


func part1(input string) int {
    count := 0
    nums := make([]int, len(input))

    for i, c := range input {
	n, err := strconv.Atoi(string(c))
	check(err)
	nums[i] = n
    }

    expanded := make([]int,0)
    
    j := 0
    for i, n := range nums {
	if i % 2 == 0 {
	    for k := 0; k < n; k++ {
		expanded = append(expanded, j)
	    }
	    j++
	}else {
	    for k:= 0; k < n; k++ {
		expanded = append(expanded, -1)
	    }
	}
    }
    for i, n := range expanded {
	if n == -1 {
	    for j := len(expanded)-1; j > i; j-- {
		if expanded[j] != -1 {
		    temp := n
		    expanded[i] = expanded[j]
		    expanded[j] = temp
		    break
		}
	    }
	}
    }

    for i, n := range expanded {
	if n == -1 {
	    break
	}
	count += i*n
    }
    return count
}

func part2(input string) int {
    count := 0
    nums := make([]int, len(input))

    for i, c := range input {
	n, err := strconv.Atoi(string(c))
	check(err)
	nums[i] = n
    }

    expanded := make([]int,0)
    
    id := 0
    for i, n := range nums {
	if i % 2 == 0 {
	    for k := 0; k < n; k++ {
		expanded = append(expanded, id)
	    }
	    id++
	}else {
	    for k:= 0; k < n; k++ {
		expanded = append(expanded, -1)
	    }
	}
    }
    id --

    var s1, s2 int

    for ; id >= 0; id-- {
	length := 0;
	started := false
	for j := len(expanded)-1; j >= 0; j-- {
	    if expanded[j] == id {
		if !started {
		    started = true
		}
		s1 = j
		length ++
	    }
	}


	emptylength := 0
	counting := false
	for j := 0; j < s1; j++ {
	    if expanded[j] == -1 {
		if !counting {
		    s2 = j
		}
		counting = true
		emptylength++
		if emptylength == length {
		    for k := 0; k < length; k++ {
			temp := expanded[s1]
			expanded[s1] = expanded[s2]
			expanded[s2] = temp
			s1 ++
			s2 ++
		    }
		    break
		}
	    }else {
		if counting {
		    counting = false
		    emptylength = 0
		}
	    }    
	}
    }

    for i, n := range expanded {
	if n == -1 {
	    continue
	}
	count += i*n
    }
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := strings.TrimSpace(string(data))

    fmt.Println("part 1:", part1(input))
    fmt.Println("part 2:", part2(input))

}
