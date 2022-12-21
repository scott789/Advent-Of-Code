package main

import (
    "bufio"
    "fmt"
    "os"
)

func pointsForLetter(letter string) int {
	// 65 == A
	// 97 == a
	asciiValue := int(rune(letter[0]))
	if asciiValue < 97 {
		return asciiValue - 38
	} else {
		return asciiValue - 96
	}
}

func main() {
	readFile, err := os.Open("day-3-input.txt")
	defer readFile.Close()

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)

	var runningTotal int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		length := len(line)
		fmt.Printf("%s %s\n", line[0:length/2], line[length/2:length])

		ruckSack := make(map[string]int)

		for i := 0; i < length / 2; i++ {
			letter := string(line[i])
			ruckSack[letter] = ruckSack[letter] + 1
		}

		fmt.Println(ruckSack)

		for i := length / 2; i < length; i++ {
			letter := string(line[i])
			if ruckSack[letter] != 0 {
				// Found the duplicate
				fmt.Printf("letter: %s\n", letter)
				fmt.Println(ruckSack)
				runningTotal += pointsForLetter(letter)
				break
			}
		}

		fmt.Println()
    }

	fmt.Println(runningTotal)
}