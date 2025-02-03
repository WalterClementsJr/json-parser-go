package parser

import (
	"testing"
)

func TestTokenizeNumber(t *testing.T) {
	input := "-1"
	tokResult, tokRead := tokenizeNumber(0, []rune(input))

	if tokResult.TokType != TokNumber {
		t.Errorf("Wrong tokenized result, expected number token, got %s", tokResult.TokType)
	}
	if tokRead != 2 {
		t.Errorf("Wrong string size read, expected %d, got %d", 2, tokRead)
	}
	if tokResult.TokValue.(int64) != -1 {
		t.Errorf("Wrong token value read, expected %d, got %s", -1, tokResult.TokValue)
	}
}

func TestTokenizeString(t *testing.T) {
	input := ":\"abc\""
	tokResult, tokRead := tokenizeString(1, []rune(input))

	if tokResult.TokType != TokString {
		t.Errorf("Wrong tokenized type, expected string token, got %s", tokResult.TokType)
	}
	if tokRead != 5 {
		t.Errorf("Wrong string size read, expected 5, got %d", tokRead)
	}
	if tokResult.TokValue != "abc" {
		t.Errorf("Wrong token value read, expected %s, got %s", "abc", tokResult.TokValue)
	}
}

func TestJsonSyntax(t *testing.T) {
	isSyntax, tokenType := isJsonSyntax('[')

	if isSyntax == false {
		t.Errorf("Expected JSON bracket")
	}
	if tokenType != TokLeftBracket {
		t.Errorf("Wrong syntax detected, expected %s, got %s", TokLeftBracket, tokenType)
	}
}

func TestTokenizeBoolean(t *testing.T) {
	tokResult, read := TokenizeBoolean(0, []rune("false"))

	if tokResult.TokType != TokBoolean {
		t.Errorf("Wrong token type, expected %s, got %s", TokBoolean, tokResult.TokType)
	}
	if read != 5 {
		t.Errorf("Wrong expected %d, got %d", 5, read)
	}
	if tokResult.TokValue.(bool) != false {
		t.Errorf("Wrong token value, got %s", tokResult.TokValue)
	}
}

func TestTokenizeSyntax(t *testing.T) {
	tokens := Tokenize("{:,}")
	expectedTypes := []TokenType{TokLeftBrace, TokColon, TokComma, TokRightBrace}

	if l := len(tokens); l != 4 {
		t.Errorf("Wrong number of tokens parsed, expected %d, got %d", 4, l)
	}
	for i := range len(expectedTypes) {
		if expectedTypes[i] != tokens[i].TokType {
			t.Errorf("Wrong type parsed, expedted %s, got %s", expectedTypes[i], tokens[i].TokType)
		}
	}
}

func TestTokenizeNull(t *testing.T) {
	tokResult, read := tokenizeNull(1, []rune("1null"))

	if tokResult.TokType != TokNull {
		t.Errorf("Wrong token type, expected %s, got %s", TokNull, tokResult.TokType)
	}
	if read != 4 {
		t.Errorf("Wrong expected %d, got %d", 4, read)
	}
}

func TestTokenizer(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput []Token
	}{
		{"\"a\":{}", []Token{{TokType: TokString, TokValue: "a"}, {TokType: TokColon}, {TokType: TokLeftBrace}, {TokType: TokRightBrace}}},
		{"[{}, 10]", []Token{{TokType: TokLeftBracket}, {TokType: TokLeftBrace}, {TokType: TokRightBrace}, {TokType: TokComma}, {TokType: TokNumber, TokValue: int64(10)}, {TokType: TokRightBracket}}},
	}

	for _, tt := range testCases {
		t.Run(tt.input, func(t *testing.T) {
			tokenizerResult := Tokenize(tt.input)

			if tokLen := len(tokenizerResult); tokLen != len(tt.expectedOutput) {
				t.Errorf("Mismatched token length, expected %d, got %d", len(tt.expectedOutput), tokLen)
			}
			for i, expectedToken := range tt.expectedOutput {
				resultToken := tokenizerResult[i]

				if expectedToken.TokType != resultToken.TokType {
					t.Errorf("Wrong token type, expected %s, got %s", expectedToken.TokType, resultToken.TokType)
				}
				if expectedToken.TokValue != resultToken.TokValue {
					t.Errorf("Wrong token value, expected %s, got %s", expectedToken.TokValue, resultToken.TokValue)
				}
			}
		})
	}
}

