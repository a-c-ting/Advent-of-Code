package main

import (
	"fmt"

	"./proc"
)

const ROBO_OUT = 1
const gridSize = 50

func main() {
	Intcode := proc.GetIntcodeFromInput("Day19Part1.txt")
	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	initialState := make([]int, len(Intcode))
	copy(initialState, Intcode)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, false)
	spaceDrone(inputChannel, outputChannel, Intcode, initialState)
}

func spaceDrone(inChan chan int, outChan chan int, intcode []int, original []int) {
	var tractorMap [gridSize][gridSize]int
	droneReply := 0
	pointCount := 0

	for posY := 0; posY < gridSize; posY++ {
		for posX := 0; posX < gridSize; posX++ {
			inChan <- posX
			inChan <- posY

			//actual output
			droneReply = <-outChan
			if droneReply == 1 {
				pointCount++
			}

			tractorMap[posX][posY] = droneReply

			/*
				Intcode doesn't loop on its own this time.
				Each position scan is a different run of the intcode program.
				Program Restart and cleanup is needed per iter
			*/
			droneReply = <-outChan
			if droneReply == proc.CodeNineNine {
				inChan <- 0 //cleanup

				proc.Restart()
			}
		}
	}

	displayVisuals(tractorMap)
	fmt.Println("Total Points Covered:", pointCount)
}

func displayVisuals(mapping [gridSize][gridSize]int) {
	display := ""
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			display += convertForDisplay(mapping[x][y])
		}
		display += "\n"
	}
	fmt.Println(display)
}

func convertForDisplay(value int) string {
	var image string
	switch value {
	case 0:
		image = "░"
	case 1:
		image = "▒"
	default:
		image = "█"
	}
	return image
}
