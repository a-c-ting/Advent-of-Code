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

	//specs
	Intcode[1] = 12
	Intcode[2] = 2

	processIntcode(Intcode)
	fmt.Printf("Position 0 is %d\n", Intcode[0])
}

func processIntcode(intcode []int) {
	for currentOp := 0; currentOp < len(intcode); currentOp += 4 {
		switch intcode[currentOp] {
		case 99:
			break
		case 1:
			pos1 := intcode[currentOp+1]
			pos2 := intcode[currentOp+2]
			pos3 := intcode[currentOp+3]
			intcode[pos3] = intcode[pos1] + intcode[pos2]
		case 2:
			pos1 := intcode[currentOp+1]
			pos2 := intcode[currentOp+2]
			pos3 := intcode[currentOp+3]
			intcode[pos3] = intcode[pos1] * intcode[pos2]
		}
	}
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
