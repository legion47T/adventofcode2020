package main

import (
	"bufio"
	"container/ring"
	"log"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("../test23.txt")
	file, err := os.Open("../input23.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// numMoves := 10
	// numMoves := 100
	// numMoves := 10000
	numMoves := 10000000

	initialCups := make([]int, 0)
	maxVal := -1
	for sc.Scan() {
		line := sc.Text()
		for _, sym := range line {
			val, _ := strconv.Atoi(string(sym))
			if val > maxVal {
				maxVal = val
			}
			initialCups = append(initialCups, val)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= (1000000 - 9); i++ {
		initialCups = append(initialCups, maxVal+i)
	}
	numCups := len(initialCups)

	// adapted mnml's solution, thanks for putting it online! https://github.com/mnml/aoc/blob/master/2020/23/2.go

	cups := ring.New(numCups)
	cupInRing := map[int]*ring.Ring{}

	for i := 1; i <= numCups; i++ {
		if cups.Value = i; i <= len(initialCups) {
			cups.Value = initialCups[i-1]
		}
		cupInRing[cups.Value.(int)] = cups
		cups = cups.Next()
	}

	for i := 0; i < numMoves; i++ {
		takeAways := cups.Unlink(3)
		destCup := (cups.Value.(int)-2+numCups)%numCups + 1

		unavail := map[int]bool{}
		for i := 0; i < 3; i++ {
			unavail[takeAways.Value.(int)] = true
			takeAways = takeAways.Next()
		}

		for unavail[destCup] {
			destCup = (numCups+destCup-2+numCups)%numCups + 1
		}

		cupInRing[destCup].Link(takeAways)
		cups = cups.Next()
	}

	log.Println(cupInRing[1].Next().Value.(int) * cupInRing[1].Move(2).Value.(int))
}
