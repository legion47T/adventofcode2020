package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("test02-1.txt")
	file, err := os.Open("input02-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]*)`)
	var res int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		tkns := r.FindStringSubmatch(line)
		fmt.Println(tkns)

		lo, _ := strconv.Atoi(tkns[1])
		up, _ := strconv.Atoi(tkns[2])
		chall := tkns[3]
		test := tkns[4]
		num := strings.Count(test, chall)
		if num >= lo && num <= up {
			res++
		}
	}
	fmt.Println("Result: ", res)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
