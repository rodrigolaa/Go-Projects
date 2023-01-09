package main

import (
	"fmt"
	"io"
	"os"
)

// func main() {
// 	// get file from terminal
// 	inputFile := os.Args[1]
// 	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
// 	file, err := ioutil.ReadFile(inputFile)
// 	if err != nil {
// 		fmt.Printf("Could not read the file due to this %s error \n", err)
// 	}
// 	// convert the file binary into a string using string
// 	fileContent := string(file)
// 	// print file content
// 	fmt.Println(fileContent)
// }

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
	}

	io.Copy(os.Stdout, f)

}
