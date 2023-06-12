package object

import "hash/fnv"

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

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}
