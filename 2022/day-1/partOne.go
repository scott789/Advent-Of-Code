package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

func main() {
	readFile, err := os.Open("day-1-input.txt")
  
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var runningTotal, maximum int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) != 0 {
			i, err := strconv.Atoi(line)
   		 	if err != nil {
        		panic(err)
			}
			runningTotal += i
    	} else {
			if runningTotal > maximum {
				maximum = runningTotal
			}
			runningTotal = 0
		}
    }
	readFile.Close()
	fmt.Println(maximum)
}