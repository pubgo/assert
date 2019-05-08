package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

//var goPath = build.Default.GOPATH
//var srcDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "src"), string(os.PathSeparator))
//var modDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "pkg", "mod"), string(os.PathSeparator))

func funcCaller() string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller"
	}

	//return strings.TrimPrefix(strings.TrimPrefix(_f, srcDir), modDir)
	return fmt.Sprintf("%s:%d ", file, line)
}

func If(b bool, t, f interface{}) interface{} {
	if b {
		if _t, ok := t.(FnT); ok {
			return _t()[0].Interface()
		}
		return t
	}

	if _f, ok := f.(FnT); ok {
		return _f()[0].Interface()
	}
	return f
}

func IsAllNil(args ...interface{}) bool {
	for _, _a := range args {
		if !IsNil(_a) {
			return false
		}
	}
	return true
}

func IsAllNotNil(args ...interface{}) bool {
	for _, _a := range args {
		if IsNil(_a) {
			return false
		}
	}
	return true
}

func IfEquals(args ...interface{}) bool {
	if len(args) < 2 {
		return true
	}

	if IsAllNil(args...) {
		return true
	}

	_t := args[0]
	if IsNil(_t) {
		return false
	}

	for _, _a := range args[1:] {
		if IsNil(_t) {
			return false
		}

		if _t != _a {
			return false
		}
	}
	return true
}

func IfIn(a interface{}, args ...interface{}) bool {
	if IsNil(a) == !IsAllNotNil(args...) {
		return true
	}

	if IsNil(a) {
		return false
	}

	_a := reflect.TypeOf(a).Kind()

	for _, arg := range args {
		if IsNil(arg) {
			return false
		}

		if _a == reflect.TypeOf(arg).Kind() {
			return true
		}
	}
	
	return false
}

func IfNotIn(a interface{}, args ...interface{}) bool {
	return !IfIn(a, args...)
}

func ToInt(p string) int {
	r, err := strconv.Atoi(p)
	SWrap(err, "can not convert %s to int,error(%s)", p, err)
	return r
}
