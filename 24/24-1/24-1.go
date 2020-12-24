package main

import (
	"bufio"
	"log"
	"os"
)

type coord struct {
	x, y, z int
}

func main() {
	// file, err := os.Open("../test24.txt")
	file, err := os.Open("../input24.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	floor := make(map[coord]bool)

	var lineNo int
	for sc.Scan() {
		lineNo++
		log.Println("Line", lineNo)
		line := sc.Text()
		currentTilePos := coord{}
		var north, south bool
		for _, sym := range line {
			switch sym {
			case 'n':
				north = true
			case 's':
				south = true
			case 'e':
				if north {
					currentTilePos.x++
					currentTilePos.z--
				} else if south {
					currentTilePos.y--
					currentTilePos.z++
				} else {
					currentTilePos.x++
					currentTilePos.y--
				}
				north, south = false, false
				// log.Println(currentTilePos)
			case 'w':
				if north {
					currentTilePos.y++
					currentTilePos.z--
				} else if south {
					currentTilePos.x--
					currentTilePos.z++
				} else {
					currentTilePos.x--
					currentTilePos.y++
				}
				north, south = false, false
				// log.Println(currentTilePos)
			}
		}
		if tile, ok := floor[currentTilePos]; ok {
			floor[currentTilePos] = !tile
		} else {
			floor[currentTilePos] = true
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println(floor)

	var res int

	for _, tile := range floor {
		if tile {
			res++
		}
	}

	log.Println("Result:", res)
}
