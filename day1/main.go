package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func findNumOccurences(slice []int, val int) int {
    numOccurences := 0
    for _, num := range slice {
        if num == val {
            numOccurences += 1
        }
    }
    return numOccurences
}

func find(slice []int, val int) int {
    for i, num := range slice {
        if num == val {
            return i
        }
    }
    return -1;
}

func delete(slice []int, val int) []int {
    pos := find(slice, val)
    return append(slice[:pos], slice[pos+1:]...)
}

func min(slice []int) int {
    minNum := slice[0]
    for _, num := range slice {
        if num < minNum {
            minNum = num
        }
    }
    return minNum
}

func part1(list1 []int, list2 []int) int {
    l1 := make([]int, len(list1))
    l2 := make([]int, len(list2))
    copy(l1, list1)
    copy(l2, list2)

    ans := 0
    for len(l1) > 0 {
        minNum := min(l1)
        minNum2 := min(l2)
        l1 = delete(l1, minNum)
        l2 = delete(l2, minNum2)
        ans += int(math.Abs(float64(minNum - minNum2)))
    }
    return ans
}

func part2(list1 []int, list2 []int) int {
    ans := 0
    for _, num := range(list1) {
        ans += num * findNumOccurences(list2, num)
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

    list1 := []int{}
    list2 := []int{}
    scanner := bufio.NewScanner(data)
    for scanner.Scan() {
        line := scanner.Text()
        nums := strings.Fields(line)
        num1, err := strconv.Atoi(nums[0])
        if err != nil {
            fmt.Println("Failed to convert string to integer")
            return
        }
        num2, err := strconv.Atoi(nums[1])
        if err != nil {
            fmt.Println("Failed to convert string to integer")
            return
        }
        list1 = append(list1, num1)
        list2 = append(list2, num2)
    }

    fmt.Println("part 1:", part1(list1, list2))
    fmt.Println("part 2:", part2(list1, list2))

    if scanner.Err() != nil {
        fmt.Println("Error scanning files:", scanner.Err())
    }
}
