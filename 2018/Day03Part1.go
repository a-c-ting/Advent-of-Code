package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
	"strconv"
)

type elfClaim struct {
	elfClaimId	int
	fabricRow 	int //distance from left
	fabricCol	int //distance from top
	claimedRow	int //width
	claimedCol	int //height
}

func main() {
	elfClaimList := getElfClaimsData("Day03Input.txt")
	fabric := [1000][1000]int{}

	simulateFabricUsage(&fabric, elfClaimList)

	fmt.Printf("Total Overlap is %d\n", countOverlappingFabric(&fabric))
}

func countOverlappingFabric(fabric *[1000][1000]int) int {
	totalOverlaps := 0
	for _, columns := range fabric {
		for _, value := range columns {
			if value > 1 {
				totalOverlaps += 1
			}
		}
	}
	return totalOverlaps
}

func simulateFabricUsage(fabric *[1000][1000]int, claimList []elfClaim) {
	for _, claim := range claimList {
		for i := claim.fabricRow; i < (claim.fabricRow + claim.claimedRow); i++ {
			for j := claim.fabricCol; j < (claim.fabricCol + claim.claimedCol); j++ {
				fabric[i][j] += 1
			}
		}
	}
	return
}

func getElfClaimsData(path string) []elfClaim {
	stream, err := ioutil.ReadFile(path)
	if (err != nil) {
		return nil
	}

	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}
	cleanedStream := strings.FieldsFunc(string(stream), f)
	claimCount := len(cleanedStream)/5

	result := make([]elfClaim, claimCount)
	for i := 0; i < len(cleanedStream); i += 5 {
		v := elfClaim{
			elfClaimId: toInt(cleanedStream[i]),
			fabricRow: toInt(cleanedStream[i+1]),
			fabricCol: toInt(cleanedStream[i+2]),
			claimedRow: toInt(cleanedStream[i+3]),
			claimedCol: toInt(cleanedStream[i+4]),
		}
		result[i/5] = v
	}
	return result
}

func toInt(stringForm string) int {
	//helper since struct initializers don't do multi-values
	result, _ := strconv.Atoi(stringForm)
	return result
}
