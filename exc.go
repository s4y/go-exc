package exc

import (
	"fmt"
	"reflect"
	"runtime"
)

// Panic represents an value plus a stack trace from where panicking started.
type Panic struct {
	Stack string
	Value interface{}
}

func (e Panic) Error() string {
	return fmt.Sprintf("caught panic: %#v\n\n%s", e.Value, e.Stack)
}

// CatchOnly captures a panic and its stack trace, but only if the value given
// to panic is of the passed in type (you should generally pass in the zero
// value of the error you want to catch).
func CatchOnly(f func(), errProto interface{}) (err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		if _, isRuntimeError := r.(runtime.Error); isRuntimeError {
			panic(r)
		}

		if errProto != nil && reflect.TypeOf(r) != reflect.TypeOf(errProto) {
			panic(r)
		}

		stackBuf := make([]byte, 65535)
		stackBuf = stackBuf[:runtime.Stack(stackBuf, false)]
		err = Panic{string(stackBuf), r}
		return
	}()

	f()
	return
}

// Catch captures a panic and its stack trace as an error.
func Catch(f func()) error {
	return CatchOnly(f, nil)
}
