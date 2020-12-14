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
	// file, err := os.Open("test14-2.txt")
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
			prepMask := prepareMask(ins.addr, block.mask)
			adresses := calculateAddresses(prepMask)
			for _, address := range adresses {
				addrs, _ := strconv.ParseInt(address, 2, 64)
				// addrs := applyMask(int64(ins.addr), mask)
				mem[int(addrs)] = ins.val
			}
		}
	}
	fmt.Println(mem)

	var res int64
	for _, val := range mem {
		res += val
	}

	log.Println(res)
}

func prepareMask(addr int, mask string) string {
	addrString := strconv.FormatInt(int64(addr), 2)
	zeroes := make([]rune, 0)
	for i := 1; i <= len(mask)-len(addrString); i++ {
		zeroes = append(zeroes, '0')
	}
	addrString = strings.Join([]string{string(zeroes), addrString}, "")

	maskedAddr := make([]rune, 0)
	for i, bitRune := range mask {
		if bitRune == '0' {
			maskedAddr = append(maskedAddr, []rune(addrString)[i])
		} else {
			maskedAddr = append(maskedAddr, bitRune)
		}
	}
	return string(maskedAddr)
}

func calculateAddresses(mask string) []string {
	masks := make([]string, 0)
	if strings.Contains(mask, "X") {
		for i, bit := range mask {
			if bit == 'X' {
				tempMask := []rune(mask)
				tempMask[i] = '1'
				for _, calcMask := range calculateAddresses(string(tempMask)) {
					masks = append(masks, calcMask)
				}
				tempMask[i] = '0'
				for _, calcMask := range calculateAddresses(string(tempMask)) {
					masks = append(masks, calcMask)
				}
				break
			}
		}
	} else {
		masks = append(masks, mask)
	}
	return unique(masks)
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// func applyMask(val int64, mask string) int64 {
// 	rMask := strings.TrimRight(reverse(mask), "0")
// 	for pos, bitRune := range rMask {
// 		if bitRune == '1' {
// 			val = setBit(val, pos)
// 		}
// 	}
// 	return val
// }

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
