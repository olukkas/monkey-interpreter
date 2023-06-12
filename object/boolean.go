package object

import "fmt"

type Boolean struct {
	Value bool
}

func NewBoolean(value bool) *Boolean {
	return &Boolean{Value: value}
}

func (b *Boolean) Type() ObjectType {
	return BooleanObj
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{
		Type:  b.Type(),
		Value: value,
	}
}
