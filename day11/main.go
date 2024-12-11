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

func blink(stones LinkedList) {
    for cur := stones.head; cur != nil; cur = cur.next {
	if cur.num == 0 {
	    cur.num = 1
	}else if (int(math.Floor(math.Log10(float64(cur.num)))) + 1)%2 == 0 {
	    n := (int(math.Log10(float64(cur.num))) + 1) / 2
	    temp := cur.num
	    cur.num = cur.num / int(math.Pow(10, float64(n)))
	    nn := NewNode(temp % int(math.Pow(10, float64(n))))
	    nn.next = cur.next
	    cur.next = nn
	    cur = cur.next
	}else {
	    cur.num = cur.num * 2024
	}
    }
}


func part1(stones LinkedList) int {
    count := 0
    for cur := stones.head; cur != nil; cur = cur.next {
	var newList LinkedList
	newList.head = NewNode(cur.num)
	for i:=0; i < 25; i++ {
	    blink(newList)
	}
	count += getSize(newList)

    }
    return count
}

func part2(stones LinkedList) int {
    count := 0
    for cur := stones.head; cur != nil; cur = cur.next {
	var newList LinkedList
	newList.head = NewNode(cur.num)
	for i:=0; i < 75; i++ {
	    blink(newList)
	}
	count += getSize(newList)

    }
    return count


    return getSize(stones)
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
