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

	round := 1
	for len(players[0]) != 0 && len(players[1]) != 0 {
		log.Println("Round", round)
		log.Println("Player 1:", players[0])
		log.Println("Player 2:", players[1])
		round++
		p1top := players[0][0]
		p2top := players[1][0]
		players[0] = remove(players[0], 0)
		players[1] = remove(players[1], 0)
		if p1top > p2top {
			players[0] = append(players[0], p1top, p2top)
		} else {
			players[1] = append(players[1], p2top, p1top)
		}
	}

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
