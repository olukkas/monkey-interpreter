package parser

import (
	"fmt"
	"monkey/token"
)

func curTokenIs(p *Parser, t token.TokenType) bool {
	return p.curToken.Type == t
}

func peekError(p *Parser, t token.TokenType) {
	msg := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		t,
		p.peekToken.Type,
	)

	p.errors = append(p.errors, msg)
}

func peekTokenIs(p *Parser, t token.TokenType) bool {
	return p.peekToken.Type == t
}

func expectPeek(p *Parser, t token.TokenType) bool {
	if peekTokenIs(p, t) {
		p.nextToken()
		return true
	}

	peekError(p, t)
	return false
}

func loadPrefixes(p *Parser) map[token.TokenType]prefixParseFn {
	return map[token.TokenType]prefixParseFn{
		token.Ident: p.parseIdentifier,
		token.Int:   p.parseIntegerLiteral,
	}
}
