package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type bag struct {
	color         string
	containedBags map[string]int
	mayBeIn       *concurrentSlice
}

type concurrentSlice struct {
	*sync.RWMutex
	s []string
}

type bags struct {
	*sync.RWMutex
	m map[string]bag
}

var allBags bags

var bagsContainRegex *regexp.Regexp
var numBagSplitRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test07.txt")
	// file, err := os.Open("test07-2.txt")
	file, err := os.Open("../input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int64

	bagsContainRegex = regexp.MustCompile(`(.*) bags contain (.*)\.`)
	numBagSplitRegex = regexp.MustCompile(`(\d*) (.*) bags?`)

	allBags = bags{&sync.RWMutex{}, make(map[string]bag)}
	wg := sync.WaitGroup{}
	groupScanner := bufio.NewScanner(file)
	for groupScanner.Scan() {
		sentence := groupScanner.Text()
		// sentence := "light red bags contain 1 bright white bag, 2 muted yellow bags."
		wg.Add(1)
		go createBag(sentence, &wg)
	}
	wg.Wait()

	allBags.RLock()
	for _, aBag := range allBags.m {
		for containedBag := range aBag.containedBags {
			wg.Add(1)
			go addMaybeIn(aBag.color, containedBag, &wg)
		}
	}
	allBags.RUnlock()
	wg.Wait()

	resCh := make(chan int64)
	shiny := allBags.m["shiny gold"]
	fmt.Println(shiny)
	go countContdBags("shiny gold", resCh)
	// vibrant := allBags.m["vibrant plum"]
	// fmt.Println(vibrant)
	// go countContdBags("vibrant plum", resCh)

	for sub := range resCh {
		res += sub
	}

	if err := groupScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", res)
}

func countContdBags(color string, resultChan chan<- int64) {
	defer close(resultChan)
	resChs := make(map[chan int64]int)
	contdBags := allBags.m[color].containedBags
	if len(contdBags) > 0 {
		for containedColor, times := range contdBags {
			// fmt.Println(color, "contains", containedColor, times, "times")
			resCh := make(chan int64)
			resChs[resCh] = times
			go countContdBags(containedColor, resCh)
		}
		var sum int64
		for resCh, times := range resChs {
			for res := range resCh {
				sum += (1 + res) * int64(times)
				// fmt.Println(color, ":", times, "->", res)
			}
			// fmt.Println(color, ":", times, "times", sum, "makes", (sum * int64(times)))
		}
		resultChan <- sum
	} else {
		// fmt.Println(color, "contains no other bags")
		resultChan <- 0
	}
}

func addMaybeIn(containingColor string, containedColor string, wg *sync.WaitGroup) {
	// fmt.Println("Add", containingColor, "to", containedColor)
	defer wg.Done()
	allBags.RLock()
	containedBag, _ := allBags.m[containedColor]
	allBags.RUnlock()
	containedBag.mayBeIn.Lock()
	// fmt.Println("Old content of mayBeIn of color", containedColor, ":", containedBag.mayBeIn.s)
	containedBag.mayBeIn.s = append(containedBag.mayBeIn.s, containingColor)
	// fmt.Println("New content of mayBeIn of color", containedColor, ":", containedBag.mayBeIn.s)
	containedBag.mayBeIn.Unlock()
}

func findMayBeIns(color string, ch chan string) {
	chs := make([]chan string, 0)
	for _, sub := range allBags.m[color].mayBeIn.s {
		colorCh := make(chan string)
		chs = append(chs, colorCh)
		go findMayBeIns(sub, colorCh)
	}
	for {
		var dones int
		for _, subCh := range chs {
			select {
			case answ, ok := <-subCh:
				if ok {
					ch <- answ
				} else {
					dones++
				}
			default:
			}
		}
		if dones == len(chs) {
			break
		}
	}
	ch <- color
	close(ch)
}

func createBag(sentence string, wg *sync.WaitGroup) {
	defer wg.Done()
	containedBags := make(map[string]int)
	submatches := bagsContainRegex.FindStringSubmatch(sentence)
	color := submatches[1]
	subsentence := submatches[2]
	if subsentence != "no other bags" {
		for _, part := range strings.Split(subsentence, ",") {
			num, color := findNumAndColor(strings.TrimSpace(part))
			containedBags[color] = num
		}
	}
	wg.Add(1)
	go addOrUpdateBag(bag{
		color:         color,
		containedBags: containedBags,
		mayBeIn:       &concurrentSlice{&sync.RWMutex{}, make([]string, 0)},
	}, wg)
	// wg.Add(1)
	// go updateContainingBags(color, containedBags, wg)
}

func addOrUpdateBag(newBag bag, wg *sync.WaitGroup) {
	defer wg.Done()
	// allBags.RLock()
	// foundBag, ok := allBags.m[newBag.color]
	// allBags.RUnlock()
	// if ok {
	// 	foundBag.containedBags = newBag.containedBags
	// 	return
	// }
	allBags.Lock()
	allBags.m[newBag.color] = newBag
	allBags.Unlock()
}

// func updateContainingBags(color string, containedBags map[string]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for containedColor := range containedBags {
// 		allBags.RLock()
// 		foundBag, ok := allBags.m[containedColor]
// 		allBags.RUnlock()
// 		if !ok {
// 			// newBag := bag{
// 			// 	color: containedColor,
// 			// 	mayBeIn: concurrentSlice{
// 			// 		&sync.RWMutex{},
// 			// 		[]string{color},
// 			// 	},
// 			// }
// 			// wg.Add(1)
// 			// go addOrUpdateBag(newBag, wg)
// 		} else {
// 			foundBag.mayBeIn.RLock()
// 			found := false
// 			for _, containingColor := range foundBag.mayBeIn.s {
// 				if color == containingColor {
// 					found = true
// 				}
// 			}
// 			foundBag.mayBeIn.RUnlock()
// 			if !found {
// 				foundBag.mayBeIn.Lock()
// 				foundBag.mayBeIn.s = append(foundBag.mayBeIn.s, color)
// 				foundBag.mayBeIn.Unlock()
// 			}
// 		}
// 	}
// }

func findNumAndColor(numAndColor string) (int, string) {
	submatches := numBagSplitRegex.FindStringSubmatch(numAndColor)
	num, _ := strconv.Atoi(submatches[1])
	return num, submatches[2]
}
