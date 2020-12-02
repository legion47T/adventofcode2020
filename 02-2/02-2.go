package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("test02-2.txt")
	// file, err := os.Open("input02-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]*)`)
	var res int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		tkns := r.FindStringSubmatch(line)
		// fmt.Println(tkns)

		lo, _ := strconv.Atoi(tkns[1])
		up, _ := strconv.Atoi(tkns[2])
		chall := []rune(tkns[3])
		test := []rune(tkns[4])
		if (test[lo-1] == chall[0]) != (test[up-1] == chall[0]) {
			fmt.Println(line)
			res++
		}
	}
	fmt.Println("Result: ", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
