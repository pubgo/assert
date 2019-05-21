package assert

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go/build"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

func assertFn(fn interface{}) {
	T(IsNil(fn), "the func is nil")

	_v := reflect.TypeOf(fn)
	T(_v.Kind() != reflect.Func, "func type error(%s)", _v.String())
}

var goPath = build.Default.GOPATH
var srcDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "src"), string(os.PathSeparator))
var modDir = fmt.Sprintf("%s%s", filepath.Join(goPath, "pkg", "mod"), string(os.PathSeparator))

func funcCaller() string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return "no func caller"
	}

	return strings.TrimPrefix(strings.TrimPrefix(fmt.Sprintf("%s:%d ", file, line), srcDir), modDir)
}
