package utils

import "errors"

type ECBBlock struct {
	blockLen     int
	currentBlock []byte
	text         []byte
}

//SelectBlock takes the block number (starting at block 0), returns that block, and sets current block.
func (b *ECBBlock) SelectBlock(blockNum int) ([]byte, error) {
	// really I'm just tired of of typing in text[(blockNum*blockLen):((blockNum+1)*blockLen)] again and again
	blockStart := blockNum * b.blockLen
	blockEnd := (blockNum + 1) * b.blockLen
	if blockEnd > len(b.text) {
		return nil, errors.New("Requested block does not exist or is incomplete")
	}
	b.currentBlock = b.text[blockStart:blockEnd]

	return b.currentBlock, nil
}
