package main

import (
	"fmt"
	"time"

	"./proc"
)

const ROBO_OUT = 1

func main() {
	Intcode := proc.GetIntcodeFromInput("Day17Part1.txt")
	Intcode = proc.IncreaseIntcodeMemoryTenfold(Intcode)

	Intcode[0] = 2 //Wake Up Robot

	/*Solved by hand. Need to be able to do the basics first.*/
	mainRoutine := "A,B,A,C,B,C,A,C,B,C\n"
	subA := "L,8,R,10,L,10\n"
	subB := "R,10,L,8,L,8,L,10\n"
	subC := "L,4,L,6,L,8,L,8\n"
	visualsSwitch := "y\n"
	inputSequence := convertToInputArray(mainRoutine + subA + subB + subC + visualsSwitch)

	inputChannel := make(chan int)
	outputChannel := make(chan int, ROBO_OUT)

	go proc.ProcessIntcode(Intcode, inputChannel, outputChannel, ROBO_OUT, true)
	robotAction(inputChannel, outputChannel, inputSequence)
}

func convertToInputArray(input string) []rune {
	var inputSequence []rune
	for _, xRune := range input {
		inputSequence = append(inputSequence, xRune)
	}
	return inputSequence
}

func robotAction(inChan chan int, outChan chan int, inSeq []rune) {
	seqLen := len(inSeq)
	prevReply := 0
	botReply := 0
	var displayBuffer string

	index := 0
loop:
	for {
		switch botReply {
		case proc.CodeInput:
			currentInput := int(inSeq[index])
			inChan <- currentInput
			fmt.Printf("%c", currentInput)

			if currentInput == int('\n') {
				time.Sleep(1000 * time.Millisecond)
			}

			index++ //input tracking

			botReply = <-outChan
		case proc.CodeNineNine:
			fmt.Printf("\nOpcode Nine-Nine : End of Program\n")

			break loop
		default:
			if index != seqLen {
				fmt.Printf("%c", botReply)
				botReply = <-outChan
			} else { //frame by frame display
				prevReply = botReply
				botReply = <-outChan

				//intcode returns two '\n' when finishing a whole video frame
				if prevReply == int('\n') && botReply == int('\n') {
					time.Sleep(50 * time.Millisecond)
					fmt.Printf("\n%s", displayBuffer)
					displayBuffer = ""
				} else { //buffer if not yet complete
					displayBuffer += convertToDisplay(botReply)
				}
			}
		}
	}
	fmt.Println("Total Dust Gathered:", prevReply)
}

func convertToDisplay(value int) string {
	return fmt.Sprintf("%s", string(value))
}
