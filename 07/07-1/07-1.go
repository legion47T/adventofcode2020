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

func main() {
	// file, err := os.Open("../test07.txt")
	file, err := os.Open("../input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// var res int

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

	allBags.RLock()
	shiny := allBags.m["shiny gold"]
	allBags.RUnlock()
	fmt.Println(shiny)

	chs := make([]chan string, 0)
	for _, color := range shiny.mayBeIn.s {
		colorCh := make(chan string)
		chs = append(chs, colorCh)
		go findMayBeIns(color, colorCh)
	}

	foundlings := make([]string, 0)
	for {
		done := 0
		for _, ch := range chs {
			select {
			case answ, ok := <-ch:
				if ok {
					fmt.Println("found:", answ)
					foundlings = append(foundlings, answ)
				} else {
					done++
				}
			default:
			}
		}
		if done == len(chs) {
			break
		}
	}
	foundlings = unique(foundlings)

	if err := groupScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", len(foundlings))
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
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
	r := regexp.MustCompile(`(.*) bags contain (.*)\.`)
	containedBags := make(map[string]int)
	submatches := r.FindStringSubmatch(sentence)
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
	r := regexp.MustCompile(`(\d*) (.*) bags?`)
	submatches := r.FindStringSubmatch(numAndColor)
	num, _ := strconv.Atoi(submatches[1])
	return num, submatches[2]
}
