package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"sync"
)

type sMap struct {
	mut *sync.RWMutex
	m   map[int]uint64
}

var nums []int
var outMap map[int][]int
var sumMap sMap

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
	mut := sync.RWMutex{}
	sumMap = sMap{
		mut: &mut,
		m:   make(map[int]uint64),
	}

	sumMap.m[nums[len(nums)-1]] = 1

	ch := make(chan uint64)
	go findPath(0, nums[len(nums)-1], ch)
	res = <-ch

	fmt.Println(res)
}

func findPath(curr int, target int, ch chan uint64) {
	// if curr == target {
	// 	// fmt.Println("End")
	// 	sumMap.mut.Lock()
	// 	sumMap.m[curr] = 1
	// 	sumMap.mut.Unlock()
	// 	ch <- 1
	// }

	var sum uint64
	chs := make([]chan uint64, 0)
	for _, out := range outMap[curr] {
		sumMap.mut.RLock()
		knownSum, ok := sumMap.m[out]
		sumMap.mut.RUnlock()
		if ok {
			sum += knownSum
			continue
		}
		subCh := make(chan uint64)
		chs = append(chs, subCh)
		go findPath(out, target, subCh)
	}
	// fmt.Println("curr", curr, "outMap[curr]", outMap[curr], "sum", sum)
	for _, subCh := range chs {
		sum += <-subCh
	}
	sumMap.mut.Lock()
	sumMap.m[curr] = sum
	sumMap.mut.Unlock()
	ch <- sum
}
