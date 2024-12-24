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

func convert(char string, wires map[string]bool) int{
    num := 0
    for i := 0;; i++ {
	istring := strconv.Itoa(i)
	var target string
	if i < 10 {
	    target = char + "0" + istring
	}else {
	    target = char + istring
	}
	b, exists := wires[target]
	if !exists {
	    break
	}
	if b {
	    num += int(math.Pow(2,float64(i)))
	}
    }
    return num
}

func getSignals(signals map[string]bool, gates []string) map[string]bool {
    gatesCompleted := make(map[int]bool)
    for len(gatesCompleted) < len(gates) {
	for i, gate := range gates {
	    if gatesCompleted[i] {
		continue
	    }
	    split := strings.Split(gate, " -> ")
	    split2 := strings.Split(split[0], " ")
	    w1 := split2[0]
	    operator := split2[1]
	    w2 := split2[2]
	    output := split[1]

	    b1, e1 := signals[w1]
	    b2, e2 := signals[w2]
	    if !e1 || !e2 {
		continue
	    }

	    if operator == "AND" {
		signals[output] = b1 && b2
	    }else if operator == "OR" {
		signals[output] = b1 || b2
	    }else {
		signals[output] = b1 != b2
	    }

	    gatesCompleted[i] = true
	}
    }

    return signals
}

func part1(signals map[string]bool, gates []string) int{
    signals = getSignals(signals, gates)
    return convert("z",signals)
}

func part2(signals map[string]bool, gates []string) int {
    signals = getSignals(signals, gates)
    // x := convert("x", signals)
    // z := convert("z", signals)
    // y := convert("y", signals)
    // fmt.Println(x)
    // fmt.Println(y)
    // fmt.Println(x+y)
    // fmt.Println(z)
    return 0
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)
    split := strings.Split(input, "\n\n")
    inits := strings.Split(split[0], "\n")
    gates := strings.Split(split[1], "\n")
    gates = gates[:len(gates)-1]

    signals := make(map[string]bool)
    for _, init := range inits {
	split := strings.Split(init, ": ")
	if split[1] == "0" {
	    signals[split[0]] = false
	}else {
	    signals[split[0]] = true
	}
    }

    fmt.Println("part 1:", part1(signals, gates))
    fmt.Println("part 2:", part2(signals, gates))

}
