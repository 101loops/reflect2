package reflector

import "reflect"

func IsExportableField(field reflect.StructField) bool {
	return field.PkgPath == "" // PkgPath is empty for exported fields.
}

func IsStruct(obj interface{}) bool {
	return isStructType(reflect.TypeOf(obj))
}

func isStructType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Struct
}

func IsPointer(obj interface{}) bool {
	return isPointerType(reflect.TypeOf(obj))
}

func isPointerType(typ reflect.Type) bool {
	return typ != nil && typ.Kind() == reflect.Ptr
}
