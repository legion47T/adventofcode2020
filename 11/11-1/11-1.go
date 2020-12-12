package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type seat struct {
	x, y              int
	present, occupied bool
}

type coord struct {
	x, y int
}

func main() {
	// file, err := os.Open("../test11.txt")
	file, err := os.Open("../input11.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var room, futureRoom *[][]seat
	tempRoom := make([][]seat, 0)
	room = &tempRoom
	tempRoom2 := make([][]seat, 0)
	futureRoom = &tempRoom2
	var y int
	for sc.Scan() {
		row := parseSeats(sc.Text(), y)
		*room = append(*room, row)
		row = parseSeats(sc.Text(), y)
		*futureRoom = append(*futureRoom, row)
		y++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		// fmt.Println("room:", room)
		// fmt.Println("futureRoom:", futureRoom)
		// fmt.Printf("roomaddr %p futRoomAddr %p\n", room, futureRoom)
		room, futureRoom = futureRoom, room
		// tempRoom := room
		// room = futureRoom
		// futureRoom = tempRoom
		// fmt.Printf("roomaddr %p futRoomAddr %p\n", room, futureRoom)
		for y, row := range *room {
			for x, s := range row {
				if s.present {
					// fmt.Print((*futureRoom)[y][x])
					(*futureRoom)[y][x].occupied = willBeOccupied(room, s)
					// fmt.Println(" ", (*futureRoom)[y][x])

					// fmt.Println("willBe", willBeOccupied(&room, s), "new status", futureRoom[y][x].occupied)
				}
			}
		}
		if !stateHasChanged(room, futureRoom) {
			break
		}
	}

	// fmt.Println(room)
	var res int
	res = numSeatsOccupied(room)

	fmt.Println(res)
}

func numSeatsOccupied(room *[][]seat) int {
	var sum int
	for _, row := range *room {
		for _, s := range row {
			if s.occupied {
				sum++
			}
		}
	}
	return sum
}

func stateHasChanged(currentRoom, futureRoom *[][]seat) bool {
	for y, row := range *currentRoom {
		for x, s := range row {
			if s.occupied != (*futureRoom)[y][x].occupied {
				return true
			}
		}
	}
	return false
}

func willBeOccupied(room *[][]seat, s seat) bool {
	if !s.present {
		return false
	}
	occupiedSeats := make([]coord, 0)
	for y := s.y - 1; y <= s.y+1; y++ {
		if y < 0 {
			continue
		} else if y >= len(*room) {
			break
		}
		row := (*room)[y]
		for x := s.x - 1; x <= s.x+1; x++ {
			if x == s.x && y == s.y {
				continue
			}
			if x < 0 {
				continue
			} else if x >= len(row) {
				break
			}
			currSeat := row[x]
			if currSeat.present && currSeat.occupied {
				occupiedSeats = append(occupiedSeats, coord{x: x, y: y})
			}
		}
	}

	numOccupied := len(occupiedSeats)
	if !s.occupied && numOccupied == 0 {
		return true
	}

	if s.occupied && numOccupied >= 4 {
		return false
	}

	return s.occupied
}

func parseSeats(input string, y int) []seat {
	seats := make([]seat, 0)
	for x, symb := range input {
		switch symb {
		case '.':
			seats = append(seats, seat{x: x, y: y})
		case 'L':
			seats = append(seats, seat{x: x, y: y, present: true})
		}
	}
	return seats
}
