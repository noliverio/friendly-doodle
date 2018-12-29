package challenge_10

import (
	"crypto/aes"

	"encoding/base64"
	"fmt"
	"github.com/noliverio/friendly-doodle/utils"
	"io/ioutil"
)

// implement CBC mode for a block cipher
// similar to ECB mode, but each block is xor'd with the previous block before encryption

func xor(str []byte, priorBlock []byte) []byte {
	output := make([]byte, len(str))
	for i := range str {
		blockPosition := i % len(priorBlock)
		output[i] = str[i] ^ priorBlock[blockPosition]
	}
	return output
}

func CbcEncrypt(src []byte, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	utils.Check(err)
	blockLen := len(key)
	numberOfBlocks := len(src) / blockLen

	output := make([]byte, len(src))
	blockCipher := make([]byte, blockLen)
	// xor the first block with the initialization vector then encrypt with aes
	blockCipher = xor(src[:blockLen], iv)
	block.Encrypt(output[:blockLen], blockCipher)

	// create the block chain as it were by xoring the current block with the previously encrypted block
	for i := 1; i < numberOfBlocks; i++ {
		blockCipher = xor(src[(i*16):((i+1)*16)], output[((i-1)*16):(i*16)])
		block.Encrypt(output[(i*16):((i+1)*16)], blockCipher)
	}
	return output

}

func CbcDecrypt(src []byte, iv []byte, key []byte) []byte {
	// Want both encryption and decryption working.
	// Base decryption on encyption code
	block, err := aes.NewCipher(key)
	utils.Check(err)
	blockLen := len(key)
	numberOfBlocks := len(src) / blockLen

	output := make([]byte, len(src))
	decryptedBlock := make([]byte, blockLen)
	// Decrypt the first block then xor with the initialization vector
	block.Decrypt(decryptedBlock, src[:blockLen])
	for i, value := range decryptedBlock {
		output[i] = value ^ iv[i]
	}

	// create the block chain as it were by xoring the current block with the previously encrypted block
	for i := 1; i < numberOfBlocks; i++ {
		//		currentBlock = scr[(i*16): ((i+1)*16)]
		block.Decrypt(decryptedBlock, src[(i*16):((i+1)*16)])
		for j := 0; j < len(decryptedBlock); j++ {
			output[(i*16)+j] = decryptedBlock[j] ^ src[((i-1)*16)+j]
		}
	}
	return output

}

func Main() {
	//	initVector := []byte("\x01\x02\x03\x04\x01\x02\x03\x04\x01\x02\x03\x04\x01\x02\x03\x04")
	//	key := []byte("YELLOW SUBMARINE")
	//
	//	myStringPlain := []byte("KABUTO KABUTO KABUTO KABUTO KABUTO KABUTO KABUTO KABUTO KABUTO K")
	//
	//	myStringEncrypted := CbcEncrypt(myStringPlain, initVector, key)
	//	myStringDecrypted := CbcDecrypt(myStringEncrypted, initVector, key)
	//
	//	fmt.Printf("%s\n", myStringPlain)
	//	fmt.Printf("%s\n", myStringEncrypted)
	//	err := utils.PrintBlocks(myStringEncrypted, key)
	//	utils.Check(err)
	//
	//	fmt.Printf("%s\n", myStringDecrypted)
	//	 sweet! works on my string! now test against challenge file
	contents, err := ioutil.ReadFile("challenge_10/challenge_10.txt")
	utils.Check(err)
	src := make([]byte, base64.StdEncoding.DecodedLen(len(contents)))
	_, err = base64.StdEncoding.Decode(src, contents)
	utils.Check(err)

	intiVector := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	key := []byte("YELLOW SUBMARINE")

	decrypted := CbcDecrypt(src, intiVector, key)

	fmt.Printf("%s\n", decrypted)

	// nice works with the given ciphertext as well.
}
