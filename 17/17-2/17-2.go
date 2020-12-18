package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cube struct {
	c            coord
	isActive     bool
	willBeActive bool
}

type coord struct {
	x, y, z, w int
}

func main() {
	// file, err := os.Open("../test17.txt")
	file, err := os.Open("../input17.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	iterations := 6

	pocketDimension := make(map[coord]cube, 0)
	var maxCoords, minCoords coord
	var y int
	for sc.Scan() {
		for x, r := range sc.Text() {
			var current cube
			location := coord{x, y, 0, 0}
			switch r {
			case '.':
				current = cube{location, false, false}
			case '#':
				current = cube{location, true, false}
			}
			pocketDimension[location] = current
			if x > maxCoords.x {
				maxCoords.x = x
			}
		}
		if y > maxCoords.y {
			maxCoords.y = y
		}
		y++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	printPocketDim(maxCoords, minCoords, &pocketDimension)
	for i := 1; i <= iterations; i++ {
		maxCoords = coord{maxCoords.x + 1, maxCoords.y + 1, maxCoords.z + 1, maxCoords.w + 1}
		minCoords = coord{minCoords.x - 1, minCoords.y - 1, minCoords.z - 1, minCoords.w - 1}
		prepareNeighbors(maxCoords, minCoords, &pocketDimension)
		// log.Println(i, pocketDimension[coord{0, 0, 0}])
		updateActiveState(&pocketDimension)
		// log.Println(i, pocketDimension[coord{0, 0, 0}])
		setNextActiveState(&pocketDimension)
		// log.Println(i, pocketDimension[coord{0, 0, 0}])
		// fmt.Println("Iteration", i, "max", maxCoords, "min", minCoords)
		// printPocketDim(maxCoords, minCoords, &pocketDimension)
		// maxCoords = coord{maxCoords.x + 1, maxCoords.y + 1, maxCoords.z + 1}
		// minCoords = coord{minCoords.x - 1, minCoords.y - 1, minCoords.z - 1}
	}
	// printPocketDim(maxCoords, minCoords, &pocketDimension)

	var res int
	res = numCubesActive(&pocketDimension)

	fmt.Println(res)
}

func setNextActiveState(dim *map[coord]cube) {
	for _, c := range *dim {
		(*dim)[c.c] = cube{c.c, c.willBeActive, false}
	}
}

func updateActiveState(dim *map[coord]cube) {
	for _, c := range *(dim) {
		c.setWillBeActive(dim)
	}
}

func prepareNeighbors(maxCoords, minCoords coord, dim *map[coord]cube) {
	for w := minCoords.w; w <= maxCoords.w; w++ {
		for z := minCoords.z; z <= maxCoords.z; z++ {
			for y := minCoords.y; y <= maxCoords.y; y++ {
				for x := minCoords.x; x <= maxCoords.x; x++ {
					currentLocation := coord{x, y, z, w}
					if _, ok := (*dim)[currentLocation]; !ok {
						(*dim)[currentLocation] = cube{currentLocation, false, false}
					}
				}
			}
		}
	}
}

func (c *cube) countActiveNeighbors(dim *map[coord]cube) int {
	var activeCount int
	for w := c.c.w - 1; w <= c.c.w+1; w++ {
		for z := c.c.z - 1; z <= c.c.z+1; z++ {
			for y := c.c.y - 1; y <= c.c.y+1; y++ {
				for x := c.c.x - 1; x <= c.c.x+1; x++ {
					if c.c.x == x && c.c.y == y && c.c.z == z && c.c.w == w {
						continue
					}
					neighbor := (*dim)[coord{x, y, z, w}]
					if neighbor.isActive {
						activeCount++
					}
				}
			}
		}
	}
	return activeCount
}

func (c *cube) setWillBeActive(dim *map[coord]cube) {
	activeNeighbors := c.countActiveNeighbors(dim)
	if c.isActive {
		if activeNeighbors == 2 || activeNeighbors == 3 {
			(*dim)[c.c] = cube{c.c, c.isActive, true}
			return
		}
	} else {
		if activeNeighbors == 3 {
			(*dim)[c.c] = cube{c.c, c.isActive, true}
			return
		}
	}
	(*dim)[c.c] = cube{c.c, c.isActive, false}
}

func printPocketDim(maxCoords, minCoords coord, dim *map[coord]cube) {
	for w := minCoords.w; w <= maxCoords.w; w++ {
		for z := minCoords.z; z <= maxCoords.z; z++ {
			fmt.Println("z=", z, ",w=", w)
			for y := minCoords.y; y <= maxCoords.y; y++ {
				// fmt.Printf("%d ", y)
				for x := minCoords.x; x <= maxCoords.x; x++ {
					if (*dim)[coord{x, y, z, w}].isActive {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print("\n")
			}
			fmt.Print("\n")
		}
	}
}

func numCubesActive(dim *map[coord]cube) int {
	var sum int
	for _, c := range *dim {
		if c.isActive {
			sum++
		}
	}
	return sum
}
