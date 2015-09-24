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
	return fmt.Sprintf("panic: %#v\n\n%s", e.Value, e.Stack)
}

// CatchOnly captures a panic and its stack trace, but only if the value given
// to panic is of the passed in type (you should generally pass in the zero
// value of the error you want to catch).
func CatchOnly(f func(), errProto interface{}) (err error) {
	defer func() {
		panicVal := recover()
		if panicVal == nil {
			return
		}

		stackBuf := make([]byte, 65535)
		stackBuf = stackBuf[:runtime.Stack(stackBuf, false)]
		err = Panic{string(stackBuf), panicVal}

		if errProto == nil || reflect.TypeOf(panicVal) == reflect.TypeOf(errProto) {
			return
		}

		panic(err)
	}()

	f()
	return
}

// Catch captures a panic and its stack trace as an error.
func Catch(f func()) error {
	return CatchOnly(f, nil)
}
