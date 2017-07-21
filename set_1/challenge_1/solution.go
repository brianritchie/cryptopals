package main 

import (
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

func main () {

	//Variables
	hex_string := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	byte_string, err := hex.DecodeString(hex_string)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The ASCII byte conversion is --> ", string(byte_string))
	}

	base64_encode := base64.StdEncoding.EncodeToString([]byte(byte_string))
	fmt.Println("The Base64 encoding of the above is --> ", base64_encode)
}