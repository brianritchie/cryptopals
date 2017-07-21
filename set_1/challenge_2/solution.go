package main 

import (
	"fmt"
	"encoding/hex"
)

func main() {

	// Declare Variables
	hex_string_1 := "1c0111001f010100061a024b53535009181c"
	hex_string_2 := "686974207468652062756c6c277320657965"

	// Begin decoding Hex to ASCII bytes
	byte_string_1, err := hex.DecodeString(hex_string_1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The ASCII byte conversion is --> ", string(byte_string_1))
	}

	byte_string_2, err := hex.DecodeString(hex_string_2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The second ASCII byte conversion is --> ", string(byte_string_2))
	}

	//Commence XOR Operations

	///Create empty slice of type byte with a length of the first bytes string
	store := make([]byte, len(byte_string_1))
	/// Iterate through each character of the string and XOR it and store it in the slice
	for a := 0; a < len(byte_string_1); a++ {
		store[a] = byte_string_1[a] ^ byte_string_2[a]
	}

	fmt.Println("The XOR'd result is as follows:")
	//Output the String in Hexadecimal by using %x with fmt
	fmt.Printf("%x\n", string(store))

}