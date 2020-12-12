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

	shipPosition := coord{0, 0}
	waypt := coord{10, 1}

	fmt.Println(shipPosition, waypt)
	for _, in := range instrs {
		shipPosition, waypt = moveShip(shipPosition, waypt, in)
		fmt.Println(shipPosition, waypt)
	}

	var res int
	res = abs(shipPosition.x) + abs(shipPosition.y)

	fmt.Println(res)
}

func moveShip(sPos coord, wayPt coord, i instr) (coord, coord) {
	act := i.action
	val := i.val
	if act == "F" {
		pos := sPos
		for j := 1; j <= val; j++ {
			pos.x += wayPt.x
			pos.y += wayPt.y
		}
		return pos, wayPt
	}

	if act == "L" {
		act = "R"
		val = convertLTurnToRTurn(val)
	}

	switch act {
	case "N":
		return sPos, coord{x: wayPt.x, y: wayPt.y + val}
	case "S":
		return sPos, coord{x: wayPt.x, y: wayPt.y - val}
	case "E":
		return sPos, coord{x: wayPt.x + val, y: wayPt.y}
	case "W":
		return sPos, coord{x: wayPt.x - val, y: wayPt.y}
	case "R":
		switch val % 360 {
		case 0:
			return sPos, wayPt
		case 90:
			return sPos, coord{x: wayPt.y, y: -wayPt.x}
		case 180:
			return sPos, coord{x: -wayPt.x, y: -wayPt.y}
		case 270:
			return sPos, coord{x: -wayPt.y, y: wayPt.x}
		}
	}
	return sPos, wayPt
}

func convertLTurnToRTurn(val int) int {
	switch val % 360 {
	case 90:
		return 270
	case 270:
		return 90
	default:
		return val
	}
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
