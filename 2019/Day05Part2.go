package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const MAX_MODES = 3

type opcodeModes struct {
	opcode int
	mode   [MAX_MODES]int
}

func main() {
	Intcode := getIntcodeFromInput("Day05Part1.txt")

	processIntcode(Intcode)
}

func processIntcode(intcode []int) {
	codeLength := len(intcode)
	p1pos, p2pos := 1, 2
	var stepSize int
	var pos1, pos2, pos3 int
	var ommm opcodeModes
	var opcode int
	for currentOp := 0; currentOp < codeLength; currentOp += stepSize {
		ommm = obtainCodeAndModes(intcode[currentOp])
		opcode = ommm.opcode
		switch opcode {
		case 99:
			return
		case 1:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)
			pos3 = intcode[currentOp+3] //write mode so not using immediate mode

			intcode[pos3] = intcode[pos1] + intcode[pos2]

			stepSize = 4
		case 2:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)
			pos3 = intcode[currentOp+3] //write mode so not using immediate mode

			intcode[pos3] = intcode[pos1] * intcode[pos2]

			stepSize = 4
		case 3:
			input := getInputFromUser()
			pos1 = intcode[currentOp+1] //write mode so not using immediate mode
			intcode[pos1] = input

			stepSize = 2
		case 4:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)

			fmt.Println("Output:", intcode[pos1])
			stepSize = 2
		case 5:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)

			if intcode[pos1] != 0 {
				currentOp = intcode[pos2]
				stepSize = 0
			} else {
				stepSize = 3
			}
		case 6:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)

			if intcode[pos1] == 0 {
				currentOp = intcode[pos2]
				stepSize = 0
			} else {
				stepSize = 3
			}
		case 7:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)
			pos3 = intcode[currentOp+3] //write mode so not using immediate mode

			var tBool int
			if intcode[pos1] < intcode[pos2] {
				tBool = 1
			} else {
				tBool = 0
			}

			intcode[pos3] = tBool

			stepSize = 4
		case 8:
			pos1 = getPosFromMode(ommm.mode[0], p1pos, currentOp, intcode)
			pos2 = getPosFromMode(ommm.mode[1], p2pos, currentOp, intcode)
			pos3 = intcode[currentOp+3] //write mode so not using immediate mode

			var tBool int
			if intcode[pos1] == intcode[pos2] {
				tBool = 1
			} else {
				tBool = 0
			}

			intcode[pos3] = tBool

			stepSize = 4
		default:
		}
	}
	return
}

func getPosFromMode(mode int, paramPosition int, instrPtr int, intcode []int) int {
	var pos int
	if mode == 0 {
		pos = intcode[instrPtr+paramPosition]
	} else {
		pos = instrPtr + paramPosition
	}
	return pos
}

func obtainCodeAndModes(ommm int) opcodeModes {
	//no error checking on input. Codes are expected to be in ABCDE
	stringVer := toString(ommm)
	lenStr := len(stringVer)
	parsedOmmm := opcodeModes{
		opcode: 99,
		mode:   [MAX_MODES]int{0, 0, 0},
	}
	var modes [MAX_MODES]int

	if lenStr <= 0 {
		//ends the program since it has opcode99
		return parsedOmmm
	}

	if lenStr <= 2 {
		parsedOmmm.opcode = toInt(stringVer)
	} else {
		parsedOmmm.opcode = toInt(stringVer[lenStr-2 : lenStr])
	}
	lenStr -= 2

	k := 0
	for i := lenStr; i > 0; i-- {
		modes[k] = toInt(stringVer[i-1 : i])
		k++
	}

	parsedOmmm.mode = modes
	return parsedOmmm
}

func getInputFromUser() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}

func getIntcodeFromInput(path string) []int {
	stream, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	f := func(c rune) bool {
		return c == ',' || c == '\n'
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

func toString(i int) string {
	str := strconv.Itoa(i)
	return str
}
