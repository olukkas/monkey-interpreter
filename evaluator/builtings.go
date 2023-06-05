package evaluator

import "monkey/object"

var builtings = map[string]*object.Builting{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewErrorObject("wrong number of arguments. got=%d, wat=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return object.NewIntegerObject(int64(len(arg.Value)))

			default:
				return object.NewErrorObject(
					"argument to `len` not suported, got %s",
					args[0].Type(),
				)
			}
		},
	},
}
