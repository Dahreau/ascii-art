package main

import (
	"fmt"
	"os"
)

func getContent() []byte {
	fileName := ""
	if len(os.Args) == 3 {
		fileName = os.Args[2] + ".txt"
	} else if len(os.Args) == 2 {
		fileName = "standard.txt"
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Style error")
		os.Exit(0)
	}
	return content
}
