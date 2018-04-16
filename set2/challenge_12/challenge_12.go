package challenge_12

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"crypto/rand"

	"friendly-doodle/set2/challenge_11"
	"friendly-doodle/set2/challenge_9"
	"friendly-doodle/utils"
)

// generate a random 16 byte key
var key = make([]byte, 16)
var _, _ = rand.Read(key)

// Encrypt takes a message and key, pads it appropriately, encrypts the message, and returns the cipher text
func Encrypt(msg []byte, key []byte) []byte {

	paddedMsg := challenge_9.PadMessage(msg, len(key))
	cipherText := make([]byte, len(paddedMsg))

	cipherText = challenge_11.EcbEncrypt(paddedMsg, key)
	return cipherText
}

// Use this function as a way to contain the secret and the password from the rest of the program
func encryptionWrapper(myMessage []byte) []byte {
	secret_base64 := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	secretMessage, err := base64.StdEncoding.DecodeString(secret_base64)
	utils.Check(err)

	//	key := []byte("12345678abcdefgh")
	msg := append(myMessage, secretMessage...)
	return Encrypt(msg, key)
}

// A version of the encryption wrapper w/o the real secret
func encryptionWrapperDev(myMessage []byte) []byte {
	secretMessage := []byte("Hi secret message")
	key := []byte("12345678abcdefgh")
	msg := append(myMessage, secretMessage...)
	return Encrypt(msg, key)
}

func buildAttackMap(blockLen int, blockNum int, known []byte) map[string]byte {

	// generate a map of all possible strings to thier respective byte
	plainMap := buildMap(blockLen, known)
	attackMap := make(map[string]byte)

	// encrypt each string to create an attack map
	for mapKey, value := range plainMap {
		attackBlock := utils.ECBBlock{}
		attackBlock.Text = encryptionWrapper([]byte(mapKey))
		attackBlock.BlockLen = blockLen

		// block 0 is a special case, we are not prepending any known blocks, so it will be block 0.
		if blockNum == 0 {
			_, err := attackBlock.SelectBlock(0)
			utils.Check(err)
		} else {
			_, err := attackBlock.SelectBlock(blockNum - 1)
			utils.Check(err)
		}
		// build out the map
		attackMap[string(attackBlock.CurrentBlock)] = value
	}

	return attackMap

}

// buildMap takes the block length, and the known bytes to create a map of
// all possible outputs from the oracle
func buildMap(blockLen int, known []byte) map[string]byte {
	// create a base string of one less than the block length
	// append each byte to the base string and add the new string and byte to a string to byte map
	plainMap := make(map[string]byte)
	var baseString string

	if blockLen <= len(known) {
		baseString = string(known[len(known)%blockLen+1:])
	} else {
		neededPrepends := blockLen - len(known) - 1

		var buffer bytes.Buffer
		for count := 0; count < neededPrepends; count++ {
			buffer.WriteString("\x00")
		}
		prependString := buffer.String()
		baseString = prependString + string(known)
	}
	var byt byte
	for byt = 0; byt <= 127; byt++ {
		plainMap[baseString+string(byt)] = byt
	}

	return plainMap
}

// create AttackStrings takes the secret mesage and appends bytes to the begining
func createAttackStrings(blockLen int) []utils.ECBBlock {
	attackPrefixes := make([][]byte, blockLen)
	for i := 0; i < blockLen; i++ {
		attackPrefixes[i] = make([]byte, i)
	}

	attackStrings := make([]utils.ECBBlock, len(attackPrefixes))

	for pos, str := range attackPrefixes {
		attackStrings[pos].Text = encryptionWrapper(str)
		attackStrings[pos].BlockLen = blockLen
	}

	return attackStrings
}

func attack() {
	// create the set of attack strings
	// iteratively decrypt the secret message byte by byte

	// this works because I am able to append anything to the secret string,
	// and have it sent through the same encryption process. In ECB ecnryption this will
	// cause identical inputs to create identical outputs.

	blockSize := 16
	blockNum := 0
	attackStringSelector := 15

	known := make([]byte, 0)

	attackStrings := createAttackStrings(blockSize)

	fmt.Println("BlockNum: 0")
	for range attackStrings[0].Text {
		// decrypt each byte
		attackMap := buildAttackMap(blockSize, blockNum, known)
		attackedBlock, err := attackStrings[attackStringSelector].SelectBlock(blockNum)
		utils.Check(err)

		byt, ok := attackMap[string(attackedBlock)]
		if ok {
			known = append(known, byt)
		}

		if attackStringSelector == 0 {
			attackStringSelector = 15
			blockNum++
		} else {
			attackStringSelector--
		}
		fmt.Printf("%s \n", known)
	}
}

func Main() {
	attack()
}
