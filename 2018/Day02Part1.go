package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getInputFromFile(path string) []string {
	stream, err := ioutil.ReadFile(path)
	if (err != nil) {
		return nil
	}

	result := strings.Fields(string(stream))
	return result
}

func createTally(boxid string) map[string]int {
	idLength := len(boxid)
	tally := make(map[string]int)
	for i := 0; i < idLength; i++ {
		count, _ := tally[boxid[i:i+1]]
		tally[boxid[i:i+1]] = count + 1
	}
	return tally
}

func countChecks(boxid string) (bool, bool) {
	tally := createTally(boxid)
	var twos, threes bool
	for _, count := range tally {
		if (count == 3) {
			threes = true
		}
		if (count == 2) {
			twos = true
		}
	}
	return twos, threes
}

func main() {
	stringArray := getInputFromFile("Day02Input.txt")
	var totalTwos, totalThrees int

	for _, value := range stringArray {
		twoExists, threeExists := countChecks(value)
		if twoExists {
			totalTwos += 1
		}
		if threeExists {
			totalThrees += 1
		}
	}
	checkSum := totalTwos*totalThrees
	fmt.Printf("Checksum is %d", checkSum)
}