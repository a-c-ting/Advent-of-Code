package main

import (
	"fmt"

	"./proc"
)

const ROBO_OUT = 2

type coordinates struct {
	x int
	y int
}

var itrTable = map[int]rune{
	0: '<',
	1: '^',
	2: '>',
	3: 'v',
}
var rtiTable = map[rune]int{
	'<': 0,
	'^': 1,
	'>': 2,
	'v': 3,
}

func main() {
	Intcode := proc.GetIntcodeFromInput("Day11Part1.txt")

	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, false)

	roboPainter(inputChannel, outputChannel)
}

func roboPainter(inChan chan int, outChan chan int) {
	adventurerMap := make(map[coordinates]int)

	startingPoint := coordinates{0, 0}
	startingDirection := '^'
	adventurerMap[startingPoint] = 0
	currentCoordinates := startingPoint
	currentDirection := startingDirection

	//PaintProcess
	for {
		//checkPanel
		currentPaint, alreadyExist := adventurerMap[currentCoordinates]
		if !alreadyExist {
			adventurerMap[currentCoordinates] = 0 //all panels are initially black
		}

		inChan <- currentPaint

		newPanelPaint := <-outChan
		roboTurn := <-outChan

		//halt
		if roboTurn == 99 {
			break
		}

		//paint
		adventurerMap[currentCoordinates] = newPanelPaint

		//turn and move
		currentCoordinates, currentDirection =
			roboMove(currentCoordinates, currentDirection, roboTurn)
	}
	fmt.Println(len(adventurerMap))
}

func roboMove(c coordinates, d rune, turn int) (coordinates, rune) {
	newDirection := roboChangeDirection(d, turn)
	newCoordinates := c
	switch newDirection {
	case '<':
		newCoordinates.x = c.x - 1
	case '^':
		newCoordinates.y = c.y + 1
	case '>':
		newCoordinates.x = c.x + 1
	case 'v':
		newCoordinates.y = c.y - 1
	}
	return newCoordinates, newDirection
}

func roboChangeDirection(d rune, turn int) rune {
	if turn == 0 {
		turn = -1
	}

	offset := len(itrTable)
	iPreTurn := rtiTable[d]
	iPostTurn := (iPreTurn + turn + offset) % offset
	resultDirection := itrTable[iPostTurn]

	return resultDirection
}
