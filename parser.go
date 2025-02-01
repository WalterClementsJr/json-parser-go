package main

import (
	"fmt"
)

// syntax
const (
	JSON_LEFTBRACE    = '{'
	JSON_RIGHTBRACE   = '}'
	JSON_LEFTBRACKET  = '['
	JSON_RIGHTBRACKET = ']'
	JSON_COMMA        = ','
	JSON_COLON        = ':'
	JSON_QUOTE        = '"'
)

var JSON_NUMBER = [...]rune{
	'0',
	'1',
	'2',
	'3',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'-',
	'e',
	'E',
	'.',
	'+',
}

var JSON_SYNTAX = [...]rune{
	JSON_LEFTBRACE,
	JSON_RIGHTBRACE,
	JSON_LEFTBRACKET,
	JSON_RIGHTBRACKET,
	JSON_COMMA,
	JSON_COLON,
	JSON_QUOTE,
}
var JSON_WHITESPACE = [...]rune{' ', '\t', '\b', '\n', '\r'}

type TokenType int

const (
	TokLeftBrace TokenType = iota
	TokRightBrace
	TokLeftBracket
	TokRightBracket
	TokColon
	TokComma
	TokString
	TokBoolean
	TokNumber
	TokNull
)

var TokenTypeMap = map[TokenType]string{
	TokLeftBrace:    "TokLeftBrace",
	TokRightBrace:   "TokRightBrace",
	TokLeftBracket:  "TokLeftBracket",
	TokRightBracket: "TokRightBracket",
	TokColon:        "TokColon",
	TokComma:        "TokComma",
	TokString:       "TokString",
	TokBoolean:      "TokBoolean",
	TokNumber:       "TokNumber",
	TokNull:         "TokNull",
}

func (t TokenType) String() string {
	return TokenTypeMap[t]
}

type Token struct {
	TokType  TokenType
	TokValue any // for string & number token only
}

// func (t Token) String() string {
// 	var val string
// 	if t.TokValue != nil {
// 		switch t.TokValue.(type) {
// 		case int32:
// 			return string(t.TokValue)
// 		case string:
// 			return
// 		}
// 		val = t.TokValue.(string)
// 	} else {
// 		val = "<nil>"
// 	}
// 	return fmt.Sprintf("Token [type: %s, value: %s]", t.TokType, val)
// }

var SYNTAX_TOKEN_MAP = map[rune]TokenType{
	JSON_LEFTBRACE:    TokLeftBrace,
	JSON_RIGHTBRACE:   TokRightBrace,
	JSON_LEFTBRACKET:  TokLeftBracket,
	JSON_RIGHTBRACKET: TokRightBracket,
	JSON_COMMA:        TokComma,
	JSON_COLON:        TokColon,
	// JSON_QUOTE:        TokQuote,
}

// top level API
func FromString(str string) (map[string]any, error) {
	// TODO
	return nil, nil
}

func isWhiteSpace(r rune) bool {
	for _, syn := range JSON_WHITESPACE {
		if syn == r {
			return true
		}
	}
	return false
}

func isJsonSyntax(r rune) (bool, TokenType) {
	for _, syn := range JSON_SYNTAX {
		if syn == r {
			correspondingToken, ok := SYNTAX_TOKEN_MAP[syn]

			if ok {
				return true, correspondingToken
			}
		}
	}
	return false, -1
}

func lexString(index int, runes []rune) (Token, bool) {
	if runes[index] != JSON_QUOTE {
		return Token{}, false
	}

	start := index + 1
	for i := start; i < len(runes); i++ {
		// end string
		if runes[i] == JSON_QUOTE {
			return Token{TokType: TokString, TokValue: runes[start:i]}, true
		}
	}
	// TODO: what would happen here?
	return Token{}, false
}

func lexNumber(index int, runes []rune) (Token, bool) {
	foundNumber := false

	for i := index; i < len(runes); i++ {
		cur := runes[i]

		if !isNumber(cur) {
			return Token{TokType: TokNumber, TokValue: runes[index:i]}, foundNumber
		} else {
			foundNumber = true
		}
	}
	return Token{}, false
}

func isNumber(r rune) bool {
	for _, numberToken := range JSON_NUMBER {
		if numberToken == r {
			return true
		}
	}
	return false
}

// lexing
func lex(str string) []Token {
	// TODO: how to pre allocate the token slice
	fmt.Println(str)

	lexResult := make([]Token, 0)
	input := []rune(str)

	// lexing
	for index := 0; index < len(input); {
		r := input[index]

		if numToken, ok := lexNumber(index, input); ok == true {
			lexResult = append(lexResult, numToken)
			index = index + len(numToken.TokValue.([]rune))
		} else if stringToken, ok := lexString(index, input); ok == true {
			lexResult = append(lexResult, stringToken)
			// since we don't include the quote, we have to increment index to pass the closing quote
			index = index + len(stringToken.TokValue.([]rune)) + 1
			// } else if isWhiteSpace := isWhiteSpace(r); isWhiteSpace == true {
			// TODO: we don't care about whitespace?
			// 	index++
		} else if isJsonSyntax, tokenType := isJsonSyntax(r); isJsonSyntax == true {
			lexResult = append(lexResult, Token{TokType: tokenType})
		}
		index++
	}
	fmt.Println("parse result", lexResult)
	return lexResult
}

func main() {
	json := "{\"name\": \"Jason\", \"age\": -10, \"school\": { \"name\": \"Bekerly\"}}"

	lex(json)

	// m, err := FromString(json)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("parsed resutl", m)
}
