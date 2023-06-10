package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token // the [ token
	Elements []Expression
}

func NewArray(tok token.Token) *ArrayLiteral {
	return &ArrayLiteral{
		Token:    tok,
		Elements: make([]Expression, 0),
	}
}

func (a *ArrayLiteral) expressionNode() {}

func (a *ArrayLiteral) TokenLiteral() string { return a.Token.Literal }

func (a *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
