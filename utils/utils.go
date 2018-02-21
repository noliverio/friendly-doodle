package utils

import (
	"errors"
	"fmt"
)

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PrintBlocks(ciphertext []byte, key []byte) error {
	keyLength := len(key)
	if len(ciphertext)%keyLength != 0 {
		err := errors.New("ciphertext is not multiple of cipher block size")
		return err
	}
	numberOfBlocks := len(ciphertext) / keyLength
	for i := 0; i < numberOfBlocks; i++ {
		fmt.Println(ciphertext[(i * keyLength):((i + 1) * keyLength)])
	}

	return nil
}
