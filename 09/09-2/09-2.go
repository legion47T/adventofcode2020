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
	var target int

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for i := bracketLen + 1; i < len(nums); i++ {
		if !findTarget(&nums, i) {
			target = nums[i]
			break
		}
	}
	lo, hi := findBounds(&nums, target)

	res = nums[lo] + nums[hi]

	fmt.Println("Result:", res)
}

func findBounds(nums *[]int, target int) (lo int, hi int) {
	for i := range *nums {
		for j := i + 1; j < len(*nums); j++ {
			sum := sumUp(nums, i, j)
			if sum == target {
				return i, j
			}
			if sum > target {
				break
			}
		}
	}
	return -1, -1
}

func sumUp(nums *[]int, lo int, hi int) int {
	var sum int
	for i := lo; i <= hi; i++ {
		sum += (*nums)[i]
	}
	return sum
}

func findTarget(nums *[]int, idx int) bool {
	for i := idx - bracketLen; i < idx-1; i++ {
		for j := idx - bracketLen + 1; j < idx; j++ {
			if (*nums)[i]+(*nums)[j] == (*nums)[idx] {
				return true
			}
		}
	}
	return false
}
