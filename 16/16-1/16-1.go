package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type numRange struct {
	lo, hi int
}

type constr struct {
	fieldName string
	ranges    []numRange
}

var constrRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test16.txt")
	file, err := os.Open("../input16.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	constrRegex = regexp.MustCompile(`(\w*): (\d*-\d*) or (\d*-\d*)`)

	constrs := make([]constr, 0)
	var myTicket []int
	otherTickets := make([][]int, 0)
	mode := 0
	for sc.Scan() {
		line := sc.Text()
		switch line {
		case "":
			continue
		case "your ticket:":
			mode = 1
			continue
		case "nearby tickets:":
			mode = 2
			continue
		}
		switch mode {
		case 0:
			constrs = append(constrs, parseConstr(line))
		case 1:
			myTicket = parseTickets(line)
		case 2:
			otherTickets = append(otherTickets, parseTickets(line))
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(myTicket)

	var res int64
	for _, ticket := range otherTickets {
		for _, val := range ticket {
			if !valueValid(val, &constrs) {
				res += int64(val)
			}
		}
	}

	log.Println(res)
}

func valueValid(val int, constrs *[]constr) bool {
	for _, con := range *constrs {
		for _, r := range con.ranges {
			if val >= r.lo && val <= r.hi {
				return true
			}
		}
	}
	return false
}

func parseTickets(in string) []int {
	ticket := make([]int, 0)
	vals := strings.Split(in, ",")
	for _, val := range vals {
		valNum, _ := strconv.Atoi(val)
		ticket = append(ticket, valNum)
	}
	return ticket
}

func parseConstr(in string) constr {
	submatches := constrRegex.FindStringSubmatch(in)
	name := submatches[1]
	range1 := strings.Split(submatches[2], "-")
	range2 := strings.Split(submatches[3], "-")
	range1lo, _ := strconv.Atoi(range1[0])
	range1hi, _ := strconv.Atoi(range1[1])
	range2lo, _ := strconv.Atoi(range2[0])
	range2hi, _ := strconv.Atoi(range2[1])
	ranges := []numRange{
		{lo: range1lo, hi: range1hi},
		{lo: range2lo, hi: range2hi},
	}

	return constr{fieldName: name, ranges: ranges}
}
