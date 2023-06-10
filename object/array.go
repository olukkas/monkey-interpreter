package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

func (array *Array) Type() ObjectType {
	return ArrayObj
}

func (array *Array) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range array.Elements {
		elements = append(elements, el.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
