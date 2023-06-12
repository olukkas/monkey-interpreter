package evaluator

import "monkey/object"

func wronElementesError(expected, length int) object.Object {
	return object.NewError("wrong number of arguments. got=%d, wat=%d", length, expected)
}

func wrongTypeError(funcName, expected, got string) object.Object {
	return object.NewError(
		"argument to `%s` must be %s, got %s",
		funcName, expected, got,
	)
}

var builtings = map[string]*object.Builting{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return wronElementesError(1, len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				strLen := len(arg.Value)
				return object.NewInteger(int64(strLen))

			case *object.Array:
				elmentsLen := len(arg.Elements)
				return object.NewInteger(int64(elmentsLen))

			default:
				return object.NewError(
					"argument to `len` not suported, got %s",
					args[0].Type(),
				)
			}
		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return wronElementesError(1, len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return wrongTypeError("first", "ARRAY", string(args[0].Type()))
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return Null
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return wronElementesError(1, len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return wrongTypeError("last", "ARRAY", string(args[0].Type()))
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if len(arr.Elements) > 0 {
				return arr.Elements[length-1]
			}

			return Null
		},
	},

	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return wronElementesError(1, len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return wrongTypeError("rest", "ARRAY", string(args[0].Type()))
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])

				return object.NewArray(newElements)
			}

			return Null
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return wronElementesError(2, len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return wrongTypeError("push", "ARRAY", string(args[0].Type()))
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return object.NewArray(newElements)
		},
	},
}
