package challenge_12

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"friendly-doodle/utils"

	"friendly-doodle/set2/challenge_11"
	"friendly-doodle/set2/challenge_9"
	"math/rand"
)

var key = make([]byte, 16)
var _, _ = rand.Read(key)

// Encrypt takes a message and key, pads it appropriately, encrypts the message, and returns the cipher text
func Encrypt(msg []byte, key []byte) []byte {

	paddedMsg := challenge_9.PadMessage(msg, len(key))
	cipherText := make([]byte, len(paddedMsg))

	cipherText = challenge_11.EcbEncrypt(paddedMsg, key)
	return cipherText
}

func encryptionWrapper(myMessage []byte) []byte {
	secret_base64 := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	secretMessage, err := base64.StdEncoding.DecodeString(secret_base64)
	utils.Check(err)
	key := []byte("12345678abcdefgh")
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

	plainMap := buildMap(blockLen, known)
	// encrypt each string to create an attack map
	attackMap := make(map[string]byte)

	for mapKey, value := range plainMap {
		attackMap[string(encryptionWrapper([]byte(mapKey)))[((blockNum/2)*blockLen):(((blockNum/2)+1)*blockLen)]] = value
		fmt.Printf("%v \n", encryptionWrapper([]byte(mapKey))[(blockNum*blockLen):((blockNum+1)*blockLen)])
		fmt.Printf("%v \n", encryptionWrapper([]byte(mapKey)))
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
func createAttackStrings(blockLen int) [][]byte {
	attackPrefixes := make([][]byte, blockLen)
	for i := 0; i < blockLen; i++ {
		attackPrefixes[i] = make([]byte, i)
	}

	attackStrings := make([][]byte, len(attackPrefixes))

	for pos, str := range attackPrefixes {
		attackStrings[pos] = encryptionWrapper(str)
	}

	return attackStrings
}

func attack() {
	// create the set of attack strings
	// iteratively decrypt the secret message byte by byte

	blockSize := 16
	blockNum := 0
	attackStringSelector := 15

	known := make([]byte, 0)

	attackStrings := createAttackStrings(blockSize)

	for range attackStrings[0] {
		// decrypt the first byte
		attackMap := buildAttackMap(blockSize, blockNum, known)

		attackedBlock := string(attackStrings[attackStringSelector][(blockNum * blockSize):((blockNum + 1) * blockSize)])
		byt, ok := attackMap[attackedBlock]
		if ok {
			known = append(known, byt)
			fmt.Println(byt)
		} else {
			fmt.Println("AttackedBlock:")
			fmt.Printf("%v, \n", []byte(attackedBlock))
			//			for key := range attackMap {
			//				fmt.Printf("%v \n", []byte(key))
			//			}
		}

		if attackStringSelector == 0 {
			attackStringSelector = 15
			blockNum++
			fmt.Println("BlockNum: ")
			fmt.Println(blockNum)
		} else {
			attackStringSelector--
		}
		fmt.Printf("%s \n", known)
	}
}

func Main() {
	attack()
}
