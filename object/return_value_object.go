package object

type ReturnValue struct {
	Value Object
}

func NewReturnValueObject(value Object) *ReturnValue {
	return &ReturnValue{Value: value}
}

func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
