package main

import (
	"errors"
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
    X int
    Y int
}

func solveSystem (a1,b1,c1,a2,b2,c2 int) (int,int, error){
    det := a1*b2 - a2*b1 
    if det == 0 {
	return -1,-1, errors.New("No solution")
    }
    
    Adet := (b2*c1 - b1*c2)
    Bdet := (a1*c2 - a2*c1)
    if Adet%det != 0 || Bdet%det != 0 {
	return -1, -1, errors.New("no whole number solution")
    }
    
    A := Adet / det
    B := Bdet / det
    return A,B, nil

}
func parse(input string) (int, int, int, int, int, int) {
    lines := strings.Split(input, "\n")
    if len(lines) != 3 {
	lines = lines[:3]
    }
    
    buttonA := strings.TrimPrefix(lines[0], "Button A: ")
    partsA := strings.Split(buttonA, ", ")
    aX, err := strconv.Atoi(strings.TrimPrefix(partsA[0], "X+"))
    check(err)
    aY, err := strconv.Atoi(strings.TrimPrefix(partsA[1], "Y+"))
    check(err)

    buttonB := strings.TrimPrefix(lines[1], "Button B: ")
    partsB := strings.Split(buttonB, ", ")
    bX, err := strconv.Atoi(strings.TrimPrefix(partsB[0], "X+"))
    check(err)
    bY, err := strconv.Atoi(strings.TrimPrefix(partsB[1], "Y+"))
    check(err)

    prize := strings.TrimPrefix(lines[2], "Prize: ")
    partsPrize := strings.Split(prize, ", ")
    prizeX, err := strconv.Atoi(strings.TrimPrefix(partsPrize[0], "X="))
    check(err)
    prizeY, err := strconv.Atoi(strings.TrimPrefix(partsPrize[1], "Y="))
    check(err)

    return aX, aY, bX, bY, prizeX, prizeY
}
func part1(machines []string) int {
    count := 0
    for _, m := range machines {
	aX, aY, bX, bY, prizeX, prizeY := parse(m)
	A, B, err := solveSystem(aX, bX, prizeX, aY, bY, prizeY)
	if err == nil {
	    count += 3*A + B
	}
    }
    return count
}

func part2(machines []string) int {
    count := 0
    for _, m := range machines {
	aX, aY, bX, bY, prizeX, prizeY := parse(m)
	prizeX += 10000000000000
	prizeY += 10000000000000
	A, B, err := solveSystem(aX, bX, prizeX, aY, bY, prizeY)
	if(err == nil) {
	    count += 3*A + B
	}
    }
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    machines := strings.Split(input, "\n\n")

    fmt.Println("part 1:", part1(machines))
    fmt.Println("part 2:", part2(machines))

}
