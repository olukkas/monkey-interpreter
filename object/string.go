package object

type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{value}
}

func (s *String) Type() ObjectType {
	return StringObj
}

func (s *String) Inspect() string {
	return s.Value
}
