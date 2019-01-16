package utils

import (
	"crypto/aes"
	"errors"
	//"fmt"
)

// Encrypt takes a message and key, pads it appropriately, encrypts the message, and returns the cipher text
func Encrypt(msg []byte, key []byte, args map[string][]byte) ([]byte, error) {

	paddedMsg := PadMessage(msg, len(key))
	cipherText := make([]byte, len(paddedMsg))
	var err error

	_, ok := args["blockMode"]
	if !ok {
		err := errors.New("Block mode required but not specified." +
			"Please include a \"blockMode\" parameter in agrs")
		return nil, err
	}

	switch string(args["blockMode"]) {
	case "ECB":
		cipherText, err = ecbEncrypt(paddedMsg, key)
		if err != nil {
			return nil, err
		}
	case "CBC":
		_, ok = args["iv"]
		if !ok {
			err := errors.New("Initialization vector required for block mode: CBC. Please provide an \"iv\" parameter in args")
			return nil, err
		}
		cipherText, err = cbcEncrypt(paddedMsg, key, args["iv"])
		if err != nil {
			return nil, err
		}
	default:
		err := errors.New("Supported block modes are ECB and CBC please select one of these.")
		return nil, err
	}
	return cipherText, nil
}

//
func ecbEncrypt(src []byte, key []byte) ([]byte, error) {
	keyLength := len(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	numberOfBlocks := len(src) / keyLength

	output := make([]byte, len(src))
	for i := 0; i < numberOfBlocks; i++ {
		block.Encrypt(output[(i*keyLength):((i+1)*keyLength)], src[(i*keyLength):((i+1)*keyLength)])
	}

	return output, nil
}

//
func cbcEncrypt(src []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockLen := len(key)
	numberOfBlocks := len(src) / blockLen

	output := make([]byte, len(src))
	blockCipher := make([]byte, blockLen)
	// xor the first block with the initialization vector then encrypt with aes
	blockCipher = xorSlices(src[:blockLen], iv)
	block.Encrypt(output[:blockLen], blockCipher)

	// create the block chain as it were by xoring the current block with the previously encrypted block
	for i := 1; i < numberOfBlocks; i++ {
		blockCipher = xorSlices(src[(i*16):((i+1)*16)], output[((i-1)*16):(i*16)])
		block.Encrypt(output[(i*16):((i+1)*16)], blockCipher)
	}
	return output, nil

}

//Padding functions:

// Pad takes a single block and the block length, and returns a PKCS#7 padded block
func Pad(block []byte, blockLength int) []byte {
	// apply a padding to the end of the block such that the number of bytes in the padding
	// is the byte used for the padding
	bytesNeeded := blockLength - len(block)
	paddingByte := byte(bytesNeeded)
	paddedBlock := make([]byte, blockLength)

	for i := 0; i < blockLength; i++ {
		if i < len(block) {
			paddedBlock[i] = block[i]
		} else {
			paddedBlock[i] = paddingByte

		}
	}

	return paddedBlock
}

// PadMessage takes an entire message and applies an appropriate padding to it based on the block lenght
func PadMessage(message []byte, blockLength int) []byte {
	if len(message)%blockLength == 0 {
		return message
	}
	numberOfBlocks := (len(message) / blockLength) + 1
	lastBlockStart := (numberOfBlocks - 1) * blockLength
	paddedMessage := make([]byte, numberOfBlocks*blockLength)
	lastBlock := Pad(message[lastBlockStart:], blockLength)

	for i := 0; i < lastBlockStart; i++ {
		paddedMessage[i] = message[i]
	}

	for i := 0; i < len(lastBlock); i++ {
		paddedMessage[i+lastBlockStart] = lastBlock[i]
	}

	return paddedMessage

}

//Base functions:

func xorSlices(str []byte, mask []byte) []byte {
	output := make([]byte, len(str))
	for i := range str {
		blockPosition := i % len(mask)
		output[i] = str[i] ^ mask[blockPosition]
	}
	return output
}
