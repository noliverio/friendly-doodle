package challenge_14

import (
	"encoding/base64"
	"fmt"
	utils "github.com/noliverio/friendly-doodle/utils"
)

var messageUnitMagicCode = "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

var secretKey = []byte("16-bit_secretkey")

var unknownPrefix = []byte("What's harder than challenge #12 about doing this? How would you overcome that obstacle?")

func Main() {
	// We can use an arbitraily large repeating block of characters to create a the same effect
	// as we used in challenge 12. With with enough repeating characters we can look for several
	// identical blocks and say that is our attack string. From there we can feed the final block
	// of the attack string the same attack function as used when there was no random uncontrolled prefix.
	cipherText := encryptionOracle([]byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	utils.PrintBlocks(cipherText, []byte("AAAAAAAAAAAAAAAA"))
	fmt.Println(utils.FindRepeatingBlocks(cipherText, 16))

}

func encryptionOracle(attackString []byte) []byte {
	messageUnit, err := base64.StdEncoding.DecodeString(messageUnitMagicCode)
	utils.Check(err)
	message := append(unknownPrefix, attackString...)
	message = append(message, messageUnit...)

	encryptedMessage, err := utils.Encrypt(message, secretKey, map[string][]byte{"blockMode": []byte("ECB")})
	utils.Check(err)
	return encryptedMessage
}
