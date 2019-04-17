package assert

var _a = &Assert{}

var P = _a.P
var Bool = _a.Bool
var True = _a.Bool

var Err = _a.Err
var MustNotError = _a.MustNotError

func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}

func IfNot(a, b interface{}) {

}

func IfEquals(args ...interface{}) bool {
	if len(args) == 0 {
		return true
	}

	_t := args[0]
	if _t == nil {
		return false
	}

	for i := 1; i < len(args); i++ {
		if args[i] == nil {
			return false
		}

		if _t != args[i] {
			return false
		}
	}
	return true
}

func IfIn(a interface{}, args ...interface{}) bool {
	for _, arg := range args {
		if a == arg {
			return true
		}
	}
	return false
}

func IfNotIn(a interface{}, args ...interface{}) bool {
	return IfIn(a, args...)
}
