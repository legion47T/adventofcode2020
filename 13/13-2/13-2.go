package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/v1/bigmod"
)

func main() {
	// file, err := os.Open("../test13.txt")
	file, err := os.Open("../input13.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	line := 0
	var busses map[int]int
	for sc.Scan() {
		if line == 0 {
			line++
			continue
		}
		busses = parseBusses(sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	crtEs := make([]bigmod.CrtEntry, 0, len(busses))

	for t, bus := range busses {
		crtEs = append(crtEs, bigmod.CrtEntry{A: big.NewInt(int64(-t)), N: big.NewInt(int64(bus))})
	}

	res := bigmod.SolveCrtMany(crtEs)

	log.Println(res)
}

func parseBusses(in string) map[int]int {
	busIds := strings.Split(in, ",")
	busses := make(map[int]int, 0)
	for i, bus := range busIds {
		if bus == "x" {
			continue
		}
		id, _ := strconv.Atoi(bus)
		busses[i] = id
	}
	return busses
}
