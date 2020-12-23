package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("../test23.txt")
	file, err := os.Open("../input23.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	// numMoves := 10
	numMoves := 100

	cups := make([]int, 0)
	for sc.Scan() {
		line := sc.Text()
		for _, sym := range line {
			val, _ := strconv.Atoi(string(sym))
			cups = append(cups, val)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	var currentCupIdx int
	numCups := len(cups)
	for i := 1; i <= numMoves; i++ {
		fmt.Println("-- move", i, "--")
		fmt.Println(cups)
		currentCup := cups[currentCupIdx]
		fmt.Println("current cup:", currentCup)
		takeAways := make([]int, 0)
		for j := 1; j <= 3; j++ {
			idx := (find(cups, currentCup) + 1) % len(cups)
			takeAways = append(takeAways, cups[idx])
			cups = remove(cups, idx)
		}
		fmt.Println("pick up:", takeAways)
		destCup := currentCup - 1
		if destCup < 1 {
			destCup = numCups
		}
		destIdx := find(cups, destCup)
		for destIdx == -1 {
			destCup--
			if destCup < 1 {
				destCup = numCups
			}
			destIdx = find(cups, destCup)
		}
		cups = insert(cups, takeAways, destIdx+1)
		fmt.Println("destination:", destCup)
		currentCupIdx = (find(cups, currentCup) + 1) % numCups
	}
	fmt.Println(cups)

	oneIdx := find(cups, 1)
	for i := 0; i < numCups; i++ {
		idx := (i + oneIdx) % numCups
		fmt.Print(cups[idx])
	}

	// var res int

	// log.Println(res)
}

func insert(a []int, c []int, i int) []int {
	if i < len(a) {
		return append(a[:i], append(c, a[i:]...)...)
	}
	return append(a, c...)
}

func find(slice []int, val int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return i
		}
	}
	return -1
}

func reverse(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func unique(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	// We do not need to put s[i] at the end, as it will be discarded anyway
// 	return s[:len(s)-1]
// }

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
