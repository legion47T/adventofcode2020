package main

import (
	"log"
	"strconv"
	"strings"
)

type spoken struct {
	lastTurn     int
	previousTurn int
	new          bool
}

func main() {
	// inp := "0,3,6" // 436
	// inp := "1,3,2" // 1
	inp := "6,19,0,5,7,13,1"

	lastTurnByNumber := make(map[int]spoken)
	numbers := make([]int, 0, 2020)
	for i, num := range strings.Split(inp, ",") {
		numVal, _ := strconv.Atoi(num)
		numbers = append(numbers, numVal)
		lastTurnByNumber[numVal] = spoken{lastTurn: i, previousTurn: i, new: true}
	}

	for i := len(numbers); i < 2020; i++ {
		numSpoken, _ := lastTurnByNumber[numbers[i-1]]
		if numSpoken.lastTurn == 0 {
			numbers = append(numbers, 0)
			lastTurnByNumber[0] = spoken{i, numSpoken.previousTurn, false}
			continue
		}
		num := numSpoken.lastTurn - numSpoken.previousTurn
		numbers = append(numbers, num)
		if newNum, ok := lastTurnByNumber[num]; ok {
			lastTurnByNumber[num] = spoken{i, newNum.lastTurn, false}
		} else {
			lastTurnByNumber[num] = spoken{i, i, true}
		}
	}
	log.Println(numbers)

	res := numbers[2019]

	log.Println(res)
}
