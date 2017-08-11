package main

import (

	"fmt"

)

func main() {

	//Variables

	source_string := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	
	store := make([]byte, len(source_string))

	pos_one := 0
	pos_two := 1
	pos_three := 2

	key := "ICE"

	for a := 0; a < len(source_string); a++ {

		if a == pos_one {
			
			store[a] = source_string[a] ^ key[0]
			pos_one = pos_one + 3
			
		} else if a == pos_two {
			
			store[a] = source_string[a] ^ key[1]
			pos_two = pos_two + 3
			
		} else if a == pos_three {
			
			store[a] = source_string[a] ^ key[2]
			pos_three = pos_three + 3
			
		}

	}

	fmt.Println("The original string to be encrypted is -->", source_string)

	fmt.Println("The solution is as follows -->")
	fmt.Printf("%x\n", string(store))

}