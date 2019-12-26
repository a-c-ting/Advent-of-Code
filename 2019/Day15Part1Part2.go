package main

import (
	"fmt"
	"math"
	"time"

	"./proc"
)

const ROBO_OUT = 1

type coordinates struct {
	x int
	y int
}

func main() {
	Intcode := proc.GetIntcodeFromInput("Day15Part1.txt")
	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, false)

	gameMap := make(map[coordinates]int)
	searchAndRepair(inputChannel, outputChannel, gameMap)
}

func searchAndRepair(inChan chan int, outChan chan int, gameMap map[coordinates]int) {
	var xlen, ylen, xmin, ymin int
	startingPoint := coordinates{0, 0}
	currentCoordinates := startingPoint
	gameMap[currentCoordinates] = 1 //passable terrain
	commandCount := 0

	for { //play Process
		xlen, ylen, xmin, ymin = getMapParams(gameMap)
		visualizeFrame(xlen, ylen, xmin, ymin, gameMap, currentCoordinates)

		/** Manual Play Version		**/
		/** The out label and for loop is so that invalid input won't end the entire game**/
		tempInput := translateInput(proc.GetInputFromUser())
	out:
		for {
			switch tempInput {
			case 1:
				break out
			case 2:
				break out
			case 3:
				break out
			case 4:
				break out
			case 5: //this should allow us to measure distance, used for Part 2
				commandCount = 0
				tempInput = translateInput(proc.GetInputFromUser())
			case 6:
				fmt.Println("Current Steps", commandCount)
				tempInput = translateInput(proc.GetInputFromUser())
			default:
				fmt.Println("invalid input")
				tempInput = translateInput(proc.GetInputFromUser())
			}
		}

		/** Auto Play Version		**/
		// TODO: AutoPlay version when I get more time

		inChan <- tempInput

		remoteBotReply := <-outChan

		currentCoordinates = roboRecordAndMove(remoteBotReply, currentCoordinates, tempInput, gameMap)

		if remoteBotReply == 1 {
			commandCount++
		} else if remoteBotReply == 2 { //halt
			visualizeFrame(xlen, ylen, xmin, ymin, gameMap, currentCoordinates)
			commandCount++
			fmt.Println("Total Command Count is", commandCount)
			break
		}
	}
	visualizeFrame(xlen, ylen, xmin, ymin, gameMap, currentCoordinates)
}

func roboRecordAndMove(botReply int, currentPos coordinates, movement int, gameMap map[coordinates]int) coordinates {
	result := currentPos
	cx, cy := getPosChange(movement)
	projectedPos := coordinates{x: currentPos.x + cx, y: currentPos.y + cy}

	switch botReply {
	case 0:
	case 1:
		result = projectedPos
	case 2:
	}

	gameMap[projectedPos] = botReply

	return result
}

func getPosChange(movement int) (int, int) {
	switch movement { //following cartesian plane
	case 1:
		return 0, 1
	case 2:
		return 0, -1
	case 3:
		return -1, 0
	case 4:
		return 1, 0
	default:
		return 0, 0
	}
}

func visualizeFrame(xlen, ylen, xmin, ymin int, gameMap map[coordinates]int, currentPos coordinates) {
	displayString := ""
	fmt.Println("test")
	for j := (ymin + ylen) - 1; j >= ymin; j-- {
		for i := xmin; i < (xmin + xlen); i++ {
			paintCoordinates := coordinates{x: i, y: j}
			panelColor, exists := gameMap[paintCoordinates]
			visual := ""

			if !exists {
				panelColor = 99
			}
			if (i == currentPos.x) && (j == currentPos.y) {
				panelColor = 3
			}

			switch panelColor {
			case 0: //wall
				visual = "██"
			case 1: //passable terrain
				visual = "▒▒"
			case 2: //oxigen
				visual = "[]"
			case 3: //robot
				visual = "()"
			default:
				visual = "░░" //empty
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

func translateInput(onetwothree int) int {
	//NOTE: This is no longer used for solution as it is very hard to actually play the game.
	//These function is to set the keypad as the joystey input
	//If you wish to use a non-keypad set of input (like WASD or ZXC), you'd need to
	//modify the GetInputFromUser in intcoodeProcessor.go to allow for non-Integers input
	result := 99
	switch onetwothree {
	case 5: //north
		result = 1
	case 2: //south
		result = 2
	case 1: //west
		result = 3
	case 3: //east
		result = 4
	case 7: //reset counter
		result = 5
	case 9: //print counter
		result = 6
	default: //invalid
		result = 0
	}
	return result
}
