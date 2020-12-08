package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	instr   string
	operand int
	called  bool
}

var acc int

func main() {
	file, err := os.Open("../test08.txt")
	// file, err := os.Open("../input08.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var res int64

	sc := bufio.NewScanner(file)
	instrs := make([]instruction, 0)
	for sc.Scan() {
		instrs = append(instrs, parseInstr(sc.Text()))
	}
	// fmt.Println(instrs)
	nextIdx := 0
	lastIdcs := make([]int, 0)
	var idxMod int
	for err == nil && nextIdx < len(instrs) {
		// fmt.Println("NextIdx", nextIdx)
		lastIdcs = append(lastIdcs, nextIdx)
		idxMod, err = executeInstrs(&instrs[nextIdx])
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println("IdxMod", idxMod, "error", err)
		nextIdx += idxMod
	}
	fmt.Println("Last Index", lastIdcs)

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", acc)
}

func executeInstrs(instr *instruction) (idxMod int, err error) {
	if instr.called {
		return 0, errors.New("Loop detected")
	}
	switch instr.instr {
	case "nop":
		instr.called = true
		return +1, nil
	case "acc":
		instr.called = true
		acc += instr.operand
		return +1, nil
	case "jmp":
		instr.called = true
		return instr.operand, nil
	}
	return 0, errors.New("unreachable")
}

func parseInstr(instr string) instruction {
	split := strings.Split(instr, " ")
	num, _ := strconv.Atoi(split[1])
	return instruction{
		instr:   split[0],
		operand: num,
	}
}
