package intcodeInterpreter

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

func IncreaseIntcodeMemoryTenfold(intcode []int) []int {
	power := make([]int, len(intcode)*9)
	result := append(intcode, power...)

	return result
}

func ProcessIntcode(intcode []int, inChan chan int, outChan chan int) {
	codeLength := len(intcode)
	var stepSize int
	var pos1, pos2, pos3 int
	var ommm opcodeModes
	var relativeBase int = 0
	p1offset, p2offset, p3offset := 1, 2, 3
	for currentOp := 0; currentOp < codeLength; currentOp += stepSize {
		ommm = obtainCodeAndModes(intcode[currentOp])
		switch ommm.opcode {
		case 99:
			//output 99 so intcode user know the intcode has stopped
			outChan <- 99
			outChan <- 99
			outChan <- 99

			//release channel
			<-inChan
			return
		case 1:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)
			pos3 = getPosFromMode(ommm.mode[2], p3offset, currentOp, intcode, relativeBase)

			intcode[pos3] = intcode[pos1] + intcode[pos2]

			stepSize = 4
		case 2:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)
			pos3 = getPosFromMode(ommm.mode[2], p3offset, currentOp, intcode, relativeBase)

			intcode[pos3] = intcode[pos1] * intcode[pos2]

			stepSize = 4
		case 3:
			// signal the intcode user we're asking for input. Similar to intcode 99
			outChan <- 98
			outChan <- 98
			outChan <- 98

			// input := getInputFromUser()
			input := <-inChan
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)

			intcode[pos1] = input

			stepSize = 2
		case 4:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)

			// fmt.Println("Output:", intcode[pos1])

			outChan <- intcode[pos1]

			stepSize = 2
		case 5:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)

			if intcode[pos1] != 0 {
				currentOp = intcode[pos2]
				stepSize = 0
			} else {
				stepSize = 3
			}
		case 6:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)

			if intcode[pos1] == 0 {
				currentOp = intcode[pos2]
				stepSize = 0
			} else {
				stepSize = 3
			}
		case 7:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)
			pos3 = getPosFromMode(ommm.mode[2], p3offset, currentOp, intcode, relativeBase)

			var tBool int
			if intcode[pos1] < intcode[pos2] {
				tBool = 1
			} else {
				tBool = 0
			}

			intcode[pos3] = tBool

			stepSize = 4
		case 8:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)
			pos2 = getPosFromMode(ommm.mode[1], p2offset, currentOp, intcode, relativeBase)
			pos3 = getPosFromMode(ommm.mode[2], p3offset, currentOp, intcode, relativeBase)

			var tBool int
			if intcode[pos1] == intcode[pos2] {
				tBool = 1
			} else {
				tBool = 0
			}

			intcode[pos3] = tBool

			stepSize = 4
		case 9:
			pos1 = getPosFromMode(ommm.mode[0], p1offset, currentOp, intcode, relativeBase)

			relativeBase += intcode[pos1]

			stepSize = 2
		default:
		}
	}

	return
}

func getPosFromMode(mode int, offset int, instrPtr int, intcode []int, relBase int) int {
	var pos int
	switch mode {
	case 0: //position mode
		pos = intcode[instrPtr+offset]
	case 1: //immediate mode
		pos = instrPtr + offset
	case 2: //relative mode
		relOffset := getPosFromMode(1, offset, instrPtr, intcode, relBase)
		pos = relBase + intcode[relOffset]
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

func GetInputFromUser() int {
	var input string
	fmt.Scanln(&input)
	return toInt(input)
}

func GetIntcodeFromInput(path string) []int {
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
