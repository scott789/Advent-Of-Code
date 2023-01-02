package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
	"strconv"
)

func firstInsideSecond(num1, num2, num3, num4 int) bool {
	return num1 >= num3 && num2 <= num4;
}

func SecondInsideFirst(num1, num2, num3, num4 int) bool {
	return num3 >= num1 && num4 <= num2;
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var runningTotal int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		ranges := strings.Split(line, ",")

		range0 := strings.Split(ranges[0], "-")
		range1 := strings.Split(ranges[1], "-")

		range0Num0, err := strconv.Atoi(range0[0])
		if err != nil {
			panic(err)
		}

		range0Num1, err := strconv.Atoi(range0[1])
		if err != nil {
			panic(err)
		}

		range1Num0, err := strconv.Atoi(range1[0])
		if err != nil {
			panic(err)
		}

		range1Num1, err := strconv.Atoi(range1[1])
		if err != nil {
			panic(err)
		}

		if firstInsideSecond(range0Num0, range0Num1, range1Num0, range1Num1) || SecondInsideFirst(range0Num0, range0Num1, range1Num0, range1Num1) {
			runningTotal++
		}
    }

	fmt.Println(runningTotal)
}