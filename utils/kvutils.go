package utils

import (
	"fmt"
	"sort"
	"strings"
)

// A function for parsing key value pairs in k=v format
// returns the input key value slice as a map of keys to values
func ParseKV(input []byte) map[string]string {

	var keyValueSlice []string
	returnMap := make(map[string]string)

	keyValueSlice = strings.Split(string(input), "&")

	for _, pair := range keyValueSlice {
		keyValuePair := strings.Split(pair, "=")
		if len(keyValuePair) == 2 {
			returnMap[keyValuePair[0]] = keyValuePair[1]
		}
	}

	return returnMap
}

func ProfileFor(emailAddr string) ([]byte, error) {
	if strings.Contains(emailAddr, "&") || strings.Contains(emailAddr, "=") {
		return nil, fmt.Errorf("Email contains an illegal character")
	}

	profileMap := map[string]string{"uid": "10", "role": "user"}
	profileMap["email"] = emailAddr

	return encodeMap(profileMap), nil

}

// Function updated to sort key value pairs before return due to complications in the challenge.
func encodeMap(keyValueMap map[string]string) []byte {
	var keyValueSlice []string
	var returnString string

	for key, value := range keyValueMap {
		keyValueSlice = append(keyValueSlice, (fmt.Sprintf("%s=%s", key, value)))
	}

	sort.Slice(keyValueSlice, func(i, j int) bool { return keyValueSlice[i] < keyValueSlice[j] })

	returnString = strings.Join(keyValueSlice, "&")
	return []byte(returnString)
}
