package main

import (
	"fmt"

	"./proc"
)

const ROBO_OUT = 1

func main() {
	Intcode := proc.GetIntcodeFromInput("Day17Part1.txt")
	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, false)
	shipMap := getCameraFeed(inputChannel, outputChannel)
	findTotalAlignmentParam(shipMap)
}

func findTotalAlignmentParam(shipMap [][]string) {
	rows := len(shipMap)
	col := len(shipMap[0])
	sum := 0
	for i := 1; i < rows-2; i++ { //intersection can't occur on edges
		for j := 1; j < col-2; j++ {
			up := shipMap[i+1][j]
			down := shipMap[i-1][j]
			right := shipMap[i][j+1]
			left := shipMap[i][j-1]
			if isIntersection(up, down, left, right) {
				sum += i * j
			}
		}
	}
	fmt.Println("Sum of Alignment Param is", sum)
}

func isIntersection(up, down, left, right string) bool {
	return (up == "#") &&
		(down == "#") &&
		(left == "#") &&
		(right == "#")
}

func getCameraFeed(inChan chan int, outChan chan int) [][]string {
	var shipMap [][]string
	var mapRow []string

	for {
		temp := <-outChan
		if temp == proc.CodeNineNine {
			break
		}
		fmt.Printf(convertToDisplay(temp))
		printable := convertToDisplay(temp)

		mapRow = append(mapRow, printable)
		if printable == "\n" {
			shipMap = append(shipMap, mapRow)
			mapRow = nil
		}
	}
	inChan <- 0

	return shipMap
}

func convertToDisplay(value int) string {
	return fmt.Sprintf("%s", string(value)) //this one works
}
