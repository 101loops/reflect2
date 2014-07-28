package reflector

import "reflect"

// IsNumber returns true when the passed-in value has a numeric type.
func IsNumber(obj interface{}) bool {
	return isNumberKind(reflect.TypeOf(obj).Kind())
}

// IsDecimalNumber returns true when the passed-in value is a decimal number.
func IsDecimalNumber(obj interface{}) bool {
	return isDecimalNumberKind(reflect.TypeOf(obj).Kind())
}

// IsUnsignedNumber returns true when the passed-in value is an unsigned number.
func IsUnsignedNumber(obj interface{}) bool {
	return isUnsignedNumberKind(reflect.TypeOf(obj).Kind())
}

// IsSignedNumber returns true when the passed-in value is a signed number.
func IsSignedNumber(obj interface{}) bool {
	return isSignedNumberKind(reflect.TypeOf(obj).Kind())
}

// IsComplexNumber returns true when the passed-in value is a complex number.
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
