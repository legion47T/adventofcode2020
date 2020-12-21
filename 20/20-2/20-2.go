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

type direction int

const (
	north direction = iota
	east
	south
	west
	iNorth
	iEast
	iSouth
	iWest
)

type tile struct {
	id             int
	image          []string
	borders        map[direction]int64
	neighbors      []int
	placed         bool
	foundNeighbors int
}

type coord struct {
	x, y int
}

var tileIDRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test20.txt")
	file, err := os.Open("../input20.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	tileIDRegex = regexp.MustCompile(`Tile (\d*):`)

	tiles := make(map[int]tile)
	borderMap := make(map[int64][]int)
	currentTile := tile{}
	lastLineEmpty := false
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			if lastLineEmpty {
				break
			}
			lastLineEmpty = true
			currentTile = calculateBorders(currentTile)
			tiles[currentTile.id] = currentTile
			insertIntoBorderMap(currentTile, &borderMap)
			currentTile = tile{}
			continue
		}
		if strings.HasPrefix(line, "Tile") {
			submatches := tileIDRegex.FindStringSubmatch(line)
			id, _ := strconv.Atoi(submatches[1])
			currentTile.id = id
			continue
		}
		lastLineEmpty = false
		currentTile.image = append(currentTile.image, parseTileLine(line))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	// currentTile = calculateBorders(currentTile)
	// tiles[currentTile.id] = currentTile
	// insertIntoBorderMap(currentTile, &borderMap)

	log.Println(tiles)
	log.Println(borderMap)
	for _, tileIDs := range borderMap {
		if len(tileIDs) == 2 {
			tile1 := tiles[tileIDs[0]]
			var t1neighbors []int
			if len(tile1.neighbors) == 0 {
				t1neighbors = make([]int, 0)
			} else {
				t1neighbors = tile1.neighbors
			}
			t1neighbors = append(t1neighbors, tileIDs[1])
			tiles[tileIDs[0]] = tile{id: tile1.id, image: tile1.image, borders: tile1.borders, neighbors: unique(t1neighbors)}

			tile2 := tiles[tileIDs[1]]
			var t2neighbors []int
			if len(tile2.neighbors) == 0 {
				t2neighbors = make([]int, 0)
			} else {
				t2neighbors = tile2.neighbors
			}
			t2neighbors = append(t2neighbors, tileIDs[0])
			tiles[tileIDs[1]] = tile{id: tile2.id, image: tile2.image, borders: tile2.borders, neighbors: unique(t2neighbors)}
		}
	}
	log.Println(tiles)

	prePicture := make(map[coord]tile, len(tiles))
	for _, currentTile := range tiles {
		if len(currentTile.neighbors) == 2 {
			currentTile = tile{id: currentTile.id, image: currentTile.image, borders: currentTile.borders, neighbors: currentTile.neighbors, placed: true}
			tiles[currentTile.id] = currentTile
			prePicture[coord{0, 0}] = currentTile
			break
		}
	}
	// var allPlaced bool

	// for !allPlaced {
	// 	for _, currentTile := range tiles {
	// 		if !currentTile.placed {
	// 			allPlaced = false
	// 			for _, neighbor := range currentTile.neighbors {
	// 				for position, placedTile := range prePicture {
	// 					if placedTile.id == neighbor {
	// 						for myBorderDir, myBorder := range currentTile.borders {
	// 							for placedBorderDir, placedBorder := range placedTile.borders {
	// 								if myBorder == placedBorder {
	// 									switch placedBorderDir {
	// 									case east:
	// 										switch myBorderDir {
	// 										case north:
	// 											// rotate right 90°
	// 										case east:
	// 											// rotate right 180°
	// 										case south:
	// 											// rotate right 270°
	// 										case west:
	// 											// perfect fit
	// 										case iNorth:
	// 											// rotate right 90° then flip vertical
	// 										case iEast:
	// 											// east + flip
	// 										case iSouth:
	// 											// south + flip
	// 										case iWest:
	// 											// flip horizontal
	// 										}
	// 									case south:
	// 										switch myBorderDir {

	// 										}
	// 									}
	// 									break
	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	var res int = 1
	fmt.Println(res)
}

func insertIntoBorderMap(currentTile tile, borderMap *map[int64][]int) {
	for _, borderID := range currentTile.borders {
		var adjacentTileList []int
		if tileList, ok := (*borderMap)[borderID]; ok {
			adjacentTileList = tileList
		} else {
			adjacentTileList = make([]int, 0)
		}
		adjacentTileList = append(adjacentTileList, currentTile.id)
		(*borderMap)[borderID] = adjacentTileList
	}
}

func calculateBorders(currentTile tile) tile {
	newTile := tile{id: currentTile.id, image: currentTile.image}

	yDim := len(currentTile.image)

	borders := make(map[direction]int64, 8)
	northRow := currentTile.image[0]
	borders[north], _ = strconv.ParseInt(northRow, 2, 64)
	invNorthRow := reverse(northRow)
	borders[iNorth], _ = strconv.ParseInt(invNorthRow, 2, 64)

	southRow := currentTile.image[yDim-1]
	borders[south], _ = strconv.ParseInt(southRow, 2, 64)
	invSouthRow := reverse(southRow)
	borders[iSouth], _ = strconv.ParseInt(invSouthRow, 2, 64)

	xDim := len(northRow)

	var eastSB, westSB strings.Builder

	for y := 0; y < yDim; y++ {
		eastSB.WriteByte(currentTile.image[y][xDim-1])
		westSB.WriteByte(currentTile.image[y][0])
	}

	eastRow := eastSB.String()
	borders[east], _ = strconv.ParseInt(eastRow, 2, 64)
	invEastRow := reverse(eastRow)
	borders[iEast], _ = strconv.ParseInt(invEastRow, 2, 64)

	westRow := westSB.String()
	borders[west], _ = strconv.ParseInt(westRow, 2, 64)
	invWestRow := reverse(westRow)
	borders[iWest], _ = strconv.ParseInt(invWestRow, 2, 64)

	newTile.borders = borders
	return newTile
}

func parseTileLine(in string) string {
	in = strings.ReplaceAll(in, ".", "0")
	return strings.ReplaceAll(in, "#", "1")
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
