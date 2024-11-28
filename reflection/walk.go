package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numValues = val.NumField()
		getField = val.Field
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
