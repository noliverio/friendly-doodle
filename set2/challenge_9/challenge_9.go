package challenge_9

import "fmt"

// This assumes that the message is only one block long.
// Not that usefull. lets rework to pad message of n blocks
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

func Main() {
	fmt.Println(PadMessage([]byte("YELLOW SUBMARINE"), 10))
}
