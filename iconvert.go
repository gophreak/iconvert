package iconvert

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToString converts an interface value to a string, returns error if cannot handle the type of interface.
func ToString(value interface{}) (v string, e error) {
	switch value.(type) {
	case string:
		v = (value).(string)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(value).Int()
		v = strconv.Itoa(int(val))
	case float32:
		val := float64(value.(float32))
		v = strconv.FormatFloat(val, 'f', -1, 32)
	case float64:
		val := reflect.ValueOf(value).Float()
		v = strconv.FormatFloat(val, 'f', -1, 64)
	default:
		e = fmt.Errorf("Value %T is type %s", value, reflect.TypeOf(value).Kind())
	}

	return
}

// ToFloat converts an interface value to a float64, returns error if cannot handle the type of interface.
func ToFloat(value interface{}) (v float64, e error) {
	switch value.(type) {
	case string:
		v, e = strconv.ParseFloat((value).(string), 64)
	case int, int8, int16, int32, int64:
		val := reflect.ValueOf(value).Int()
		v = float64(val)
	case float32:
		val, _ := ToString(value) // convert float32 into string, to maintain precision
		v, e = ToFloat(val)
	case float64:
		v = reflect.ValueOf(value).Float()
	default:
		e = fmt.Errorf("Value %T is type %s", value, reflect.TypeOf(value).Kind())
	}

	return
}

// ToInt converts an interface value to a int64, returns error if cannot handle the type of interface.
func ToInt(value interface{}) (v int64, e error) {
	switch value.(type) {
	case string:
		v, e = strconv.ParseInt((value).(string), 10, 64)
	case int, int8, int16, int32, int64:
		v = reflect.ValueOf(value).Int()
	case float32:
		val := value.(float32) // convert float32 into string, to maintain precision
		v = int64(val)
	case float64:
		val := reflect.ValueOf(value).Float()
		v = int64(val)
	default:
		e = fmt.Errorf("Value %T is type %s", value, reflect.TypeOf(value).Kind())
	}

	return
}
