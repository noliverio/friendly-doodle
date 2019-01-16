package utils

import (
	"bytes"
	"errors"
	//"fmt"
)

// == ECB Attack Functions ==

// = Block Functions =

// Take a ciphertext and block size, then find any repeating blocks.
// This function will return any matching blocks, so assume that
// any matches are the attack string
func FindRepeatingBlocks(ciphertext []byte, blocksize int) map[int]bool {
	blocks := make(map[int][]byte)
	matchingLines := make(map[int]bool)
	for index := 0; index < (len(ciphertext) / blocksize); index++ {
		blocks[index] = ciphertext[(index * blocksize):((index + 1) * blocksize)]
	}
	for key, value := range blocks {
		for next := key + 1; next < len(blocks); next++ {
			same, err := compareBlock(value, blocks[next])
			Check(err)
			if same {
				matchingLines[key] = true
				matchingLines[next] = true

			}
		}
	}
	return matchingLines

}

// Compare two equal sized blocks and return true if they are identical.
// If the blocks are not of the same size return an error.
func compareBlock(block1, block2 []byte) (bool, error) {
	if len(block1) != len(block2) {
		err := errors.New("Blocks are not same length")
		return false, err
	}
	if !bytes.Equal(block1, block2) {
		return false, nil
	}
	return true, nil
}

// = Decryption Attack Functions =

func BuildEcbAttackString(blockLen int)[]byte {
    return []byte("123")}

// create AttackStrings takes the secret mesage and appends bytes to the begining
//func createAttackStrings(blockLen int) []utils.ECBBlock {
//	attackPrefixes := make([][]byte, blockLen)
//	for i := 0; i < blockLen; i++ {
//		attackPrefixes[i] = make([]byte, i)
//	}
//
//	attackStrings := make([]utils.ECBBlock, len(attackPrefixes))
//
//	for pos, str := range attackPrefixes {
//		attackStrings[pos].Text = encryptionWrapper(str)
//		attackStrings[pos].BlockLen = blockLen
//	}
//
//	return attackStrings
//}

func BuildEcbAttackMap() {}
