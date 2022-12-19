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

	var runningTotal, max1, max2, max3 int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) != 0 {
			i, err := strconv.Atoi(line)
   		 	if err != nil {
        		panic(err)
			}
			runningTotal += i
    	} else {
			if runningTotal > max1 {
				max3 = max2
				max2 = max1
				max1 = runningTotal
			} else if runningTotal > max2 {
				max3 = max2
				max2 = runningTotal
			} else if runningTotal > max3 {
				max3 = runningTotal
			}
			runningTotal = 0
		}
    }
	readFile.Close()
	fmt.Println(max1 + max2 + max3)
}