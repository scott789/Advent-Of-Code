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
	var counter int
	var ruckSack0, ruckSack1, ruckSack2 map[string]int
  
    for fileScanner.Scan() {
		line := fileScanner.Text()
		length := len(line)

		ruckSack := make(map[string]int)

		for i := 0; i < length; i++ {
			letter := string(line[i])
			ruckSack[letter] = ruckSack[letter] + 1
		}

		if counter == 0 {
			ruckSack0 = ruckSack
		} else if counter == 1 {
			ruckSack1 = ruckSack
		} else {
			ruckSack2 = ruckSack
		}

		fmt.Println(ruckSack)

		if counter == 2 {
			// Compare A to B and store off intersection
			intersectionSackOneAndTwo := make(map[string]int)
			for k, v := range ruckSack0 {
				if ruckSack1[k] != 0 {
					// Found a duplicate between A and B
					fmt.Printf("letter: %s %d\n", k, v)
					intersectionSackOneAndTwo[k] = intersectionSackOneAndTwo[k] + 1
				}
			}

			fmt.Printf("intersection of 0 and 1: %v\n", intersectionSackOneAndTwo)
			
			// Then compare AB to C and store off intersection (short-circuit on the first match)
			for k, v := range intersectionSackOneAndTwo {
				letter := k
				if ruckSack2[letter] != 0 {
					// Found a duplicate between AB and C
					fmt.Printf("letter: %s %d\n", letter, v)
					runningTotal += pointsForLetter(letter)
					break
				}
			}

			// Reset to 0
			counter = 0
			ruckSack0 = nil
			ruckSack1 = nil
			ruckSack2 = nil
		} else {
			counter++
		}

		fmt.Println()
    }

	fmt.Println(runningTotal)
}