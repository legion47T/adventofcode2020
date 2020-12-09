package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	instr   string
	operand int
	called  bool
}

var acc, bracketLen int

func main() {
	// file, err := os.Open("../test09.txt")
	// bracketLen = 5
	file, err := os.Open("../input09.txt")
	bracketLen = 25

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int
	nums := make([]int, 0)

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	for i := bracketLen + 1; i < len(nums); i++ {
		if !findSum(&nums, i) {
			res = nums[i]
		}
	}

	fmt.Println("Result:", res)
}

func findSum(nums *[]int, idx int) bool {
	for i := idx - bracketLen; i < idx-1; i++ {
		for j := idx - bracketLen + 1; j < idx; j++ {
			if (*nums)[i]+(*nums)[j] == (*nums)[idx] {
				return true
			}
		}
	}
	return false
}
