package utils

import (
	"encoding/base64"
	//	"fmt"
	"reflect"
	"testing"
)

// Tests for encryption functions
//TODO tests for when encrypting messages successfully

func TestEBCEncrypt(t *testing.T) {
	message := []byte("1234")
	key := []byte("1234567890abcdef")
	args := map[string][]byte{"blockMode": []byte("ECB")}
	ciphertext, err := Encrypt(message, key, args)
	encodedResult := base64.StdEncoding.EncodeToString(ciphertext)
	expectedResult := "8U8ZdxsyiATvLaJ6eHIq4Q=="
	if err != nil {
		t.Errorf("Encrypt failed recieved unexpected error %v", err)
	}
	if encodedResult != expectedResult {
		t.Errorf("Encrypt failed expected %v got %v", expectedResult, encodedResult)
	}

}

func TestCBCEncrypt(t *testing.T) {
	message := []byte("123456789012345678901234567890")
	key := []byte("1234567890abcdef")
	args := map[string][]byte{
		"blockMode": []byte("CBC"),
		"iv":        []byte("abcdef0123456789"),
	}
	ciphertext, err := Encrypt(message, key, args)
	encodedResult := base64.StdEncoding.EncodeToString(ciphertext)
	expectedResult := "nfxP6BscaPbvH2lmsfiDdaWcKN/z4f8pg1NBOiWBJBc="
	if err != nil {
		t.Errorf("Encrypt failed recieved unexpected error %v", err)
	}
	if encodedResult != expectedResult {
		t.Errorf("Encrypt failed expected %v got %v", expectedResult, encodedResult)
	}

}

// Tests for Encrypt function erroring
//add test for invalid keysize for CBC block mode
func TestEncryptNoBlockMode(t *testing.T) {
	message := []byte("1234")
	key := []byte("1234")
	args := map[string][]byte{"iv": []byte("CBC")}
	_, err := Encrypt(message, key, args)
	if err == nil {
		t.Errorf("Encrypt failed expected error")
	}
}

func TestEncryptUnsupportedBlockMode(t *testing.T) {
	message := []byte("1234")
	key := []byte("1234")
	args := map[string][]byte{"blockMode": []byte("CTR")}
	_, err := Encrypt(message, key, args)
	if err == nil {
		t.Errorf("Encrypt failed expected error")
	}
}

func TestECBEncryptInvalidKey(t *testing.T) {
	message := []byte("1234")
	key := []byte("1234")
	args := map[string][]byte{"blockMode": []byte("ECB")}
	_, err := Encrypt(message, key, args)
	if err == nil {
		t.Errorf("Encrypt failed expected error")
	}
}

func TestCBCEncryptNoIV(t *testing.T) {
	message := []byte("1234")
	key := []byte("1234")
	args := map[string][]byte{"blockMode": []byte("CBC")}
	_, err := Encrypt(message, key, args)
	if err == nil {
		t.Errorf("Encrypt failed expected error")
	}
}

// Tests for padding functions
func TestPad(t *testing.T) {
	// Message is 6 characters long, block length is 16, need to pad with 10 10's (/x0a)
	messageBlock := []byte("Yellow")
	blockLength := 16
	expectedResult := []byte("Yellow\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a")
	result := Pad(messageBlock, blockLength)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Pad failed on block, expected %v, length: %d \ngot %v \n lenght: %d", expectedResult, len(expectedResult), result, len(result))
	}
}

func TestPadMessage(t *testing.T) {
	// Message is 89 characters long, block length is 16, need to pad with 6 6's (/x06)
	messageBlock := []byte("test message test message test message test message test message test message test message")
	blockLength := 16
	expectedResult := []byte("test message test message test message test message test message test message test message\x06\x06\x06\x06\x06\x06")
	result := PadMessage(messageBlock, blockLength)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("PadMessage failed on message, expected %v, length: %d \ngot %v \n lenght: %d", expectedResult, len(expectedResult), result, len(result))
	}
}

func TestPadMessageNoAdditionalPadding(t *testing.T) {
	messageBlock := []byte("0123456789abcdef")
	blockLength := 16
	expectedResult := messageBlock
	result := PadMessage(messageBlock, blockLength)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("PadMessage failed on message with no additional padding, expected %v, length: %d \ngot %v \n lenght: %d", expectedResult, len(expectedResult), result, len(result))
	}
}

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
