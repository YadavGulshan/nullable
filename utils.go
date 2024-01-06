package nullable

import (
	"reflect"
	sys "time"
)

var null = []byte("null")
var empty = []byte("\"\"")


type time struct {
	sys.Time
}

func typeMatch(expected, actual interface{}) bool {
	if expected != nil && actual != nil {
		expectedType := reflect.TypeOf(expected)
		actualType := reflect.TypeOf(actual)

		return actualType.AssignableTo(expectedType)
	} else if expected == nil && actual == nil {
		return true
	} else {
		return false
	}
}
