package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"math/bits"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func hamming(str1 []byte, str2 []byte) (int, error) {
	// use hamming distance code from challenge 6 to find common blocks

	if len(str1) != len(str2) {
		err := errors.New("string length mismatch")
		return 0, err
	}
	count := 0
	for i := range str1 {
		xord := str1[i] ^ str2[i]
		count += bits.OnesCount8(xord)
	}

	return count, nil
}
func detect_ecb(ciphertext []byte, keyLen int) int {
	//		// assume that the overlap will occur in the first two blocks
	//		blockOne := ciphertext[:keyLen]
	//		blockTwo := ciphertext[keyLen : keyLen*2]
	//
	//		simm, err := hamming(blockOne, blockTwo)
	//		check(err)
	//		fmt.Println(simm)
	//
	//		return true
	// bad assumption

	numberOfBlocks := len(ciphertext) / keyLen
	blockOne := make([]byte, keyLen)
	blockTwo := make([]byte, keyLen)
	var overlap int
	var total int
	var err error

	for i := 0; i < numberOfBlocks-1; i++ {
		for j := 0; j <= numberOfBlocks-1; j++ {
			// really don't need to do every comparison, lots of redundant work
			// lets only do half of that
			if i > j {
				blockOne = ciphertext[(i * 16):((i + 1) * 16)]
				blockTwo = ciphertext[(j * 16):((j + 1) * 16)]
				overlap, err = hamming(blockOne, blockTwo)
				check(err)
				total += overlap
			}
		}
	}

	// total = total / 40
	//fmt.Println(total)

	//	// looking at the normalized values 45 is by far the lowest
	//	// print that out block by block
	//	if total == 45 {
	//		for i := 0; i < 10; i++ {
	//			fmt.Println(ciphertext[(i * 16):((i + 1) * 16)])
	//		}
	//	}
	// ok yea, thats the one now lets just find it  without the hard coded value
	total = total / 40
	return total
}

func main() {
	// Read in the file line by line
	file, err := os.Open("challenge_8.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	// track which ciphertext has the smallest difference between block
	best := 


	for scanner.Scan() {
		hex_decoded, err := hex.DecodeString(scanner.Text())
		check(err)

		// each ciphertext is 10 blocks long
		// fmt.Println(len(hex_decoded))

		detect_ecb(hex_decoded, 16)
	}
}
