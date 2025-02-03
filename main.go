package main

import (
	"fmt"

	"github.com/walterclementsjr/json-parser-go/internal/parser"
)

func main() {
	json := `{
		"id": null,
		"name": "Jason",
		"age": -10,
		"phone": [],
		"school": {
		    "name": "Bekerly"
		},
		"dob": "2020"
	}`
	fmt.Println(json)

	tokens := parser.Tokenize(json)
	fmt.Println("Tokenizer result", tokens)
	fmt.Println()

	parseResult := parser.Parse(tokens)

	switch res := parseResult.(type) {
	case parser.JsonArray:
		fmt.Println(res[0])
	case parser.JsonObject:
		fmt.Println(res["id"])
		fmt.Println(res["name"])
		fmt.Println(res["age"])
		fmt.Println(res["phone"])
		fmt.Println(res["dob"])
		fmt.Println(res["school"].(parser.JsonObject)["name"])
	}
}
