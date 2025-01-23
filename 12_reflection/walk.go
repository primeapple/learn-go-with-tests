package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
    case reflect.Chan:
        for {
            if v, ok := val.Recv(); ok {
                walkValue(v)
            } else {
                break
            }
        }
    case reflect.Func:
        valFnResult := val.Call(nil)
        for _, res := range valFnResult {
            walkValue(res)
        }
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
