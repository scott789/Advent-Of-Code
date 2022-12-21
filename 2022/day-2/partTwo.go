package main

import (
    "bufio"
    "fmt"
    "os"
)

func score(myShape, theirShape string) int {
	var score int

	if myShape == "X" { // I need to lose
		if theirShape == "A" {
			score += 3
		} else if theirShape == "B" {
			score += 1
		} else if theirShape == "C" {
			score += 2
		}
	} else if myShape == "Y" { // I need to draw
		score += 3

		if theirShape == "A" {
			score += 1
		} else if theirShape == "B" {
			score += 2
		} else if theirShape == "C" {
			score += 3
		}
	} else if myShape == "Z" { // I need to win
		score += 6

		if theirShape == "A" {
			score += 2
		} else if theirShape == "B" {
			score += 3
		} else if theirShape == "C" {
			score += 1
		}
	}

	return score
}

func main() {
	readFile, err := os.Open("day-2-input.txt")
  
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var runningTotal int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		var myShape = line[2:3]
		var theirShape = line[0:1]

		runningTotal += score(myShape, theirShape)
    }

	readFile.Close()
	fmt.Println(runningTotal)
}