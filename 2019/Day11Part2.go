package main

import (
	"fmt"
	"math"

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

	adventurerMap := make(map[coordinates]int)
	roboPainter(inputChannel, outputChannel, adventurerMap)

	xlen, ylen, xmin, ymin := getMapParams(adventurerMap)
	visualizePanels(xlen, ylen, xmin, ymin, adventurerMap)
}

func visualizePanels(xlen, ylen, xmin, ymin int, adventurerMap map[coordinates]int) {
	for j := (ymin + ylen) - 1; j >= ymin; j-- {
		for i := xmin; i < (xmin + xlen); i++ {
			paintCoordinates := coordinates{x: i, y: j}
			panelColor := adventurerMap[paintCoordinates]
			if panelColor == 0 {
				fmt.Printf("░░░")
			} else {
				fmt.Printf("███")
			}
		}
		fmt.Printf("\n")
	}
}

func getMapParams(adventurerMap map[coordinates]int) (int, int, int, int) {
	var minX, maxX, minY, maxY int
	for keys := range adventurerMap {
		if keys.x > maxX {
			maxX = keys.x
		}
		if keys.x < minX {
			minX = keys.x
		}
		if keys.y > maxY {
			maxY = keys.y
		}
		if keys.y < minY {
			minY = keys.y
		}
	}
	return int(math.Abs(float64(maxX)) + math.Abs(float64(minX)) + 1),
		int(math.Abs(float64(maxY)) + math.Abs(float64(minY)) + 1),
		minX, minY
}

func roboPainter(inChan chan int, outChan chan int, adventurerMap map[coordinates]int) {
	startingPoint := coordinates{0, 0}
	startingDirection := '^'
	adventurerMap[startingPoint] = 1 //first panel is white == 1
	currentCoordinates := startingPoint
	currentDirection := startingDirection

	//PaintProcess
	for {
		//checkPanel
		currentPaint, alreadyExist := adventurerMap[currentCoordinates]
		if !alreadyExist {
			adventurerMap[currentCoordinates] = 0 //all panels are initially black except for start
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
