package challenge_11

// Assuming an attacker controlled string, detect if EBC or CBC mode is used.

import (
	"crypto/aes"
	"friendly-doodle/set2/challenge_10"
	"friendly-doodle/set2/challenge_9"
	"math/rand"
	"reflect"

	"fmt"
	"friendly-doodle/utils"
)

func EcbEncrypt(src []byte, key []byte) []byte {
	keyLength := len(key)
	block, err := aes.NewCipher(key)
	utils.Check(err)
	numberOfBlocks := len(src) / keyLength

	output := make([]byte, len(src))
	for i := 0; i < numberOfBlocks; i++ {
		block.Encrypt(output[(i*16):((i+1)*16)], src[(i*16):((i+1)*16)])
	}

	return output
}

func encrypt(msg []byte) []byte {

	randomizedMsg := randomPadding(msg)
	paddedMsg := challenge_9.PadMessage(randomizedMsg, 16)
	cipherText := make([]byte, len(paddedMsg))

	// use Ecb when encryptionMethod is 0 and CBC when it is 1
	key := make([]byte, 16)
	_, err := rand.Read(key)
	utils.Check(err)
	encryptionMethod := rand.Intn(2)
	if encryptionMethod == 0 {
		cipherText = EcbEncrypt(paddedMsg, key)
	} else if encryptionMethod == 1 {
		iv := make([]byte, 16)
		cipherText = challenge_10.CbcEncrypt(paddedMsg, iv, key)

	}

	return cipherText
}

func randomPadding(src []byte) []byte {
	prepend := make([]byte, 5+rand.Intn(6))
	append := make([]byte, 5+rand.Intn(6))
	padded := make([]byte, (len(src) + len(prepend) + len(append)))

	_, err := rand.Read(prepend)
	utils.Check(err)

	_, err = rand.Read(append)
	utils.Check(err)
	for i, value := range prepend {
		padded[i] = value
	}
	for i, value := range src {
		padded[i+len(prepend)] = value
	}
	for i, value := range append {
		padded[i+len(prepend)+len(src)] = value
	}

	return padded
}

func detectBlockMode(cipherText []byte) string {
	// I am assuming a 16 byte key, and my 43 byte repeating char string
	// With these two assumptions blocks 2 and 3 will be the same in ebc.
	// So just look for that...
	blockOne := cipherText[16:32]
	blockTwo := cipherText[32:48]
	var blockMode string
	if reflect.DeepEqual(blockOne, blockTwo) {
		blockMode = "ECB"
	} else {
		blockMode = "CBC"
	}

	return blockMode

}

func Main() {
	// using a 16 byte key, a random number bewteen 5 and 10 as a prepend, and PKCS#7 padding a string of 43 repeating bytes
	// can force a repeated identical block in ECB mode.
	// 43 because w/o random append/prepend would require 32 chars for two identical blocks.
	// With random 5-10 byte prepend first block will be different, so pad it out with 6-11 bytes.
	// Then I will be able to reliably look at blocks 2 and 3
	// |-> I probably could go less than 43 and apply some hueristic is the stw strings are xyz closeness then ecb...
	message := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	// Verify my solution is working by running a few times and verifying output
	for i := 0; i < 100; i++ {
		fmt.Println("iter:")
		cipherText := encrypt(message)
		for i := 0; i < 4; i++ {
			fmt.Println(cipherText[i*16 : (i+1)*16])
		}
		fmt.Println(detectBlockMode(cipherText))
	}
}
