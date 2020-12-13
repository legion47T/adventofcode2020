package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instr struct {
	action string
	val    int
}

type coord struct {
	x, y int
}

var instrRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test13.txt")
	file, err := os.Open("../input13.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	line := 0
	var arrival int
	var busses []int
	for sc.Scan() {
		if line == 0 {
			arrival, _ = strconv.Atoi(sc.Text())
			line++
			continue
		}
		busses = parseBusses(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	waitingTimes := make([]int, 0)
	for _, busID := range busses {
		waitingTime := busID - arrival%busID
		waitingTimes = append(waitingTimes, waitingTime)
	}

	var res int

	shortestTime := math.MaxInt16
	for i, waitingTime := range waitingTimes {
		if waitingTime < shortestTime {
			shortestTime = waitingTime
			res = shortestTime * busses[i]
		}
	}

	fmt.Println(res)
}

func parseBusses(in string) []int {
	busIds := strings.Split(in, ",")
	busses := make([]int, 0)
	for _, bus := range busIds {
		if bus == "x" {
			continue
		}
		id, _ := strconv.Atoi(bus)
		busses = append(busses, id)
	}
	return busses
}
