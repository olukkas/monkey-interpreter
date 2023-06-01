package object

import "fmt"

type Boolean struct {
	Value bool
}

func NewBooleanObject(value bool) *Boolean {
	return &Boolean{Value: value}
}

func (b *Boolean) Type() ObjectType { return BooleanObj }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
