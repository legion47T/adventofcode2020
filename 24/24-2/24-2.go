package main

import (
	"bufio"
	"log"
	"os"

	co "github.com/daniel-dsouza/hexago/coordinate"
	"github.com/daniel-dsouza/hexago/storage"
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

	floor := make(storage.SimpleMap)

	var lineNo int
	for sc.Scan() {
		lineNo++
		// log.Println("Line", lineNo)
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
		newTile := co.NewCube(currentTilePos.x, currentTilePos.y, currentTilePos.z)
		if tile, ok := floor.Get(newTile); ok {
			floor.Set(newTile, !tile.(bool))
		} else {
			floor.Set(newTile, true)
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	// log.Println(floor)

	// var res int

	for day := 1; day <= 100; day++ {
		// log.Println("Floor", floor)
		futureFloor := make(storage.SimpleMap)
		for pos, tile := range floor {
			if tile.(bool) {
				// log.Println("Pos", pos)
				var numBlack int
				immediateNeighbors := pos.GetNeighbors()
				d2Neighbors := make(storage.SimpleMap)
				for _, neighborPos := range immediateNeighbors {

					if neighbor, ok := floor.Get(neighborPos); ok && neighbor.(bool) {
						numBlack++
					}

					for _, neighbor := range neighborPos.GetNeighbors() {
						if pos.Distance(neighbor) <= 2 {
							if _, ok := d2Neighbors.Get(neighbor); !ok {
								if neighborVal, ok := floor.Get(neighbor); ok {
									d2Neighbors.Set(neighbor, neighborVal)
								} else {
									d2Neighbors.Set(neighbor, false)
								}
							}
						}
					}
				}

				if numBlack == 1 || numBlack == 2 {
					futureFloor.Set(pos, true)
					// log.Println(pos, "will be true")
				} else {
					futureFloor.Set(pos, false)
					// log.Println(pos, "will be false")
				}

				// log.Println("d2Neighbors", d2Neighbors, "len", len(d2Neighbors))
				for d2tile, d2TileVal := range d2Neighbors {
					// if we dont have it on the futureFloor already and it's current value is false
					// if _, ok := futureFloor.Get(d2tile); !ok && !d2TileVal.(bool) {
					if !d2TileVal.(bool) {
						var numBlack int
						// take a look at all neighbors in distance 2 of current main tile
						for _, neighbor := range floor.Neighbors(d2tile) {
							if neighborVal, ok := floor.Get(neighbor); ok && neighborVal.(bool) {
								numBlack++
							}
						}
						if numBlack == 2 {
							futureFloor.Set(d2tile, true)
						}
					}
				}
				// log.Println("FutureFloor", futureFloor)
			}
		}

		floor = futureFloor
		// log.Println("Floor", floor)

		if day < 10 || day%10 == 0 {
			var numBlack int
			for _, tile := range floor {
				if tile.(bool) {
					numBlack++
				}
			}
			log.Println("Day", day, ":", numBlack)
		}
	}

	// log.Println("Result:", res)
}
