package object

import "fmt"

type Error struct {
	Message string
}

func NewError(format string, args ...any) *Error {
	return &Error{
		Message: fmt.Sprintf(format, args...),
	}
}

func (e *Error) Type() ObjectType { return ErrorObj }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
