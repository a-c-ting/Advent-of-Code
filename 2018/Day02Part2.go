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

func getCheckSum(stringArray []string) int {
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
	return checkSum
}

func compareBoxId(str1, str2 string) (bool, string) {
	length := len(str1)
	reducedLen := length-1
	if (length != len(str2)) {
		return false, ""
	}

	for i := 0; i < reducedLen; i++ {
		reducedStr1 := str1[:i] + str1[i+1:]
		reducedStr2 := str2[:i] + str2[i+1:]
		if reducedStr1 == reducedStr2 {
			return true, reducedStr1
		}
	}
	return false, ""
}

func searchCorrectBoxes(boxidList []string) string {
	var correctBoxesFound bool
	var almostMatchingIds string
	for index, boxid1 := range boxidList {
		if correctBoxesFound {
			break
		}
		//index+1 is needed to avoid comparison with self. Another way might be to skip compareBoxId if boxid1 == boxid2
		for _, boxid2 := range boxidList[index+1:] {
			correctBoxesFound, almostMatchingIds = compareBoxId(boxid1, boxid2)
			if correctBoxesFound {
				fmt.Println(almostMatchingIds)
				break
			}
		}
	}
	return almostMatchingIds
}

func main() {
	boxidList := getInputFromFile("Day02Input.txt")

	//Part 1
	fmt.Printf("Checksum is: %d\n", getCheckSum(boxidList))

	//Part 2
	fmt.Printf("Correct Common Letters are: %s", searchCorrectBoxes(boxidList))
}