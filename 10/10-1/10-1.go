package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type instruction struct {
	instr   string
	operand int
	called  bool
}

func main() {
	// file, err := os.Open("../test10a.txt")
	// file, err := os.Open("../test10b.txt")
	file, err := os.Open("../input10.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int
	nums := make([]int, 1)

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		num, _ := strconv.Atoi(sc.Text())
		nums = append(nums, num)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	sort.IntSlice(nums).Sort()
	nums = append(nums, nums[len(nums)-1]+3)
	fmt.Println(nums)

	var ones, twos, threes int
	for i := 1; i < len(nums); i++ {
		switch nums[i] - nums[i-1] {
		case 1:
			ones++
		case 2:
			twos++
		case 3:
			threes++
		}
	}
	res = ones * threes

	fmt.Println(res)
}
