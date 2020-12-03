package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("test03-1.txt")
	file, err := os.Open("input03-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var landscape [][]rune
	var res int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		landscape = append(landscape, []rune(scanner.Text()))
	}

	// fmt.Println(landscape)

	for i, level := range landscape {
		if i == 0 {
			continue
		}
		if level[(i*3)%len(level)] == '#' {
			res++
		}
	}
	fmt.Println("Result: ", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
