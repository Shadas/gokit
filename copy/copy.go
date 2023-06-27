package copy

// learn with https://github.com/mohae/deepcopy/blob/master/deepcopy.go

import "reflect"

func DeepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	origin := reflect.ValueOf(src)
	result := reflect.New(origin.Type()).Elem()
	copyRecursive(origin, result)
	return result.Interface()
}

func copyRecursive(origin, target reflect.Value) {
	if origin.CanInterface() {
		if copier, ok := origin.Interface().(interface{ DeepCopy() error }); ok {
			target.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch origin.Kind() {
	case reflect.Ptr:
		originValue := origin.Elem()

		if !originValue.IsValid() {
			return
		}

		target.Set(reflect.New(originValue.Type()))
		copyRecursive(originValue, target.Elem())
	case reflect.Interface:
		if origin.IsNil() {
			return
		}
		originValue := origin.Elem()

		copyValue := reflect.New(originValue.Type()).Elem()
		copyRecursive(originValue, copyValue)
		target.Set(copyValue)
	case reflect.Struct:
		for i := 0; i < origin.NumField(); i++ {
			copyRecursive(origin.Field(i), target.Field(i))
		}
	case reflect.Slice:
		if origin.IsNil() {
			return
		}
		target.Set(reflect.MakeSlice(origin.Type(), origin.Len(), origin.Cap()))
		for i := 0; i < origin.Len(); i++ {
			copyRecursive(origin.Index(i), target.Index(i))
		}
	case reflect.Map:
		if origin.IsNil() {
			return
		}
		target.Set(reflect.MakeMap(origin.Type()))
		for _, key := range origin.MapKeys() {
			originValue := origin.MapIndex(key)
			copyValue := reflect.New(originValue.Type()).Elem()
			copyRecursive(originValue, copyValue)
			copyKey := DeepCopy(key.Interface())
			target.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}
	default:
		target.Set(origin)
	}
}
