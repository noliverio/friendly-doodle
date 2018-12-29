package utils

import (
	//"fmt"
	"reflect"
	"testing"
)

// Tests for encryption functions
//TODO

// Tests for padding functions
//TODO PadMessage, Pad

// Tests for base functions:

// Test xor
func TestXorSlices(t *testing.T) {
	slice1 := []byte("12345")
	slice2 := []byte("abcde")
	result := xorSlices(slice1, slice2)
	expectedResult := []byte{80, 80, 80, 80, 80}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("xor failed on slice, expected %v, got %v", expectedResult, result)
	}

}
