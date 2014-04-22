package reflect2

import (
	"fmt"
	"reflect"
)

func Float2Number(source float64, target interface{}) (err error) {
	switch target := target.(type) {
	case *int:
		*target = int(source)
	case *int8:
		*target = int8(source)
	case *int16:
		*target = int16(source)
	case *int32:
		*target = int32(source)
	case *int64:
		*target = int64(source)
	case *uint:
		*target = uint(source)
	case *uint8:
		*target = uint8(source)
	case *uint16:
		*target = uint16(source)
	case *uint32:
		*target = uint32(source)
	case *uint64:
		*target = uint64(source)
	case *float32:
		*target = float32(source)
	case *float64:
		*target = source
	default:
		err = fmt.Errorf("target (%T) is not a number", target)
	}
	return
}

func Number2Float(source interface{}) (target float64, err error) {
	switch num := source.(type) {
	case int:
		target = float64(num)
	case int8:
		target = float64(num)
	case int16:
		target = float64(num)
	case int32:
		target = float64(num)
	case int64:
		target = float64(num)
	case uint:
		target = float64(num)
	case uint8:
		target = float64(num)
	case uint16:
		target = float64(num)
	case uint32:
		target = float64(num)
	case uint64:
		target = float64(num)
	case float32:
		target = float64(num)
	case float64:
		target = num
	default:
		err = fmt.Errorf("source (%T) is not a number", source)
	}
	return
}

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

func isNumberKind(kind reflect.Kind) bool {
	return isSignedNumberKind(kind) || isUnsignedNumberKind(kind) || isDecimalNumberKind(kind)
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
