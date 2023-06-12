package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func NewHashLiteral(tok token.Token) *HashLiteral {
	return &HashLiteral{
		Token: tok,
		Pairs: make(map[Expression]Expression),
	}
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+" => "+value.String())
	}

	out.WriteString("{\n")
	out.WriteString(strings.Join(pairs, ", \n"))
	out.WriteString("}\n")

	return out.String()
}
