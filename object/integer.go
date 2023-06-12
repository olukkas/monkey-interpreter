package object

import "fmt"

type Integer struct {
	Value int64
}

func NewInteger(value int64) *Integer {
	return &Integer{Value: value}
}

func (i *Integer) Type() ObjectType { return IntegerObj }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
