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

type block struct {
	mask   string
	instrs []instr
}

type instr struct {
	addr int
	val  int64
}

var instrRegex *regexp.Regexp

func main() {
	// file, err := os.Open("../test14.txt")
	file, err := os.Open("../input14.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	instrRegex = regexp.MustCompile(`mem\[(\d*)\] = (\d*)`)

	blocks := make([]block, 0)
	currBlock := block{}
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "mask") {
			if len(currBlock.mask) > 0 {
				blocks = append(blocks, currBlock)
			}
			currBlock = block{}
			currBlock.mask = parseMask(line)
			continue
		}
		currBlock.instrs = append(currBlock.instrs, parseInstr(line))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	blocks = append(blocks, currBlock)
	fmt.Println(blocks)

	mem := make(map[int]int64, 0)
	for _, block := range blocks {
		for _, ins := range block.instrs {
			mem[ins.addr] = applyMask(ins.val, block.mask)
		}
	}
	fmt.Println(mem)

	var res int64
	for _, val := range mem {
		res += val
	}

	log.Println(res)
}

func applyMask(val int64, mask string) int64 {
	rMask := reverse(mask)
	for pos, bitRune := range rMask {
		switch bitRune {
		case 'X':
			continue
		case '0':
			val = clearBit(val, pos)
		case '1':
			val = setBit(val, pos)
		}
	}
	return val
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

// Sets the bit at pos in the integer n.
func setBit(n int64, pos int) int64 {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int64, pos int) int64 {
	mask := int64(^(1 << pos))
	n &= mask
	return n
}

func parseMask(in string) string {
	return strings.Split(in, " ")[2]
}

func parseInstr(in string) instr {
	submatches := instrRegex.FindStringSubmatch(in)
	addr, _ := strconv.Atoi(submatches[1])
	val, err := strconv.ParseInt(submatches[2], 10, 64)
	if err != nil {
		log.Fatalf("Could not parse mem value: %v", err)
	}
	return instr{addr: addr, val: val}
}
