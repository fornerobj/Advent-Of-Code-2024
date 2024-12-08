package main

import (
	"fmt"
	"os"
	"strings"
	"math"
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

func distance(x1, y1, x2, y2 int) float64 {
    return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func getSlope(x1, y1, x2, y2 int) float64 {
    return float64(y2-y1)/float64(x2-x1)
}

//y = m*x + b
func getYIntercept(x, y int, m float64) float64 {
    return float64(y)-(m*float64(x))
}

func getY(x int, m, b float64) float64 {
    return (m*float64(x)) + b 
}

func part1(lines []string) int {
    width := len(lines[0])
    height := len(lines)
    antinodes := make(map[Pos]bool)
    positions := make(map[string][]Pos)
    for row, line := range lines {
	for col, c := range line {
	    character := string(c)
	    if character != "." {
		positions[character] = append(positions[character], Pos{col, row})
	    }
	}
    }
    for character := range positions {
	for i, pos1 := range positions[character] {
	    for j, pos2 := range positions[character] {
		if i == j {
		    continue
		}
		dx := pos2.X - pos1.X
		dy := pos2.Y - pos1.Y
		newX := 2*dx+pos1.X
		newY := 2*dy+pos1.Y
		if newX >= 0 && newX < width && newY >= 0 && newY < height {
		    _, ok := antinodes[Pos{newX,newY}]
		    if !ok {
			antinodes[Pos{newX, newY}] = true
		    }
		}
	    } 
	}
    }
    return len(antinodes)
}



func part2(lines []string) int {
    width := len(lines[0])
    height := len(lines)
    antinodes := make(map[Pos]bool)
    positions := make(map[string][]Pos)

    for row, line := range lines {
	for col, c := range line {
	    character := string(c)
	    if character != "." {
		positions[character] = append(positions[character], Pos{col, row})
	    }
	}
    }

    for character := range positions {
	for i, pos1 := range positions[character] {
	    for j, pos2 := range positions[character] {
		if i == j {
		    continue
		}
		dx := pos2.X-pos1.X
		dy := pos2.Y-pos1.Y

		for k := -width; k <= width; k++ { // Range sufficiently large for both directions
		    newX := pos1.X + k*dx
		    newY := pos1.Y + k*dy

		    if newX < 0 || newX >= width || newY < 0 || newY >= height {
			continue
		    }
		    _, ok := antinodes[Pos{newX, newY}]
		    if !ok {
			antinodes[Pos{newX, newY}] = true
		    }
		}

	    } 
	}
    }
    return len(antinodes)
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
