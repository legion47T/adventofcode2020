package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("../test18.txt")
	file, err := os.Open("../input18.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var res int
	for sc.Scan() {
		val, _ := parseCalculation(sc.Text())
		// fmt.Println(val)
		res += val
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func parseCalculation(in string) (int, int) {
	var result int
	var operator rune
	var advanceTo int
	for i, symbol := range in {
		if i < advanceTo {
			continue
		}
		switch {
		case symbol == ' ':
			continue
		case symbol >= '1' && symbol <= '9':
			val, _ := strconv.Atoi(string(symbol))
			if result == 0 {
				result = val
			} else {
				switch operator {
				case '+':
					result += val
				case '*':
					result *= val
				}
			}
		case symbol == '+':
			operator = '+'
			continue
		case symbol == '*':
			operator = '*'
			continue
		case symbol == '(':
			val, adv := parseCalculation(in[i+1:])
			advanceTo = adv + i + 1
			if result == 0 {
				result = val
			} else {
				switch operator {
				case '+':
					result += val
				case '*':
					result *= val
				}
			}
		case symbol == ')':
			return result, i + 1
		}
	}
	return result, 0
}
