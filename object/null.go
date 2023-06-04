package object

type Null struct{}

func NewNullObject() *Null {
	return &Null{}
}

func (n *Null) Type() ObjectType { return NullObj }
func (n *Null) Inspect() string  { return "null" }
