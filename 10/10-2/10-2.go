package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var nums []int
var outMap map[int][]int
var sumMap map[int]uint64

func main() {
	// file, err := os.Open("../test10a.txt")
	// file, err := os.Open("../test10b.txt")
	file, err := os.Open("../input10.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res uint64
	nums = make([]int, 1)

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

	outMap = make(map[int][]int)
	maxIdx := len(nums) - 1
	for i, jolt := range nums {
		upper := 3
		if maxIdx-i < upper {
			upper = maxIdx - i
		}
		for j := 1; j <= upper; j++ {
			desc := nums[i+j]
			if desc-jolt <= 3 {
				outMap[jolt] = append(outMap[jolt], desc)
			}
		}
	}

	fmt.Println(outMap)

	sumMap = make(map[int]uint64)

	res = findPath(0, nums[len(nums)-1])

	fmt.Println(res)
}

func findPath(curr int, target int) uint64 {
	if curr == target {
		// fmt.Println("End")
		sumMap[curr] = 1
		return 1
	}

	var sum uint64
	for _, out := range outMap[curr] {
		if knownSum, ok := sumMap[out]; ok {
			sum += knownSum
			continue
		}
		sum += findPath(out, target)
	}
	// fmt.Println("curr", curr, "outMap[curr]", outMap[curr], "sum", sum)
	sumMap[curr] = sum
	return sum
}
