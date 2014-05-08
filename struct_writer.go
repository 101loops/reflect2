package reflector

import (
	"fmt"
	"reflect"
)

type StructWriter struct {
	*StructReader
}

func NewStructWriter(dst interface{}) (*StructWriter, error) {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("reflector: writer requires pointer to struct")
	}

	reader, err := NewStructReader(dst)
	if err != nil {
		return nil, err
	}

	return &StructWriter{reader}, nil
}

// SetField sets the struct's field to the provided value.
func (writer *StructWriter) SetFieldValue(name string, value interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("reflector: panic: %v", r)
		}
	}()

	structValue := reflect.ValueOf(writer.obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("reflector: no such field '%s' in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("reflector: cannot set '%s' field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return fmt.Errorf("reflector: value type did not match struct's field type")
	}

	structFieldValue.Set(val)
	return
}

func (writer *StructWriter) SetFieldFloatValue(field string, source float64) (err error) {
	if fieldKind, err := writer.FieldKind(field); err == nil {
		var apply interface{}
		switch fieldKind {
		case reflect.Int:
			apply = int(source)
		case reflect.Int8:
			apply = int8(source)
		case reflect.Int16:
			apply = int16(source)
		case reflect.Int32:
			apply = int32(source)
		case reflect.Int64:
			apply = int64(source)
		case reflect.Uint:
			apply = uint(source)
		case reflect.Uint8:
			apply = uint8(source)
		case reflect.Uint16:
			apply = uint16(source)
		case reflect.Uint32:
			apply = uint32(source)
		case reflect.Uint64:
			apply = uint64(source)
		case reflect.Float32:
			apply = float32(source)
		case reflect.Float64:
			apply = source
		default:
			return fmt.Errorf("reflector: field '%T' is not a number", field)
		}
		err = writer.SetFieldValue(field, apply)
	}
	return
}
