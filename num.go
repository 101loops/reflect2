package reflect2

import (
	"fmt"
	"reflect"
)

func Float2Number(src float64, dst interface{}) (err error) {
	switch target := target.(type) {
	case *int:
		*target = int(src)
	case *int8:
		*target = int8(src)
	case *int16:
		*target = int16(src)
	case *int32:
		*target = int32(src)
	case *int64:
		*target = int64(src)
	case *uint:
		*target = uint(src)
	case *uint8:
		*target = uint8(src)
	case *uint16:
		*target = uint16(src)
	case *uint32:
		*target = uint32(src)
	case *uint64:
		*target = uint64(src)
	case *float32:
		*target = float32(src)
	case *float64:
		*target = src
	default:
		err = fmt.Errorf("target (%T) is not a number", target)
	}
	return
}

func Number2Float(src interface{}) (target float64, err error) {
	switch num := src.(type) {
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
		err = fmt.Errorf("source (%T) is not a number", src)
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
