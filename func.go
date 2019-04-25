package assert

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var goPath = build.Default.GOPATH
var srcDir = filepath.Join(goPath, "src")

func funcCaller() string {
	_, file, line, _ := runtime.Caller(callDepth)
	return strings.TrimPrefix(fmt.Sprintf("%s:%d ", file, line), fmt.Sprintf("%s%s", srcDir, string(os.PathSeparator)))
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

func FileExists(path string) bool {
	_, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		NotNil(err)
	}
	return true
}

func DirExists(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		NotNil(err)
	}
	return info.IsDir()
}
