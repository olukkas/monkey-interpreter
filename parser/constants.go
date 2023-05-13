package parser

import "monkey/token"

const (
	_ int = iota
	Lowest
	Equals
	LessGreater // > or <
	Sum
	Product
	Prefix
	Call
)

var precedences = map[token.TokenType]int{
	token.Eq:       Equals,
	token.NotEq:    Equals,
	token.Lt:       LessGreater,
	token.Gt:       LessGreater,
	token.Plus:     Sum,
	token.Minus:    Sum,
	token.Slash:    Product,
	token.Asterisk: Product,
}
