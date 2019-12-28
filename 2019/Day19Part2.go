package main

import (
	"fmt"

	"./proc"
)

const ROBO_OUT = 1

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
	droneReply := 0
	weFoundSanta := false

	index := 0 //start at 1300 if you wanna see the answer quickly
	for {
		for posY := 0; posY < index+1; posY++ { //index+1
			droneReply = droneCheck(inChan, outChan, index, posY)

			if droneReply == 1 {
				edge1 := droneCheck(inChan, outChan, index+99, posY)
				edge2 := droneCheck(inChan, outChan, index, posY+99)
				edge3 := droneCheck(inChan, outChan, index+99, posY+99)
				if (edge1 == 1) && (edge2 == 1) && (edge3 == 1) {
					weFoundSanta = true
					defer fmt.Println("Santa Location:", index, posY, "AoC Answer:", index*10000+posY)
				}
			}
		}
		for posX := 0; posX < index; posX++ {
			droneReply = droneCheck(inChan, outChan, posX, index)

			if droneReply == 1 {
				edge1 := droneCheck(inChan, outChan, posX+99, index)
				edge2 := droneCheck(inChan, outChan, posX, index+99)
				edge3 := droneCheck(inChan, outChan, posX+99, index+99)
				if (edge1 == 1) && (edge2 == 1) && (edge3 == 1) {
					weFoundSanta = true
					defer fmt.Println("Santa Location:", posX, index, "AoC Answer:", posX*10000+index)
				}
			}
		}

		//index = map length/width
		fmt.Println("Index Count:", index, "santa: ", weFoundSanta)

		if weFoundSanta {
			break
		}

		index++
	}
}

func droneCheck(inChan chan int, outChan chan int, xpos int, ypos int) int {
	inChan <- xpos
	inChan <- ypos

	droneFeed := <-outChan

	droneReply := <-outChan
	if droneReply == proc.CodeNineNine {
		inChan <- 0 //cleanup

		proc.Restart()
	}
	return droneFeed
}
