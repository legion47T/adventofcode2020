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

func main() {
	// file, err := os.Open("../test08.txt")
	file, err := os.Open("../input08.txt")

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
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(instrs)
	nextIdx := 0
	// lastIdcs := make([]int, 0)
	// var idxMod int
	var acc int
	resCh := make(chan int)
	go execute(&instrs, acc, true, nextIdx, resCh)
	// for err == nil && nextIdx < len(instrs) {
	// 	// fmt.Println("NextIdx", nextIdx)
	// 	lastIdcs = append(lastIdcs, nextIdx)
	// 	var accMod int
	// 	idxMod, accMod, err = executeInstr(&instrs[nextIdx], false)
	// 	acc += accMod
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	// fmt.Println("IdxMod", idxMod, "error", err)
	// 	nextIdx += idxMod
	// }
	// fmt.Println("Last Index", lastIdcs)
	acc = <-resCh

	fmt.Println("Result:", acc)
}

func execute(instrs *[]instruction, acc int, mayBranch bool, nextIdx int, resCh chan<- int) {
	// log.Println("Instrs:", instrs, "acc:", acc, "mayBranch", mayBranch, "nextIdx", nextIdx)
	defer close(resCh)
	var err error
	var accMod int
	var idxMod int
	copyInstrs := make([]instruction, len(*instrs))
	copy(copyInstrs, *instrs)

	myCh := make(chan int)
	for err == nil && nextIdx < len(*instrs) {
		copy(copyInstrs, *instrs)
		if mayBranch && (*instrs)[nextIdx].instr != "acc" {
			idxMod, _, _ := executeInstr(&copyInstrs[nextIdx], true)
			go execute(&copyInstrs, acc, false, nextIdx+idxMod, myCh)
		}
		idxMod, accMod, err = executeInstr(&(*instrs)[nextIdx], false)
		acc += accMod
		nextIdx += idxMod
	}

	if nextIdx == len(*instrs) {
		log.Println("Finished!", acc)
		resCh <- acc
		return
	}
	// if err != nil {
	// 	log.Println("Error", err)
	// 	return
	// }

	acc = -1
	log.Println("No own result")
	acc, ok := <-myCh
	if ok {
		log.Println("Received", acc, "from branch")
		resCh <- acc
	}
	// for (*instrs)[nextIdx].instr == "acc" && err == nil && nextIdx < len(*instrs) {
	// 	copy(copyInstrs, *instrs)
	// 	idxMod, accMod, err = executeInstr(&(*instrs)[nextIdx], false)
	// 	acc += accMod
	// 	// fmt.Println("IdxMod", idxMod, "error", err)
	// 	nextIdx += idxMod
	// }

	// if nextIdx == len(*instrs) {
	// 	log.Println("Finished 1!", acc)
	// 	resCh <- acc
	// 	return
	// }
	// if err != nil {
	// 	log.Println("Error", err)
	// 	return
	// }

	// myCh := make(chan int)
	// if mayBranch {
	// 	idxMod, _, _ := executeInstr(&copyInstrs[nextIdx], true)
	// 	go execute(&copyInstrs, acc, false, nextIdx+idxMod, myCh)
	// }
	// for err == nil && nextIdx < len(*instrs) {
	// 	idxMod, accMod, err = executeInstr(&(*instrs)[nextIdx], false)
	// 	acc += accMod
	// 	// fmt.Println("IdxMod", idxMod, "error", err)
	// 	nextIdx += idxMod
	// }
	// if nextIdx == len(*instrs) {
	// 	log.Println("Finished 2!", acc)
	// 	resCh <- acc
	// 	return
	// }
	// if mayBranch {
	// 	acc = -1
	// 	log.Println("No own result")
	// 	acc, ok := <-myCh
	// 	if ok {
	// 		log.Println("Received", acc, "from branch")
	// 		resCh <- acc
	// 	}
	// }
}

func executeInstr(instr *instruction, doBranch bool) (idxMod int, accMod int, err error) {
	if instr.called {
		return 0, 0, errors.New("Loop detected")
	}
	if !doBranch {
		instr.called = true
	}
	switch instr.instr {
	case "nop":
		if !doBranch {
			return +1, 0, nil
		}
		return instr.operand, 0, nil
	case "acc":
		return +1, instr.operand, nil
	case "jmp":
		if !doBranch {
			return instr.operand, 0, nil
		}
		return +1, 0, nil
	}
	return 0, 0, errors.New("unreachable")
}

func parseInstr(instr string) instruction {
	split := strings.Split(instr, " ")
	num, _ := strconv.Atoi(split[1])
	return instruction{
		instr:   split[0],
		operand: num,
	}
}
