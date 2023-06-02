package object

import (
	"bytes"
	"monkey/ast"
	"strings"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func NewFunctionObject(
	params []*ast.Identifier,
	body *ast.BlockStatement,
	env *Environment,
) *Function {
	return &Function{
		Parameters: params,
		Body:       body,
		Env:        env,
	}
}

func (f *Function) Type() ObjectType { return FunctionObj }

func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
