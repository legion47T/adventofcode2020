package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instr struct {
	action string
	val    int
}

type ship struct {
	heading  int
	position coord
}

type coord struct {
	x, y int
}

var instrRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test12.txt")
	file, err := os.Open("../input12.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	instrRegex = regexp.MustCompile(`(\w)(\d*)`)

	instrs := make([]instr, 0)
	for sc.Scan() {
		instrs = append(instrs, parseInstruction(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(instrs)

	s := ship{90, coord{0, 0}}

	// fmt.Println(s)
	for _, in := range instrs {
		s = moveShip(s, in)
		// fmt.Println(s)
	}

	var res int
	res = abs(s.position.x) + abs(s.position.y)

	fmt.Println(res)
}

func moveShip(s ship, i instr) ship {
	act := i.action
	if act == "F" {
		act = parseDirection(s.heading)
	}
	switch act {
	case "N":
		return ship{heading: s.heading, position: coord{x: s.position.x, y: s.position.y + i.val}}
	case "S":
		return ship{heading: s.heading, position: coord{x: s.position.x, y: s.position.y - i.val}}
	case "E":
		return ship{heading: s.heading, position: coord{x: s.position.x + i.val, y: s.position.y}}
	case "W":
		return ship{heading: s.heading, position: coord{x: s.position.x - i.val, y: s.position.y}}
	case "L":
		return ship{heading: (s.heading - i.val) % 360, position: coord{x: s.position.x, y: s.position.y}}
	case "R":
		return ship{heading: (s.heading + i.val) % 360, position: coord{x: s.position.x, y: s.position.y}}
	}
	return s
}

func parseDirection(dir int) string {
	neg := isNeg(dir)
	switch abs(dir % 360) {
	case 0:
		return "N"
	case 90:
		if neg {
			return "W"
		}
		return "E"
	case 180:
		return "S"
	case 270:
		if neg {
			return "E"
		}
		return "W"
	}
	return "X"
}

func isNeg(x int) bool {
	if x < 0 {
		return true
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInstruction(in string) instr {
	submatches := instrRegex.FindStringSubmatch(in)
	val, _ := strconv.Atoi(submatches[2])
	return instr{action: submatches[1], val: val}
}
