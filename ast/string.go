package ast

import "monkey/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func NewStringLiteral(tok token.Token, value string) *StringLiteral {
	return &StringLiteral{
		Token: tok,
		Value: value,
	}
}

func (s *StringLiteral) expressionNode() {}

func (s *StringLiteral) String() string {
	return s.Token.Literal
}

func (s *StringLiteral) TokenLiteral() string {
	return s.Token.Literal
}
