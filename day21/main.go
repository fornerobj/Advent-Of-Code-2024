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

type Pos struct {
    R int
    C int
}

type Node struct {
    pos Pos
    cost int 
    sequence string
}

var directions []Pos = []Pos {
    {-1, 0}, //up
    {0, 1}, //right
    {1, 0}, //down
    {0, -1}, //left
}

var numpad [][]string = [][]string {
    {"7", "8", "9"},
    {"4", "5", "6"},
    {"1", "2", "3"},
    {"", "0", "A"},
}

var dPad [][]string = [][]string {
    {"", "^", "A"}, 
    {"<","v", ">"},
}

var arrowToMap = map[Pos]string{
    {0, 1} :">", 
    {0, -1}:"<", 
    {-1, 0}:"^", 
    {1, 0} :"v",
}

func part1(codes []string) int {
    count := 0
    return count
}

func part2() int {
    count := 0
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    lines := strings.Split(input, "\n")
    codes := lines[:len(lines)-1]

    fmt.Println("part 1:", part1(codes))
    fmt.Println("part 2:", part2())

}
