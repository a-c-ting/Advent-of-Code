package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

const LAYER_WIDTH = 25
const LAYER_HEIGHT = 6

type encLayer struct {
	layerID   int
	layerInfo [LAYER_HEIGHT][LAYER_WIDTH]int
}

func main() {
	encodedSequence := getEncodedStreamFromInput("Day08Part1.txt")

	layers := createLayers(&encodedSequence)
	fmt.Println(layers)

	targetLayer := findFewestZeroLayerInfo(&layers)
	fmt.Printf("Target Layer is %d\n", targetLayer)
	_, ones, twos := countZeroOneTwos(&layers[targetLayer])
	fmt.Printf("Check Number is %d\n", ones*twos)
}

func findFewestZeroLayerInfo(layers *[]encLayer) int {
	//assumption is only one fewest zero layer as per specs (problem)
	lowestZeroes := math.MaxInt32
	lowestZeroLayer := -1
	for _, individualLayer := range *layers {
		var zeros int

		zeros, _, _ = countZeroOneTwos(&individualLayer)

		if zeros < lowestZeroes {
			lowestZeroes = zeros
			lowestZeroLayer = individualLayer.layerID
		}
	}
	return lowestZeroLayer
}

func countZeroOneTwos(layer *encLayer) (int, int, int) {
	var zeros, ones, twos int
	for _, innerArray := range layer.layerInfo {
		for _, elements := range innerArray {
			switch elements {
			case 0:
				zeros++
			case 1:
				ones++
			case 2:
				twos++
			default:
			}
		}
	}
	return zeros, ones, twos
}

func createLayers(encStream *[]byte) []encLayer {
	var layers []encLayer
	len := len(*encStream)
	layerCapacity := LAYER_WIDTH * LAYER_HEIGHT

	for cLayerNum := 0; cLayerNum*layerCapacity < len; cLayerNum++ {
		streamSlice := (*encStream)[cLayerNum*layerCapacity : (cLayerNum+1)*layerCapacity]
		layers = append(layers, createSingleLayer(streamSlice, cLayerNum))
	}

	return layers
}

func createSingleLayer(slice []byte, id int) encLayer {
	layer := encLayer{layerID: id}
	sliceIndex := 0

	for i := 0; i < LAYER_HEIGHT; i++ {
		for j := 0; j < LAYER_WIDTH; j++ {
			layer.layerInfo[i][j] = toInt(string(slice[sliceIndex]))
			sliceIndex++
		}
	}
	return layer
}

func toInt(stringForm string) int {
	//helper since struct initializers don't do multi-values
	result, _ := strconv.Atoi(stringForm)
	return result
}

func getEncodedStreamFromInput(path string) []byte {
	stream, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	return bytes.TrimSpace(stream)
}
