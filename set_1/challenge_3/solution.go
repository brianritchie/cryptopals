package main 

import (
	"fmt"
	"encoding/hex"
)

func main() {

	//Variables

	encrypted_hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	xor_byte_string, err := hex.DecodeString(encrypted_hex)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The XOR'ed byte string is --> ", string(xor_byte_string))
	}

	// Commence XOR Operations

	/// Empty Slice to store temporary bytes

	store := make([]byte, len(xor_byte_string))

	// c = a ^ b, thus b = c ^ a. Given a is a single character - between 1 to 255, we can brute force

	var b int
	for b = 255; b > 0; b-- {
		for a := 0; a < len (xor_byte_string); a++ {

			store[a] = xor_byte_string[a] ^ byte(b)
		}

		fmt.Println("b is", b, "and the plaintext is --> ",string(store))
	}

}