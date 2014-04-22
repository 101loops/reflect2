package reflect2

import (
	"reflect"
	"time"
)

func IsExportableField(field reflect.StructField) bool {
	return field.PkgPath == "" // PkgPath is empty for exported fields.
}

func IsStruct(obj interface{}) bool {
	return isStructType(reflect.TypeOf(obj))
}

func IsPointer(obj interface{}) bool {
	return isPointerType(reflect.TypeOf(obj))
}

func IsDefault(obj interface{}) (b bool) {
	switch v := obj.(type) {
	case bool:
		b = (v == false)
	case string:
		b = (v == "")
	case []byte:
		b = len(v) == 0
	case int:
		b = (v == int(0))
	case int8:
		b = (v == int8(0))
	case int16:
		b = (v == int16(0))
	case int32:
		b = (v == int32(0))
	case int64:
		b = (v == int64(0))
	case uint:
		b = (v == uint(0))
	case uint8:
		b = (v == uint8(0))
	case uint16:
		b = (v == uint16(0))
	case uint32:
		b = (v == uint32(0))
	case uint64:
		b = (v == uint64(0))
	case float32:
		b = (v == float32(0))
	case float64:
		b = (v == float64(0))
	case time.Time:
		b = v.IsZero()
	default:
		// TODO: use reflection
	}
	return
}

func isStructType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Struct
}

func isPointerType(typ reflect.Type) bool {
	return typ != nil && typ.Kind() == reflect.Ptr
}
