package main

import (
	"bufio"
	"log"
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

	var t uint64 = 100000000000000
	tick := uint64(busses[0])
	for {
		found := true
		for dt, bus := range busses {
			if dt == 0 {
				continue
			}
			if (t+uint64(dt))%uint64(bus) != 0 {
				found = false
				break
			}
		}
		if found {
			break
		}
		if t%10000000 == 0 {
			log.Println(t)
		}
		t += tick
	}

	res := t

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
