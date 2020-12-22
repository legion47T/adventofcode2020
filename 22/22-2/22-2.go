package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("../test22.txt")
	// file, err := os.Open("test22-2.txt")
	file, err := os.Open("../input22.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	players := make([][]int, 0)
	playerParsing := 0
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			playerParsing++
			continue
		}
		if strings.HasPrefix(line, "Player") {
			players = append(players, make([]int, 0))
			continue
		}
		val, _ := strconv.Atoi(line)
		players[playerParsing] = append(players[playerParsing], val)
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	playGame(&players[0], &players[1])

	var res int
	mult := 1
	for _, player := range players {
		for i := len(player) - 1; i >= 0; i-- {
			res += mult * player[i]
			mult++
		}
	}

	log.Println(res)
}

func playGame(deck1, deck2 *[]int) int {
	playedConfigurations := make([][][]int, 0)
	for len(*deck1) != 0 && len(*deck2) != 0 {

		for _, config := range playedConfigurations {
			if testEq(config[0], *deck1) && testEq(config[1], *deck2) {
				return 0
			}
		}

		deck1copy := make([]int, len(*deck1))
		deck2copy := make([]int, len(*deck2))
		copy(deck1copy, *deck1)
		copy(deck2copy, *deck2)
		playedConfigurations = append(playedConfigurations, [][]int{deck1copy, deck2copy})
		p1top := (*deck1)[0]
		p2top := (*deck2)[0]
		*deck1 = remove(*deck1, 0)
		*deck2 = remove(*deck2, 0)
		subGameWinner := -1
		if len(*deck1) >= p1top && len(*deck2) >= p2top {
			subDeck1 := make([]int, p1top)
			subDeck2 := make([]int, p2top)
			copy(subDeck1, (*deck1)[:p1top])
			copy(subDeck2, (*deck2)[:p2top])
			subGameWinner = playGame(&subDeck1, &subDeck2)
		}
		switch subGameWinner {
		case 0:
			*deck1 = append(*deck1, p1top, p2top)
			continue
		case 1:
			*deck2 = append(*deck2, p2top, p1top)
			continue
		}
		if p1top > p2top {
			*deck1 = append(*deck1, p1top, p2top)
		} else {
			*deck2 = append(*deck2, p2top, p1top)
		}
	}
	if len(*deck1) > 0 {
		return 0
	} else {
		return 1
	}
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func reverse(value string) string {
	// Convert string to rune slice.
	// ... This method works on the level of runes, not bytes.
	data := []rune(value)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}

func unique(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

// func remove(s []int, i int) []int {
// 	s[i] = s[len(s)-1]
// 	// We do not need to put s[i] at the end, as it will be discarded anyway
// 	return s[:len(s)-1]
// }

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
