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

type codeT struct {
	full   string
	row    string
	column string
}

type seat struct {
	code   codeT
	row    int
	column int
	id     int
}

func main() {
	// file, err := os.Open("test04-2_valid.txt")
	// file, err := os.Open("test04-2_invalid.txt")
	// file, err := os.Open("test04-2.txt")
	file, err := os.Open("../input05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := findSeat(scanner.Text())
		if seat.id > res {
			res = seat.id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", res)
}

func findSeat(code string) seat {
	c := splitCode(code)
	row := findRowOrColumn(c.row)
	column := findRowOrColumn(c.column)
	id := calculateID(row, column)

	return seat{code: c, row: row, column: column, id: id}
}

func findRowOrColumn(code string) int {
	code = strings.ReplaceAll(code, "F", "0")
	code = strings.ReplaceAll(code, "B", "1")
	code = strings.ReplaceAll(code, "L", "0")
	code = strings.ReplaceAll(code, "R", "1")
	res, err := strconv.ParseInt(code, 2, 64)
	if err != nil {
		log.Fatalf("parse error findRowOrColumn: %s", err)
	}
	return int(res)
}

func splitCode(code string) codeT {
	r := regexp.MustCompile(`(\w{7})(\w{3})`)
	submatches := r.FindStringSubmatch(code)
	if len(submatches) != 3 {
		log.Fatal("Code", code, "does not conform to schema")
	}
	return codeT{code, submatches[1], submatches[2]}
}

func calculateID(row int, column int) int {
	return row*8 + column
}
