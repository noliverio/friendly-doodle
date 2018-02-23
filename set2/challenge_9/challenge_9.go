package challenge_9

import "fmt"

func Pad(message []byte, blockLength int) []byte {
	// apply a padding to the end of the message such that the number of bytes in the padding
	// is the byte used for the padding
	bytesNeeded := blockLength - len(message)
	paddingByte := byte(bytesNeeded)
	paddedMessage := make([]byte, blockLength)

	for i := 0; i < blockLength; i++ {
		if i < len(message) {
			paddedMessage[i] = message[i]
		} else {
			paddedMessage[i] = paddingByte

		}
	}

	return paddedMessage
}

func Main() {
	fmt.Println(Pad([]byte("YELLOW SUBMARINE"), 20))
}
