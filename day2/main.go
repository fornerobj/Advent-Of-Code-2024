package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func isSafe(nums []int) bool {
    isIncreasing := false
    prev := nums[0]
    for i, val := range nums {
        if i == 0 {
            continue
        }else if i == 1 {
            isIncreasing = val > prev
        } 
        dif := int(math.Abs(float64(val-prev)))
        if dif < 1 || dif > 3 {
            return false
        }
        if isIncreasing {
            if val < prev {
                return false
            }
        } else {
            if val > prev {
                return false
            }
        }
        prev = val
    }

    return true
}

func isSafeWithDampener(nums []int) bool {
    //basically, try isSafe, and if not safe, try isSafe with each number missing
    //probably not the best way, but its 1 am
    if isSafe(nums) {
        return true
    }
    for i := 0; i < len(nums); i++ {
        numsWithoutI := make([]int, len(nums)-1)
        copy(numsWithoutI, nums[:i])
        copy(numsWithoutI[i:], nums[i+1:])
        if isSafe(numsWithoutI) {
            return true
        }
    }

    return false
}

func part1(list [][]int) int {
    ans := 0
    for _, l := range list {
        if isSafe(l){
            ans += 1
        }
    }
    return ans
}

func part2(list [][]int) int {
    ans := 0
    for _, l := range list {
        if isSafeWithDampener(l){
            ans += 1
        }
    }
    return ans
}

func main() {
    filename := os.Args[1]

    data, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file. Possible invalid file name")
        return
    }
    defer data.Close()

    list := [][]int{}
    scanner := bufio.NewScanner(data)
    for scanner.Scan() {
        line := scanner.Text()
        nums := strings.Fields(line)

        listToAdd := []int{}
        for _, num := range nums {
            numToAdd, err := strconv.Atoi(num);
            if(err != nil) {
                return
            }
            listToAdd = append(listToAdd, numToAdd)
        }
        list = append(list, listToAdd)
    }

    fmt.Println("part 1:", part1(list))
    fmt.Println("part 2:", part2(list))

    if scanner.Err() != nil {
        fmt.Println("Error scanning files:", scanner.Err())
    }
}
