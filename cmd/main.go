package main

import "fmt"
import "github.com/walterclementsjr/json-parser-go/internal/parser"

func main() {
	json := "{\"name\": \"Jason\", \"age\": -10, \"phone\": null, \"school\": { \"name\": \"Bekerly\"}}"
	fmt.Println(json)
	tokens := parser.Lex(json)
	fmt.Println("parse result", tokens)
}

