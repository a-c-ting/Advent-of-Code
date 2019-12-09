package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func getInputFromFile(path string) []string {
	stream, err := ioutil.ReadFile(path)
	if (err != nil) {
		return nil
	}

	result := strings.Fields(string(stream))
	return result
}

func main() {
	stringArray := getInputFromFile("Day01Input.txt")
	reachedFreqValues := make(map[int]int)
	var finalFreq int
	alreadyExists := false

	/*	Note: Remove the outer for loop to get the Part 1 Answer	*/
	i:=0
	for !alreadyExists {
		for _, value := range stringArray {
			freqMovement, _ := strconv.Atoi(value)
			finalFreq += freqMovement
			_, alreadyExists = reachedFreqValues[finalFreq]
			if !alreadyExists {
				reachedFreqValues[finalFreq] = 1
				// fmt.Println("not found!")
			} else {
				// fmt.Println(finalFreq)
				break
			}
			// fmt.Println(elem)
		}
		// fmt.Println(i)
		i++
	}

	fmt.Println(finalFreq)
	fmt.Printf("Total List Repetition: %d", i)
}