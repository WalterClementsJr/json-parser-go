package main

import (
	"fmt"

	"github.com/walterclementsjr/json-parser-go/internal/parser"
)

func main() {
	json := `[null, 10, "abc", [], {}]`
	fmt.Println(json)

	tokens := parser.Tokenize(json)
	fmt.Println("Tokenizer result", tokens)
	fmt.Println()

	parseResult := parser.Parse(tokens)

	switch res := parseResult.(type) {
	case []any:
		fmt.Println(res)
	case map[string]any:
	}
}
