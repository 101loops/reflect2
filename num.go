package reflector

import "reflect"

func IsNumber(obj interface{}) bool {
	return isNumberKind(reflect.TypeOf(obj).Kind())
}

func IsDecimalNumber(obj interface{}) bool {
	return isDecimalNumberKind(reflect.TypeOf(obj).Kind())
}

func IsUnsignedNumber(obj interface{}) bool {
	return isUnsignedNumberKind(reflect.TypeOf(obj).Kind())
}

func IsSignedNumber(obj interface{}) bool {
	return isSignedNumberKind(reflect.TypeOf(obj).Kind())
}

func IsComplexNumber(obj interface{}) bool {
	return isComplexNumberKind(reflect.TypeOf(obj).Kind())
}

func isNumberKind(kind reflect.Kind) bool {
	return isSignedNumberKind(kind) || isUnsignedNumberKind(kind) || isDecimalNumberKind(kind) || isComplexNumberKind(kind)
}

func isSignedNumberKind(kind reflect.Kind) (b bool) {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		b = true
	}
	return
}

func isUnsignedNumberKind(kind reflect.Kind) (b bool) {
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		b = true
	}
	return
}

func isDecimalNumberKind(kind reflect.Kind) (b bool) {
	switch kind {
	case reflect.Float32, reflect.Float64:
		b = true
	}
	return
}

func isComplexNumberKind(kind reflect.Kind) (b bool) {
	switch kind {
	case reflect.Complex64, reflect.Complex128:
		b = true
	}
	return
}
