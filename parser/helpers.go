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

func loadPrefixesFuncs(p *Parser) map[token.TokenType]prefixParseFn {
	return map[token.TokenType]prefixParseFn{
		token.Ident: p.parseIdentifier,
		token.Int:   p.parseIntegerLiteral,
		token.Bang:  p.parsePrefixExpression,
		token.Minus: p.parsePrefixExpression,
		token.True:  p.parseBooleanExpression,
		token.False: p.parseBooleanExpression,
	}
}

func loadInfixFuncs(p *Parser) map[token.TokenType]infixParseFn {
	return map[token.TokenType]infixParseFn{
		token.Plus:     p.parseInfixExpression,
		token.Minus:    p.parseInfixExpression,
		token.Slash:    p.parseInfixExpression,
		token.Eq:       p.parseInfixExpression,
		token.NotEq:    p.parseInfixExpression,
		token.Lt:       p.parseInfixExpression,
		token.Gt:       p.parseInfixExpression,
		token.Asterisk: p.parseInfixExpression,
	}
}

func noPrefixParseFnError(p *Parser, t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s", t)
	p.errors = append(p.errors, msg)
}

func peekPrecedence(p *Parser) int {
	if precedence, ok := precedences[p.peekToken.Type]; ok {
		return precedence
	}

	return Lowest
}

func curPrecedence(p *Parser) int {
	if precedence, ok := precedences[p.curToken.Type]; ok {
		return precedence
	}

	return Lowest
}
