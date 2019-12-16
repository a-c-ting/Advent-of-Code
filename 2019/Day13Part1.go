package main

import (
	"fmt"
	"math"

	proc "./Day13"
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

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel)

	gameMap := make(map[coordinates]int)
	mapMaker(inputChannel, outputChannel, gameMap)

	xlen, ylen, xmin, ymin := getMapParams(gameMap)

	visualizePanels(xlen, ylen, xmin, ymin, gameMap)

	blockTileCount := 0
	for _, v := range gameMap {
		if v == 2 {
			blockTileCount++
		}
	}

	fmt.Println("Total Block Tiles Count is", blockTileCount)
}

func visualizePanels(xlen, ylen, xmin, ymin int, gameMap map[coordinates]int) {
	for j := ymin; j < ylen; j++ {
		for i := xmin; i < xlen; i++ {
			paintCoordinates := coordinates{x: i, y: j}
			panelColor := gameMap[paintCoordinates]
			if panelColor == 0 {
				fmt.Printf("░░")
			} else if panelColor == 1 {
				fmt.Printf("██")
			} else if panelColor == 2 {
				fmt.Printf("▒▒")
			} else if panelColor == 3 {
				fmt.Printf("▓▓")
			} else if panelColor == 4 {
				fmt.Printf("()")
			}
		}
		fmt.Printf("\n")
	}
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

func mapMaker(inChan chan int, outChan chan int, gameMap map[coordinates]int) {
	for {
		xcoord := <-outChan
		ycoord := <-outChan
		tileID := <-outChan

		if tileID == 99 { //halt
			break
		}

		currentCoordinates := coordinates{xcoord, ycoord}
		gameMap[currentCoordinates] = tileID //populate map
	}
}
