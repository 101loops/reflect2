package reflector

import "reflect"

// IsStruct returns true when the passed-in value is a struct type.
func IsStruct(obj interface{}) bool {
	return isStructType(reflect.TypeOf(obj))
}

func isStructType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Struct
}

// IsPointer returns true when the passed-in value is a pointer.
func IsPointer(obj interface{}) bool {
	return isPointerType(reflect.TypeOf(obj))
}

func isPointerType(typ reflect.Type) bool {
	return typ != nil && typ.Kind() == reflect.Ptr
}
