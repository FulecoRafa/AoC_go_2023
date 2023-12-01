package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	lib "github.com/fulecorafa/aoc_2023_lib"
)

func getFirstNum(src string) (byte, bool) {
    for i := 0 ; i < len(src) ; i++ {
        char := src[i]
        if char >= '0' && char <= '9' {
            return char, true
        }
    }
    return 0, false
}

func getLastNum(src string) (byte, bool) {
    for i := len(src) - 1; i >= 0 ; i-- {
        char := src[i]
        if char >= '0' && char <= '9' {
            return char, true
        }
    }
    return 0, false
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanLines)

    result := 0
    for fileScanner.Scan() {
        line := fileScanner.Text()

        firstNum, ok := getFirstNum(line)
        lib.MayPanicBool(ok)
        
        lastNum, ok := getLastNum(line)
        lib.MayPanicBool(ok)

        strNum := string([]byte{firstNum, lastNum})

        num, err := strconv.Atoi(strNum)
        lib.MayPanic(err)
        result += num
    }

    fmt.Println("Result: ", result)
}
