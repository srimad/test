package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <file_path>")
		return
	}

	inpfl, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	} 
	defer inpfl.Close()

	return
}
