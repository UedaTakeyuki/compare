package compare

import (
	"runtime"
	"testing"
)

func Compare(t *testing.T, actual interface{}, expected interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)

	if actual != expected {
		t.Errorf("got: %v, %T\nwant: %v, %T\ncalled from %s on %s:%d", actual, actual, expected, expected,
			f.Name(), file, line)
	}
}

func CompareNotEqual(t *testing.T, actual interface{}, expected interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)

	if actual == expected {
		t.Errorf("got: %v, %T\nnot want: %v, %T\ncalled from %s on %s:%d", actual, actual, expected, expected,
			f.Name(), file, line)
	}
}
