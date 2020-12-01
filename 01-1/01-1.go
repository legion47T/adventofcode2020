package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// data, err := ioutil.ReadFile("test01-1.txt")
	data, err := ioutil.ReadFile("input01-1.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	numberstrings := strings.Split(string(data), "\n")
	fmt.Println(numberstrings)
	var numbers []int
	for _, s := range numberstrings {
		out, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Can't parse string:", s)
			return
		}
		numbers = append(numbers, out)
	}
	for i, n := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if n+numbers[j] == 2020 {
				fmt.Println("Result:", n*numbers[j])
				return
			}
		}
	}
}
