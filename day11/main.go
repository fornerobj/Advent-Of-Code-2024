package main

import (
	"fmt"
	"math"
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
    num int
    t int
}

type Node struct {
    num int
    next *Node
}

func NewNode(num int) *Node {
    return &Node{num: num, next: nil}
}

type LinkedList struct {
    head *Node
}

func getSize(l LinkedList) int {
    cur := l.head
    size := 0
    for ; cur != nil; cur = cur.next {
	size ++
    }
    return size
}

func printList(l LinkedList) {
    fmt.Print("[ ")
    for cur := l.head; cur != nil; cur = cur.next {
	fmt.Print(cur.num, " ")
    }
    fmt.Print("]")
    fmt.Println()
}

func score(num, numBlinks int, memo map[Pos]int) int {
    pos := Pos{num, numBlinks}
    _, ok := memo[pos]
    if ok {
	return memo[pos]
    }

    if numBlinks == 0 {
	return 1
    }

    if num == 0 {
	memo[pos] = score(1, numBlinks-1, memo)
	return memo[pos]
    }

    numDigits := int(math.Floor(math.Log10(float64(num)))) + 1
    if numDigits % 2 == 0 {
	num1 := num / int(math.Pow(10, float64(numDigits/2)))
	num2 := num % int(math.Pow(10, float64(numDigits/2)))
	memo[pos] = score(num1, numBlinks-1, memo) + score(num2, numBlinks-1, memo)
	return memo[pos]
	
    }
    
    memo[pos] = score(num*2024, numBlinks-1, memo)
    return memo[pos]
}

func blink(stones LinkedList, numBlinks int) int {
    memo := make(map[Pos]int)
    count := 0
    for cur := stones.head; cur != nil; cur = cur.next {
	count += score(cur.num, numBlinks, memo)	
    }
    return count
}


func part1(stones LinkedList) int {
    return blink(stones, 25)
}

func part2(stones LinkedList) int {
    return blink(stones, 75)
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := strings.TrimSpace(string(data))
    numsString := strings.Split(input, " ")
    nums := make([]int, len(numsString))
    for i, n := range numsString {
	nums[i], err = strconv.Atoi(n)
	check(err)
    }

    var stones LinkedList
    var prev *Node = nil
    for i, n := range nums {
	nn := NewNode(n)
	if i == 0 {
	    stones.head = nn
	    prev = stones.head
	}else {
	    prev.next = nn
	    prev = prev.next
	}
    }

    fmt.Println("part 1:", part1(stones))
    fmt.Println("part 2:", part2(stones))

}
