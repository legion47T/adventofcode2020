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
	// file, err := os.Open("test16-2.txt")
	file, err := os.Open("../input16.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	constrRegex = regexp.MustCompile(`(.*): (\d*)-(\d*) or (\d*)-(\d*)`)

	constrs := make([]constr, 0)
	var myTicket []int
	validTickets := make([][]int, 0)
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
			myTicket = parseTicket(line)
		case 2:
			ticket := parseTicket(line)
			if ticketValid(&ticket, &constrs) {
				validTickets = append(validTickets, ticket)
			}
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(myTicket)

	candidates := make([]map[string]bool, len(constrs))
	for _, ticket := range validTickets {
		for i, val := range ticket {
			column := candidates[i]
			cons := validConstrs(val, &constrs)
			if len(column) == 0 {
				candidates[i] = cons
				continue
			}
			for fieldName, possible := range cons {
				if column[fieldName] != possible {
					column[fieldName] = false
				} else {
					column[fieldName] = possible
				}
			}
		}
	}

	for !assignmentDefinite(&candidates) {
		for i, row := range candidates {
			if definiteField, definite := rowDefinite(&row); definite {
				clearOtherRows(i, definiteField, &candidates)
			}
		}
	}

	fields := make([]string, 0)
	for idx, column := range candidates {
		log.Println(idx, column)
		for fieldName, possible := range column {
			if possible {
				fields = append(fields, fieldName)
			}
		}
	}
	log.Println(fields)

	var res int = 1
	for i, fieldName := range fields {
		if strings.HasPrefix(fieldName, "departure") {
			log.Println(fieldName, myTicket[i])
			res *= myTicket[i]
		}
	}

	log.Println(res)
}

func clearOtherRows(idx int, definiteField string, candidates *[]map[string]bool) {
	for i, row := range *candidates {
		if idx != i {
			row[definiteField] = false
		}
	}
}

func rowDefinite(row *map[string]bool) (string, bool) {
	foundOne := false
	var foundFieldName string
	for fieldName, possible := range *row {
		if possible {
			if foundOne {
				return "", false
			}
			foundOne = true
			foundFieldName = fieldName
		}
	}
	return foundFieldName, true
}

func assignmentDefinite(candidates *[]map[string]bool) bool {
	for _, row := range *candidates {
		if _, definite := rowDefinite(&row); !definite {
			return false
		}
	}
	return true
}

func validConstrs(val int, constrs *[]constr) map[string]bool {
	out := make(map[string]bool)
	for _, c := range *constrs {
		if constrValid(val, &c) {
			out[c.fieldName] = true
		} else {
			out[c.fieldName] = false
		}
	}
	return out
}
func ticketValid(ticket *[]int, constrs *[]constr) bool {
	for _, val := range *ticket {
		if !valueValid(val, constrs) {
			return false
		}
	}
	return true
}

func constrValid(val int, con *constr) bool {
	for _, r := range con.ranges {
		if val >= r.lo && val <= r.hi {
			return true
		}
	}
	return false
}

func valueValid(val int, constrs *[]constr) bool {
	for _, con := range *constrs {
		if constrValid(val, &con) {
			return true
		}
	}
	return false
}

func parseTicket(in string) []int {
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
	range1lo, _ := strconv.Atoi(submatches[2])
	range1hi, _ := strconv.Atoi(submatches[3])
	range2lo, _ := strconv.Atoi(submatches[4])
	range2hi, _ := strconv.Atoi(submatches[5])
	ranges := []numRange{
		{lo: range1lo, hi: range1hi},
		{lo: range2lo, hi: range2hi},
	}

	return constr{fieldName: name, ranges: ranges}
}
