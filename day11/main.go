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

    numDigits := int(math.Floor(math.Log10(float64(num)))) + 1 //floor of log base 10 of num + 1 gives length
    if numDigits % 2 == 0 {
	num1 := num / int(math.Pow(10, float64(numDigits/2)))
	num2 := num % int(math.Pow(10, float64(numDigits/2)))
	memo[pos] = score(num1, numBlinks-1, memo) + score(num2, numBlinks-1, memo)
	return memo[pos]
    }
    
    memo[pos] = score(num*2024, numBlinks-1, memo)
    return memo[pos]
}

func part1(stones []int, memo map[Pos]int) int {
    count := 0
    for _, num := range stones {
	count += score(num, 25, memo)	
    }
    return count
}

func part2(stones []int, memo map[Pos]int) int {
    count := 0
    for _, num := range stones {
	count += score(num, 75, memo)	
    }
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := strings.TrimSpace(string(data))
    numsString := strings.Split(input, " ")
    stones := make([]int, len(numsString))
    for i, n := range numsString {
	stones[i], err = strconv.Atoi(n)
	check(err)
    }

    memo := make(map[Pos]int)
    fmt.Println("part 1:", part1(stones, memo))
    fmt.Println("part 2:", part2(stones, memo))

}
