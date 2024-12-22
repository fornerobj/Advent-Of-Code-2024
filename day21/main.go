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
    {1, 0}, 
    {0, -1}, 
    {0, 1}, 
    {-1, 0},
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

type state struct {
    dp1 Pos
    dp2 Pos
    dp3 Pos
    np  Pos
}

var arrowToMap = map[Pos]string{
    {0, 1} :">", 
    {0, -1}:"<", 
    {-1, 0}:"^", 
    {1, 0} :"v",
}

func sequenceToButton(pad [][]string, curPos Pos, target string) (string, Pos){
    w, h := len(pad[0]), len(pad)
    q := []Node{Node{curPos, 0, ""}}
    seen := make(map[Pos]bool)

    for len(q) > 0 {
	cur := q[0]
	q = q[1:]
	if seen[cur.pos] {
	    continue
	}
	seen[cur.pos] = true

	if cur.pos.R < 0 || cur.pos.R >= h || cur.pos.C < 0 || cur.pos.C >= w || pad[cur.pos.R][cur.pos.C] == ""{
	    continue
	}

	if pad[cur.pos.R][cur.pos.C] == target {
	    return cur.sequence+"A", cur.pos
	}

	for _, dir := range directions {
	    newPos := Pos{cur.pos.R + dir.R, cur.pos.C + dir.C}
	    sequenceCharacter := arrowToMap[dir]
	    q = append(q, Node{newPos, cur.cost+1, cur.sequence+sequenceCharacter})
	}
	
    }
    return "NOT FOUND", Pos{-1,-1}
}

func part1(codes []string) int {
    count := 0
    r1 := Pos{0,2}
    r2 := Pos{0,2}
    r3 := Pos{3,2}

    
    for _, code := range codes {
	fmt.Println(code)
	numberSequence := ""
	for _, digit := range strings.Split(code,"") {
	    newPiece, newPos := sequenceToButton(numpad, r3, digit)
	    numberSequence += newPiece
	    r3 = newPos
	}

	directionSequence1 := ""
	for _, digit2 := range strings.Split(numberSequence, "") {
	    newPiece, newPos := sequenceToButton(dPad, r2, digit2)
	    directionSequence1 += newPiece
	    r2 = newPos
	}

	directionSequence2 := ""
	for _, digit3 := range strings.Split(directionSequence1, "") {
	    newPiece, newPos := sequenceToButton(dPad, r1, digit3)
	    directionSequence2 += newPiece
	    r1 = newPos
	}
	n := len(directionSequence2)
	numericPart, err := strconv.Atoi(code[:3])
	check(err)
	fmt.Println(n, numericPart)
	count += numericPart*n
    }
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
    seq, _ := sequenceToButton(dPad, Pos{1, 0}, "A")
    fmt.Println(seq)
    fmt.Println("part 2:", part2())

}
