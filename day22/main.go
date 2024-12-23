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

func evolve (num int) int {
    temp := 64*num
    num ^= temp
    num = num % 16777216

    temp = num/32
    num ^= temp
    num = num % 16777216

    temp = num*2048
    num ^= temp
    num = num % 16777216

    return num
}

func part1(lines []string) int{
    count := 0
    for _, line := range lines {
	num, err := strconv.Atoi(line) 
	check(err)

	for i:=0; i < 2000; i++ {
	    num = evolve(num)
	}
	count += num
    }
    return count
}

func compare (a1, a2 []int) bool {
    if len(a1) != len(a2) {
	return false
    }
    for i, n := range a1 {
	if n != a2[i] {
	    return false
	}
    }
    return true
}

func getPrices(lines []string, sellSequence []int) int {
    count := 0
    for _, line := range lines {
	num, err := strconv.Atoi(line) 
	check(err)
	lastPrice := num % 10
	sequence := make([]int, 0)

	for i:=0; i < 2000; i++ {
	    num = evolve(num)
	    price := num % 10
	    diff := price - lastPrice

	    sequence = append(sequence, diff)
	    if len(sequence) > 4 {
		sequence = sequence[1:]
	    }

	    if compare(sequence, sellSequence) {
		count += price
		break
	    }

	    lastPrice = price
	}
    }
    return count
}

func keyWithMaxValue(m map[string]int) (string, int) {
	if len(m) == 0 {
		return "", 0
	}

	var maxKey string
	var maxValue int
	first := true

	for k, v := range m {
		if first || v > maxValue {
			maxKey = k
			maxValue = v
			first = false
		}
	}
	return maxKey, maxValue
}

func seqToString(seq []int) string {
    str := ""
    for _, num := range seq {
	numString := strconv.Itoa(num)
	str += numString + ","
    }
    return str
}

func part2(lines []string) int {
    prices := make([][]int, len(lines))
    diffs := make([][]int, len(lines))
    for i, line := range lines {
	num, err := strconv.Atoi(line)
	check(err)
	monkeyPrices := make([]int, 2000)
	monkeyPrices[0] = num % 10
	diffsPerMonkey := make([]int, 1999)
	for j := 1; j < 2000; j++ {
	    num = evolve(num)
	    monkeyPrices[j] = num % 10
	    diffsPerMonkey[j-1] = monkeyPrices[j] - monkeyPrices[j-1]
	}
	prices[i] = monkeyPrices
	diffs[i] = diffsPerMonkey
    }

    bananasPerSequence := make(map[string]int)
    for m, diffsPerMonkey := range diffs {
	seqSeen := make(map[string]bool)
	for i := 0; i < len(diffsPerMonkey)-3; i++ {
	    seq := diffsPerMonkey[i:i+4]
	    stringSeq := seqToString(seq)
	    if !seqSeen[stringSeq] {
		seqSeen[stringSeq] = true
		if _, exists := bananasPerSequence[stringSeq]; !exists {
		    bananasPerSequence[stringSeq] = 0
		}
		bananasPerSequence[stringSeq] += prices[m][i+4]
	    }
	}
    }
    
    k, v := keyWithMaxValue(bananasPerSequence)
    fmt.Println(k)

    return v


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
