package challenge_12

import (
	"encoding/base64"
	//	"fmt"

	"friendly-doodle/utils"

	"friendly-doodle/set2/challenge_11"
	"friendly-doodle/set2/challenge_9"
	"math/rand"
)

func encrypt(msg []byte, key []byte) []byte {

	paddedMsg := challenge_9.PadMessage(randomizedMsg, 16)
	cipherText := make([]byte, len(paddedMsg))

	cipherText = challenge_11.EcbEncrypt(paddedMsg, key)
	return cipherText
}

func createAttackSlice(cipherText) []byte {
	// could use blockMode, but won't acually use it in the attack
	//	blockMode := challenge_11.DetectBlockMode(cipherText)
	blockLength := findBlockLength()

	attackString := make([]byte, blockLength-1)
	for i := range attackString {
		attackString[i] = byte("A")
	}
}

func main() {
	secret_base64 := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	secret, err := base64.StdEncoding.DecodeString(secret_base64)
	utils.Check(err)
	key := make([]byte, 16)
	_, err = rand.Read(key)
	utils.Check(err)

	attack_slice := []byte("AAAAAAAAAAAAA")
	msg := append(attack_slice, secret...)

	encrypt(msg, key)

}
