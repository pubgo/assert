package assert

import (
	"fmt"
	"runtime"
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
		return t
	}
	return f
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
	return !IfIn(a, args...)
}