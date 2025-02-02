package parser

import (
	"testing"
)

func TestLexNumber(t *testing.T) {
	input := "-1"
	tokResult, tokRead := lexNumber(0, []rune(input))

	if tokResult.TokType != TokNumber {
		t.Errorf("Wrong lex result, expected number token, got %s", tokResult.TokType)
	}
	if tokRead != 2 {
		t.Errorf("Wrong string size read, expected %d, got %d", 2, tokRead)
	}
	if tokResult.TokValue.(int64) != -1 {
		t.Errorf("Wrong token value read, expected %d, got %s", -1, tokResult.TokValue)
	}
}

func TestLexString(t *testing.T) {
	input := ":\"abc\""
	tokResult, tokRead := lexString(1, []rune(input))

	if tokResult.TokType != TokString {
		t.Errorf("Wrong lex type result, expected string token, got %s", tokResult.TokType)
	}
	if tokRead != 5 {
		t.Errorf("Wrong string size read, expected %d, got %d", 5, tokRead)
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

func TestLexBoolean(t *testing.T) {
	tokResult, read := lexBoolean(0, []rune("false"))

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

func TestLexNull(t *testing.T) {
	tokResult, read := lexNull(1, []rune("1null"))

	if tokResult.TokType != TokNull {
		t.Errorf("Wrong token type, expected %s, got %s", TokNull, tokResult.TokType)
	}
	if read != 4 {
		t.Errorf("Wrong expected %d, got %d", 4, read)
	}
}
