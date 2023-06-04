package object

type BuiltinFunction func(args ...Object) Object

type Builting struct {
	Fn BuiltinFunction
}

func (b *Builting) Type() ObjectType {
	return BuiltingObj
}

func (b *Builting) Inspect() string {
	return "builting function"
}
