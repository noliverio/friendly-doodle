package utils

import "testing"

//Tests for compareBlock
func testCompareBlockUnequalLengths(t *testing.T) {
	block1 := []byte("0123456789abcdef")
	block2 := []byte("0123456789")
	result, err := compareBlock(block1, block2)
	if result {
		t.Errorf("Unequal length blocks %s and %s marked as same", block1, block2)
	}
	if err == nil {
		t.Errorf("Failed to return error for block of unequal length")
	}
}
func testCompareBlockDifferentBlocks(t *testing.T) {
	block1 := []byte("0123456789abcdef")
	block2 := []byte("abcdef0123456789")
	result, err := compareBlock(block1, block2)
	if err != nil {
		t.Errorf("returned error for block of equal length")
	}
	if result {
		t.Errorf("Different blocks %s and %s marked as same", block1, block2)
	}
}
func testCompareBlockEqualBlocks(t *testing.T) {
	block1 := []byte("0123456789abcdef")
	block2 := []byte("0123456789abcdef")
	result, err := compareBlock(block1, block2)
	if err != nil {
		t.Errorf("returned error for block of equal length")
	}
	if result {
		t.Errorf("Same blocks %s and %s marked as different", block1, block2)
	}
}
