package object

type ObjectType string

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	StringObj      = "STRING_OBJ"
	BuiltingObj    = "BUILTING_OBJ"
	ArrayObj       = "ARRAY_OBJ"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}
