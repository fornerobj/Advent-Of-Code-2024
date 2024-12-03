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

//function to see if the content following "mul" is valid
func is_valid(input string, i int) (bool, int) {
    for ; i < len(input); i++ {
        char := string(input[i])
        isNumber := false

        _, err := strconv.Atoi(char)
        if(err == nil) {
            isNumber = true
        }

        if !isNumber && char != "," && char != "(" && char != ")" {
            return false, -1
        } 

        if char == ")" {
            return true, i
        }
    }

    return false, -1
}

func mul(input string) int{
    //split "x,y" into ints x and y
    nums := strings.Split(input, ",")

    num1, err := strconv.Atoi(nums[0])
    check(err)

    num2, err := strconv.Atoi(nums[1])
    check(err)

    return num1 * num2
}


func part1(input string) int {
    count := 0
    for i := 0; i < len(input) - 3; i++ {
        //sliding window technique
        if input[i:i+3] == "mul" {
            valid, endIndex := is_valid(input, i+3)
            if valid {
                count += mul(input[i+4:endIndex])
            }
        }
    }
    return count
}

func part2(input string) int {
    count := 0
    enabled := true
    for i := 0; i < len(input) - 3; i++ {
        //several sliding windows to collect various key words
        if i < len(input)-4 && input[i:i+4] == "do()" {
            enabled = true
            continue
        }
        if i < len(input)-7 && input[i:i+7] == "don't()" {
            enabled = false
            continue
        }
        //change to only work when enabled
        if input[i:i+3] == "mul" && enabled {
            valid, endIndex := is_valid(input, i+3)
            if valid {
                count += mul(input[i+4:endIndex])
            }
        }
    }
    
    return count
}

func main() {
    filename := os.Args[1]

    data, err := os.ReadFile(filename)
    check(err)

    input := string(data)

    fmt.Println("part 1:", part1(input))
    fmt.Println("part 2:", part2(input))

}
