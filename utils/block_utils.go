package utils

import "errors"

type ECBBlock struct {
	BlockLen     int
	CurrentBlock []byte
	Text         []byte
}

//SelectBlock takes the block number (starting at block 0), returns that block, and sets current block.
func (b *ECBBlock) SelectBlock(blockNum int) ([]byte, error) {
	// really I'm just tired of of typing in text[(blockNum*blockLen):((blockNum+1)*blockLen)] again and again
	blockStart := blockNum * b.BlockLen
	blockEnd := (blockNum + 1) * b.BlockLen
	if blockEnd > len(b.Text) {
		return nil, errors.New("Requested block does not exist or is incomplete")
	}
	b.CurrentBlock = b.Text[blockStart:blockEnd]
	return b.CurrentBlock, nil
}
