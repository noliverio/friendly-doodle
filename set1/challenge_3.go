package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	input_string := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	input_deco := hex.DecodeString(input_string)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input_deco)
}
