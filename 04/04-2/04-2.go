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

type reqFields struct {
	byr bool
	iyr bool
	eyr bool
	hgt bool
	hcl bool
	ecl bool
	pid bool
}

func main() {
	// file, err := os.Open("test04-2_valid.txt")
	// file, err := os.Open("test04-2_invalid.txt")
	// file, err := os.Open("test04-2.txt")
	file, err := os.Open("../input04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numValid int

	scanner := bufio.NewScanner(file)
	scanner.Split(doubleNLSplitFunc)
	for scanner.Scan() {
		record := strings.Fields(strings.ReplaceAll(scanner.Text(), "\n", " "))
		tkns := mapFromList(record)
		if isValid(tkns) {
			numValid++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", numValid)
}

func isValid(tkns map[string]string) bool {
	checklist := reqFields{}
	for tkn, val := range tkns {
		switch tkn {
		case "byr":
			{
				birthYear, err := strconv.Atoi(val)
				if err != nil {
					break
				}
				checklist.byr = isValBetween(birthYear, 1920, 2002)
			}
		case "iyr":
			{
				issueYear, err := strconv.Atoi(val)
				if err != nil {
					break
				}
				checklist.iyr = isValBetween(issueYear, 2010, 2020)
			}
		case "eyr":
			{
				expYear, err := strconv.Atoi(val)
				if err != nil {
					break
				}
				checklist.eyr = isValBetween(expYear, 2020, 2030)
			}
		case "hgt":
			checklist.hgt = isHeightValid(val)
		case "hcl":
			checklist.hcl = isHairClrValid(val)
		case "ecl":
			checklist.ecl = isEyeClrValid(val)
		case "pid":
			checklist.pid = isPidValid(val)
		case "cid":
			continue
		default:
			break
		}
	}
	return checklist.byr && checklist.iyr && checklist.eyr && checklist.hgt && checklist.hcl && checklist.ecl && checklist.pid
}

func isPidValid(pid string) bool {
	return matchRegex(`\b\d{9}\b`, pid)
}

func isEyeClrValid(inClr string) bool {
	validClrs := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, clr := range validClrs {
		if inClr == clr {
			return true
		}
	}
	return false
}

func isHairClrValid(clr string) bool {
	return matchRegex(`#[0-9a-f]{6}`, clr)
}

func matchRegex(regex string, challenge string) bool {
	res, err := regexp.MatchString(regex, challenge)
	if err != nil {
		log.Fatalf("Invlid regex pattern: %s", err)
	}
	return res
}

func isHeightValid(heigth string) bool {
	r := regexp.MustCompile(`(\d{2,3})(\w{2})`)
	submatches := r.FindStringSubmatch(heigth)
	if len(submatches) != 3 {
		return false
	}
	unit := r.FindStringSubmatch(heigth)[2]
	heightVal, err := strconv.Atoi(r.FindStringSubmatch(heigth)[1])
	if err != nil {
		return false
	}
	switch unit {
	case "cm":
		return isValBetween(heightVal, 150, 193)
	case "in":
		return isValBetween(heightVal, 59, 76)
	}
	return false
}

func isValBetween(val int, lo int, hi int) bool {
	return (val >= lo) && (val <= hi)
}

func mapFromList(in []string) map[string]string {
	r := regexp.MustCompile(`(\w{3}):(.*)`)
	res := make(map[string]string)
	for _, info := range in {
		res[r.FindStringSubmatch(info)[1]] = r.FindStringSubmatch(info)[2]
	}
	return res
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
