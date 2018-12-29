package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestParseCookie(t *testing.T) {

	testSlice := []byte("foo=bar&baz=qux&zap=zazzle")
	result := ParseKV(testSlice)
	expectedMap := map[string]string{"foo": "bar", "baz": "qux", "zap": "zazzle"}

	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("result was incorrect, expected %s, got %s", result, expectedMap)
	}
}

//This test is done by calling ParseKV because the encodeMap function being tested can create multiple different valid results.
//This is because maps have random iteration order in Go, and for this use case we do not need to maintain any specific order.
// Note: Function has been updated to sort by keys before return, this format is no longer required of the test. Leaving because it
// Does not seem harmful
func TestEncodeMap(t *testing.T) {

	expectedMap := map[string]string{"foo": "bar", "baz": "qux", "zap": "zazzle"}
	result := ParseKV(encodeMap(expectedMap))

	if !reflect.DeepEqual(result, expectedMap) {
		t.Errorf("result was incorrect, expected %s, got %s", result, expectedMap)
	}
}

func TestProfileForWithAmpersand(t *testing.T) {
	result, err := ProfileFor("123&abc.com")
	if err == nil {
		t.Errorf("expected an error, got %s", result)
	}
	if err.Error() != errors.New("Email contains an illegal character").Error() {
		t.Errorf("expected a different error")
	}
}

func TestProfileForWithEqualSign(t *testing.T) {
	result, err := ProfileFor("123=abc.com")
	if err == nil {
		t.Errorf("expected an error, got %s", result)
	}
	if err.Error() != errors.New("Email contains an illegal character").Error() {
		t.Errorf("expected a different error")
	}
}

func TestValidProfileFor(t *testing.T) {
	result, err := ProfileFor("123@abc.com")
	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}
	expectedString := "email=123@abc.com&role=user&uid=10"

	if string(result) != expectedString {
		t.Errorf("result was incorrect, expected %s, got %s", result, expectedString)
	}

}
