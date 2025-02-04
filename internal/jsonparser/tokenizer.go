package jsonparser

import (
	"strconv"
	"strings"
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

const (
	JsonBooleanTrue  = "true"
	JsonBooleanFalse = "false"
	JsonNull         = "null"
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
	TokObject
	TokArray
	TokInvalid
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
	TokObject:       "TokObject",
	TokArray:        "TokArray",
	TokInvalid:      "TokInvalid",
}

func (t TokenType) String() string {
	return TokenTypeMap[t]
}

type Location struct {
	Line int
	Col  int
}

type Token struct {
	TokType  TokenType
	TokValue any // for string, boolean & number token only
	Loc      Location
}

var SYNTAX_TOKEN_MAP = map[rune]TokenType{
	JSON_LEFTBRACE:    TokLeftBrace,
	JSON_RIGHTBRACE:   TokRightBrace,
	JSON_LEFTBRACKET:  TokLeftBracket,
	JSON_RIGHTBRACKET: TokRightBracket,
	JSON_COMMA:        TokComma,
	JSON_COLON:        TokColon,
}

// top level API
func FromString(str string) (map[string]any, error) {
	// TODO: call tokenizer and parser
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

func tokenizeString(index int, runes []rune) (Token, int) {
	if runes[index] != JSON_QUOTE {
		return Token{}, -1
	}

	start := index + 1
	for i := start; i < len(runes); i++ {
		if runes[i] == JSON_QUOTE {
			return Token{TokType: TokString, TokValue: string(runes[start:i]), Loc: Location{Col: i}}, i - index + 1
		}
	}
	// TODO: what would happen here?
	return Token{}, -1
}

func tokenizeNumber(index int, runes []rune) (Token, int) {
	runeRead := 0

	for i := index; i < len(runes); i++ {
		cur := runes[i]

		if !isJsonNumber(cur) {
			break
		} else {
			runeRead++
		}
	}

	numberString := string(runes[index : index+runeRead])
	var value any
	var err error

	// this is float
	if strings.ContainsAny(numberString, ".eE") {
		value, err = strconv.ParseFloat(numberString, 64)
	} else {
		// integer
		value, err = strconv.ParseInt(numberString, 10, 64)
	}
	if err != nil {
		// TODO: what happen when can't parse
		return Token{TokType: TokNumber, TokValue: nil}, runeRead
	}
	return Token{TokType: TokNumber, TokValue: value}, runeRead
}

func isJsonNumber(r rune) bool {
	for _, numberToken := range JSON_NUMBER {
		if numberToken == r {
			return true
		}
	}
	return false
}

func TokenizeBoolean(index int, input []rune) (Token, int) {
	inputLen := len(input[index:])

	trueTokenLen := len(JsonBooleanTrue)
	falseTokenLen := len(JsonBooleanFalse)

	if inputLen >= trueTokenLen && string(input[index:index+trueTokenLen]) == JsonBooleanTrue {
		return Token{TokType: TokBoolean, TokValue: true}, trueTokenLen
	}
	if inputLen >= falseTokenLen && string(input[index:index+falseTokenLen]) == JsonBooleanFalse {
		return Token{TokType: TokBoolean, TokValue: false}, falseTokenLen
	}
	return Token{}, -1
}

func tokenizeNull(index int, input []rune) (Token, int) {
	inputLen := len(input[index:])
	nullTokenLen := len(JsonNull)

	if inputLen >= nullTokenLen && string(input[index:index+nullTokenLen]) == JsonNull {
		return Token{TokType: TokNull}, nullTokenLen
	}
	return Token{}, -1
}

func Tokenize(str string) []Token {
	// TODO: if pre-allocating the token slice is possible
	tokens := make([]Token, 0)
	input := []rune(str)

	for index := 0; index < len(input); {
		currentRune := input[index]

		if numToken, tokenRead := tokenizeNumber(index, input); tokenRead > 0 {
			tokens = append(tokens, numToken)
			index = index + tokenRead
		} else if stringToken, tokenRead := tokenizeString(index, input); tokenRead > 0 {
			tokens = append(tokens, stringToken)
			index = index + tokenRead
		} else if nullToken, tokenRead := tokenizeNull(index, input); tokenRead > 0 {
			tokens = append(tokens, nullToken)
			index = index + tokenRead
		} else if boolToken, tokenRead := TokenizeBoolean(index, input); tokenRead > 0 {
			tokens = append(tokens, boolToken)
			index = index + tokenRead
		} else if isJsonSyntax, tokenType := isJsonSyntax(currentRune); isJsonSyntax == true {
			tokens = append(tokens, Token{TokType: tokenType})
			index++
		} else {
			// whitespace
			index++
		}
	}
	return tokens
}
