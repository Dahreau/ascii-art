package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please enter a text to output as ASCII")
		os.Exit(0)
	} else if len(os.Args) > 4 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	content := getContent()
	var contentTable [][]byte
	if len(os.Args) == 3 && os.Args[2] == "thinkertoy" {
		contentTable = bytes.Split(content, []byte{13, 10})
	} else {
		contentTable = bytes.Split(content, []byte{10})

	}
	textToTransform := strings.Split(os.Args[1], `\n`)
	for k := 0; k < len(textToTransform); k++ {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			for j := 0; j < len(textToTransform[k]); j++ {
				startChar := ((int(textToTransform[k][j]) - 32) * 9) + 1
				lines[i] += string(contentTable[startChar+i])
			}
		}
		if lines[0] != "" {
			result := strings.Join(lines, "\n")
			fmt.Print(result)
		}
		if k != len(textToTransform) {
			fmt.Println()
		}
	}
}
