package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// file, err := os.Open("../test25.txt")
	file, err := os.Open("../input25.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var publicKey1, publicKey2 *big.Int
	for sc.Scan() {
		line := sc.Text()
		val, _ := strconv.Atoi(string(line))
		// if publicKey1.Cmp(big.NewInt(0)) == 0 {
		if publicKey1 == nil {
			publicKey1 = big.NewInt(int64(val))
			continue
		}
		publicKey2 = big.NewInt(int64(val))
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	round := 0
	lastNumber := big.NewInt(1)
	subjectNumber := big.NewInt(7)
	modulus := big.NewInt(20201227)
	var exp1, exp2 *big.Int
	for {
		round++
		lastNumber = lastNumber.Mod(lastNumber.Mul(lastNumber, subjectNumber), modulus)
		if publicKey1.Cmp(lastNumber) == 0 {
			exp1 = big.NewInt(int64(round))
			break
		}
		if publicKey2.Cmp(lastNumber) == 0 {
			exp2 = big.NewInt(int64(round))
			break
		}
	}

	var encKey big.Int
	if exp1 != nil {
		encKey.Exp(publicKey2, exp1, modulus)
	}
	if exp2 != nil {
		encKey.Exp(publicKey1, exp2, modulus)
	}

	// if encKey < 0 {
	// 	encKey += modulus
	// }
	// var res int

	log.Println(encKey)
}
