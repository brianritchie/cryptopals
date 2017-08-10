package main

import (

	"fmt"
	"bufio"
	"os"
	"encoding/hex"
	"regexp"
)

// This function will check for errors and output if one occurs
func check(e error) {

	if e != nil {
		panic(e)
	}
	
}

func main() {
	
	//Open the file and check if there are any errors that occur
	source_file, err := os.Open("4.txt")
	check(err)

	// Lets now read the file line by line
	file_scanner := bufio.NewScanner(source_file)
	for file_scanner.Scan() {
		hex_string := file_scanner.Text()
		

		//Decoding to bytes
		decoded_byte_string, err := hex.DecodeString(hex_string)
		check(err)
		store := make([]byte, len(decoded_byte_string))
		var b int
		for b = 255; b > 0; b-- {
			for a := 0; a < len (decoded_byte_string); a++ {

				store[a] = decoded_byte_string[a] ^ byte(b)
			}

			
			common_english_characters := regexp.MustCompile("E|T|A|O|I|N|S|H|R|D|L|U")

			match_check := common_english_characters.FindAllStringIndex(string(store), -1)
			
			
			if (len(match_check)) > 15 {
				fmt.Println("The string reviewed is -->", hex_string)
				fmt.Println("b is", b, "and the plaintext is --> ",string(store))
				fmt.Println("The amount of matches to the English Language Common characters is -->", len(match_check))
			}
		}

		
	}

	
	// Close the file after usage
	source_file.Close()
}