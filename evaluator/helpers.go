package evaluator

import "monkey/object"

func isThruthy(obj object.Object) bool {
	switch obj {
	case Null:
		return false
	case True:
		return true
	case False:
		return false
	default:
		return true
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ErrorObj
	}

	return false
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}

	return False
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return object.NewErrorObject("not a function: %s", fn.Type())
	}

	extendedEnv := extendsFucntionEnv(function, args)
	evaluated := Eval(function.Body, extendedEnv)

	return unwrapReturnValue(evaluated)
}

func operandsSameType(left, right object.Object) bool {
	return left.Type() == right.Type()
}

func operandsAre(left, right object.Object, tp object.ObjectType) bool {
	return operandsSameType(left, right) && left.Type() == tp
}

func extendsFucntionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		env.Set(param.Value, args[i])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}
