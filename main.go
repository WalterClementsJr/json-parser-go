package main

import (
	"fmt"
	"io"
	"os"

	"github.com/walterclementsjr/json-parser-go/internal/jsonparser"
)

func main() {
	var reader io.Reader

	if len(os.Args) > 1 {
		// Read from file
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		// Read from STDIN
		reader = os.Stdin
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
	}

	input := string(bytes)

	result := jsonparser.Parse(jsonparser.Tokenize([]rune(input)))
	fmt.Println(jsonparser.Dump(result))
}
