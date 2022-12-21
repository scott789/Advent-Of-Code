package main

import (
    "bufio"
    "fmt"
    "os"
)

func score(myShape string, theirShape string) int {
	var score int

	if myShape == "X" {
		score += 1
		if theirShape == "A" {
			score += 3
		} else if theirShape == "C" {
			score += 6
		}
	} else if myShape == "Y" {
		score += 2
		if theirShape == "A" {
			score += 6
		} else if theirShape == "B" {
			score += 3
		}
	} else if myShape == "Z" {
		score += 3
		if theirShape == "B" {
			score += 6
		} else if theirShape == "C" {
			score += 3
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