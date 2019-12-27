package main

import (
	"fmt"
	"math"
	"time"

	"./proc"
)

const ROBO_OUT = 3

type coordinates struct {
	x int
	y int
}

func main() {
	Intcode := proc.GetIntcodeFromInput("Day13Part1.txt")
	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	Intcode[0] = 2 //freeplay mode

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, true)

	gameMap := make(map[coordinates]int)
	game(inputChannel, outputChannel, gameMap)
}

func visualizeFrame(xlen, ylen, xmin, ymin int, gameMap map[coordinates]int) {
	displayString := ""
	for j := ymin; j < ylen; j++ {
		for i := xmin; i < xlen; i++ {
			paintCoordinates := coordinates{x: i, y: j}
			panelColor := gameMap[paintCoordinates]
			visual := ""
			switch panelColor {
			case 0:
				visual = "░░"
			case 1:
				visual = "██"
			case 2:
				visual = "▒▒"
			case 3:
				visual = "▓▓"
			case 4:
				visual = "()"
			default:
				visual = "XX"
			}
			displayString += visual
		}
		displayString += "\n"
	}
	fmt.Printf("%s", displayString)
	time.Sleep(20 * time.Millisecond) //delay for framerate
}

func getMapParams(gameMap map[coordinates]int) (int, int, int, int) {
	var minX, maxX, minY, maxY int
	for keys := range gameMap {
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

func game(inChan chan int, outChan chan int, gameMap map[coordinates]int) {
	var xlen, ylen, xmin, ymin int

	for { //play Process
		xcoord := <-outChan
		ycoord := <-outChan
		tileID := <-outChan

		if tileID == proc.CodeNineNine { //halt
			scoreXY := coordinates{-1, 0}
			fmt.Println("Final Score is :", gameMap[scoreXY])
			break
		} else if tileID == proc.CodeInput { //input, and is also the "framerate" of our game
			if (xlen == 0) && (ylen == 0) { //we use only the initial boundaries
				xlen, ylen, xmin, ymin = getMapParams(gameMap)
			}
			visualizeFrame(xlen, ylen, xmin, ymin, gameMap)

			/** Manual Play Version		**/
			// inputTemp := proc.GetInputFromUser()
			// inChan <- translateInput(inputTemp)

			/** Auto-Play Version 		**/
			inChan <- autoPlay(gameMap)
		}

		currentCoordinates := coordinates{xcoord, ycoord}

		gameMap[currentCoordinates] = tileID //populate map
	}
}

func autoPlay(gameMap map[coordinates]int) int {
	var ball coordinates
	var paddle coordinates

	for k, v := range gameMap {
		if v == 4 { //ball
			ball = k
		} else if v == 3 { //padle
			paddle = k
		}
	}

	if ball.x > paddle.x {
		return 1
	} else if ball.x < paddle.x {
		return -1
	} else {
		return 0
	}
}

func translateInput(onetwothree int) int {
	//NOTE: This is no longer used for solution as it is very hard to actually play the game.
	//These function is to set the keypad as the joystey input
	//If you wish to use a non-keypad set of input (like WASD or ZXC), you'd need to
	//modify the GetInputFromUser in intcoodeProcessor.go to allow for non-Integers input
	result := 99
	switch onetwothree {
	case 1: //go left
		result = -1
	case 3: //go right
		result = 1
	default:
		//case 2 and every other kind of input for easier play
		result = 0
	}
	return result
}
