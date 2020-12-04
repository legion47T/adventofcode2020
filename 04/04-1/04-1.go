package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type tkns struct {
	byr bool
	iyr bool
	eyr bool
	hgt bool
	hcl bool
	ecl bool
	pid bool
	cid bool
}

func main() {
	// file, err := os.Open("../test04.txt")
	file, err := os.Open("../input04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	req := tkns{byr: true, ecl: true, eyr: true, hcl: true, hgt: true, iyr: true, pid: true, cid: false}
	var numValid int

	scanner := bufio.NewScanner(file)
	scanner.Split(doubleNLSplitFunc)
	for scanner.Scan() {
		record := strings.Fields(strings.ReplaceAll(scanner.Text(), "\n", " "))
		if isValid(req, findTkns(record)) {
			numValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", numValid)
}

func isValid(need tkns, has tkns) bool {
	return (!need.byr || has.byr) && (!need.iyr || has.iyr) && (!need.eyr || has.eyr) && (!need.hgt || has.hgt) && (!need.hcl || has.hcl) && (!need.ecl || has.ecl) && (!need.pid || has.pid) && (!need.cid || has.cid)
}

func findTkns(in []string) tkns {
	r := regexp.MustCompile(`(\w{3}):.*`)
	found := tkns{}
	for _, info := range in {
		tkn := r.FindStringSubmatch(info)[1]
		if tkn == "byr" {
			found.byr = true
		} else if tkn == "iyr" {
			found.iyr = true
		} else if tkn == "eyr" {
			found.eyr = true
		} else if tkn == "hgt" {
			found.hgt = true
		} else if tkn == "hcl" {
			found.hcl = true
		} else if tkn == "ecl" {
			found.ecl = true
		} else if tkn == "pid" {
			found.pid = true
		} else if tkn == "cid" {
			found.cid = true
		}
	}
	return found
}

func doubleNLSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a newline followed by a
	// pound sign.
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}
