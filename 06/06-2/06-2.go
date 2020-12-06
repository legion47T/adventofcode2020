package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// file, err := os.Open("../test06.txt")
	file, err := os.Open("../input06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var res int

	groupScanner := bufio.NewScanner(file)
	groupScanner.Split(doubleNLSplitFunc)
	for groupScanner.Scan() {
		group := groupScanner.Text()
		// fmt.Println("Group:\n", group, "\n")
		personScanner := bufio.NewScanner(strings.NewReader(group))

		groupChoice := make(map[rune]int)
		var groupSize int

		for personScanner.Scan() {
			person := personScanner.Text()
			groupSize++

			for _, choice := range person {
				groupChoice[choice]++
			}
		}
		// fmt.Println("groupChoice:\n", groupChoice)

		if err := personScanner.Err(); err != nil {
			log.Fatal(err)
		}

		// fmt.Println("groupSize:", groupSize, "\n\n")
		for _, num := range groupChoice {
			if num == groupSize {
				res++
			}
		}
	}

	if err := groupScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", res)
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
