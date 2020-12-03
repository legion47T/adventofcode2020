package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	x int
	y int
}

func main() {
	// file, err := os.Open("test03-2.txt")
	file, err := os.Open("input03-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var landscape [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		landscape = append(landscape, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var res uint64
	res = 1

	slopes := []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	// fmt.Println(landscape)
	for _, slope := range slopes {
		var score uint64
		fmt.Println("Slope", slope)
		var lastx int
		for i := slope.y; i < len(landscape); i += slope.y {
			level := landscape[i]
			lastx += slope.x
			// fmt.Println("LastX:", lastx)
			idx := lastx % len(level)
			fmt.Println("x:", idx, "y:", i)
			// fmt.Println("Index:", idx)
			if level[idx] == '#' {
				score++
			}
		}
		fmt.Println("Score:", score)
		res *= score
	}
	fmt.Println("Result: ", res)
}
