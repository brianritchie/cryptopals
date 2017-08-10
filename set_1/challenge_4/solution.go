package main

import (

	"fmt"
	"bufio"
	"os"
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
		fmt.Println(file_scanner.Text())
		
	}

	
	// Close the file after usage
	source_file.Close()
}