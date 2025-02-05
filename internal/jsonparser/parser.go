package jsonparser

import (
	"fmt"
	"strings"
)

type Parser struct {
	sourceTokens []Token
	tokenIndex   int

	// TODO: cumulative errors
	errors []error
}

func (p *Parser) peekCurrentTok() Token {
	return p.sourceTokens[p.tokenIndex]
}

func (p *Parser) peekNextTok() Token {
	return p.sourceTokens[p.tokenIndex+1]
}

func (p *Parser) parseArray() any {
	jsonArray := make([]any, 0)

	for {
		if p.peekCurrentTok().TokType == TokRightBracket {
			p.tokenIndex++
			return jsonArray
		}
		item := p.parseValue()
		jsonArray = append(jsonArray, item)

		if p.peekCurrentTok().TokType == TokComma {
			p.tokenIndex++
			continue
		} else {
			break
		}
	}
	return jsonArray
}

func (p *Parser) parseObject() map[string]any {
	obj := make(map[string]any)

	for {
		if p.peekCurrentTok().TokType == TokRightBrace {
			// stop
			break
		}
		if p.peekCurrentTok().TokType != TokString || p.peekNextTok().TokType != TokColon {
			// TODO: push errors
			panic(fmt.Sprintf("JSON object member has wrong format, must be a value associated with a key, got %s, %s", p.peekCurrentTok().TokType, p.peekNextTok().TokType))
		}
		key := p.peekCurrentTok().TokValue.(string)

		// advance 2 tokens
		p.tokenIndex += 2

		obj[key] = p.parseValue()

		// stop parsing
		if p.tokenIndex >= len(p.sourceTokens) {
			break
		}
		if p.peekCurrentTok().TokType == TokComma {
			// continue on colon
			p.tokenIndex++
			continue
		}
	}
	return obj
}

// main parsing
func (p *Parser) parseValue() any {
	if p.tokenIndex >= len(p.sourceTokens) {
		panic("unexpected parsing error")
	}

	token := p.peekCurrentTok()
	tokenType := token.TokType

	if tokenType == TokLeftBrace {
		p.tokenIndex++
		result := p.parseObject()
		p.tokenIndex++
		return result
	} else if tokenType == TokLeftBracket {
		p.tokenIndex++
		arr := p.parseArray()
		p.tokenIndex++
		return arr
	} else if tokenType == TokNull {
		p.tokenIndex++
		return nil
	} else if (tokenType == TokBoolean || tokenType == TokString || tokenType == TokNumber) && token.TokValue != nil {
		p.tokenIndex++
		return token.TokValue
	} else {
		// TODO: accumulate errors
		panic(fmt.Errorf("unexpected token %s while parsing", token.TokType))
	}
}

func Dump(jsonValue any) string {
	builder := strings.Builder{}
	dump(jsonValue, &builder, 0)
	return builder.String()
}

func dump(jsonValue any, builder *strings.Builder, indent int) {
	indentMultiplier := 2 // 2 space per indent level

	appendIndent := func(b *strings.Builder, indentLevel int) {
		if indentLevel == 0 {
			return
		}
		b.WriteString(strings.Repeat(" ", indentLevel*indentMultiplier))
	}

	switch val := jsonValue.(type) {

	case map[string]any:
		fmt.Fprint(builder, "{\n")

		index := 0
		for k, v := range val {
			appendIndent(builder, indent+1)
			fmt.Fprintf(builder, "\"%s\": ", k)

			dump(v, builder, indent + 1)
			if index < len(val)-1 {
				fmt.Fprint(builder, ",")
			}
			fmt.Fprintln(builder)
			index++
		}
		appendIndent(builder, indent)
		fmt.Fprintf(builder, "}")

	case []any:
		fmt.Fprint(builder, "[\n")

		for index, member := range val {
			appendIndent(builder, indent+1)
			dump(member, builder, indent+1)

			if index < len(val)-1 {
				fmt.Fprint(builder, ",")
			}
			fmt.Fprintln(builder)
		}
		appendIndent(builder, indent)
		fmt.Fprint(builder, "]")
	case string:
		fmt.Fprintf(builder, "%q", val)
	case nil:
		builder.WriteString("null")
	default: // int, float, bool, null
		fmt.Fprintf(builder, fmt.Sprint(val))
	}
}

func Parse(tokens []Token) any {
	p := Parser{sourceTokens: tokens}
	return p.parseValue()
}
