package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	Intcode := getIntcodeFromInput("Day02Part1.txt")

	desiredOutput := 19690720
	fmt.Println("noun*100 + verb =", findMatchingInput(Intcode, desiredOutput, 0, 99))
}

func findMatchingInput(intcode []int, desiredOutput int, minRange int, maxRange int) int {
	for noun := minRange; noun < maxRange+1; noun++ {
		for verb := minRange; verb < maxRange+1; verb++ {
			temp := make([]int, len(intcode))
			//copy is needed to make a duplicate of of intcode, which AFAIK is a slice due to append function
			copy(temp, intcode)
			temp[1] = noun
			temp[2] = verb

			processIntcode(temp)

			if temp[0] == desiredOutput {
				return 100*noun + verb
			}
		}
	}
	return -1
}

func processIntcode(intcode []int) {
	codeLength := len(intcode)
	for currentOp := 0; currentOp < codeLength; currentOp += 4 {
		switch intcode[currentOp] {
		case 99:
			break
		case 1:
			// if (currentOp + 3) > codeLength {
			// 	return
			// }
			pos1 := intcode[currentOp+1]
			pos2 := intcode[currentOp+2]
			pos3 := intcode[currentOp+3]
			// if (pos1 > codeLength) || (pos2 > codeLength) || (pos3 > codeLength) {
			// 	return
			// }
			intcode[pos3] = intcode[pos1] + intcode[pos2]
		case 2:
			// if (currentOp + 3) > codeLength {
			// 	return
			// }
			pos1 := intcode[currentOp+1]
			pos2 := intcode[currentOp+2]
			pos3 := intcode[currentOp+3]
			// if (pos1 > codeLength) || (pos2 > codeLength) || (pos3 > codeLength) {
			// 	return
			// }
			intcode[pos3] = intcode[pos1] * intcode[pos2]
		}
	}
	return
}

func getIntcodeFromInput(path string) []int {
	stream, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	delimRemovedStream := strings.FieldsFunc(string(stream), f)

	intcodeArray := []int{}
	for _, value := range delimRemovedStream {
		intcodeArray = append(intcodeArray, toInt(value))
	}

	return intcodeArray
}

func toInt(s string) int {
	//only numbers allowed so err should be impossible
	val, _ := strconv.Atoi(s)
	return val
}
