package reflector

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Struct struct {
	obj interface{}
}

type Field struct {
	Index int
	Name  string
	Label string
	Tags  []string
	Value interface{}
	Type  reflect.Type
}

// FACTORY ========================================================================================

func NewStruct(obj interface{}) *Struct {
	return &Struct{obj}
}

// PUBLIC METHODS =================================================================================

// Fields iterates over the struct's fields and calls the provided function with each field.
func (self *Struct) Iterate(tagNames []string, fn func(*Field) error) error {
	fields, err := self.Fields(tagNames)
	if err == nil {
		for _, f := range fields {
			err = fn(f)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// Fields returns the struct's fields.
func (self *Struct) Fields(tagNames []string) (res []*Field, err error) {
	res = make([]*Field, 0)
	err = self.iterate(func(i int, v reflect.Value, f reflect.StructField) error {
		meta, err := self.meta(v, f, tagNames)
		if meta != nil {
			meta.Index = i
			res = append(res, meta)
		}
		return err
	})
	return
}

// FieldNames returns the struct fields name list.
func (self *Struct) FieldNames() ([]string, error) {
	var fields []string
	err := self.iterate(func(_ int, _ reflect.Value, f reflect.StructField) error {
		fields = append(fields, f.Name)
		return nil
	})
	return fields, err
}

// Tags lists the struct tag fields.
func (self *Struct) Tags(key string) (map[string]string, error) {
	tags := make(map[string]string)
	err := self.iterate(func(_ int, _ reflect.Value, f reflect.StructField) error {
		tags[f.Name] = f.Tag.Get(key)
		return nil
	})
	return tags, err
}

// FieldValue returns the value of the struct field.
func (self *Struct) FieldValue(name string) (interface{}, error) {
	objValue, _, err := self.ValType()
	if err != nil {
		return nil, err
	}

	field := objValue.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("no such field '%s' in obj", name)
	}

	return field.Interface(), nil
}

// FieldType returns the kind of the struct field.
func (self *Struct) FieldType(name string) (reflect.Type, error) {
	objValue, _, err := self.ValType()
	if err != nil {
		return nil, err
	}

	field := objValue.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("no such field '%s' in obj", name)
	}

	return field.Type(), nil
}

// FieldKind returns the kind of the struct field.
func (self *Struct) FieldKind(name string) (reflect.Kind, error) {
	typ, err := self.FieldType(name)
	if err != nil {
		return reflect.Invalid, err
	}

	return typ.Kind(), nil
}

// FieldTag returns the struct's field tag value.
func (self *Struct) FieldTag(name, tagKey string) (string, error) {
	_, objType, err := self.ValType()
	if err != nil {
		return "", err
	}

	field, ok := objType.FieldByName(name)
	if !ok {
		return "", fmt.Errorf("no such field '%s' in obj", name)
	}

	if !IsExportableField(field) {
		return "", errors.New("cannot FieldTag on a non-exported struct field")
	}

	return field.Tag.Get(tagKey), nil
}

// HasField checks if the provided field name is part of the struct.
func (self *Struct) HasField(name string) (bool, error) {
	_, objType, err := self.ValType()
	if err != nil {
		return false, err
	}

	field, ok := objType.FieldByName(name)
	if !ok || !IsExportableField(field) {
		return false, nil
	}

	return true, nil
}

// Items returns the field/value struct pairs as a map.
func (self *Struct) KeyVal() (map[string]interface{}, error) {
	objValue, objType, err := self.ValType()
	if err != nil {
		return nil, err
	}

	fieldsCount := objType.NumField()

	items := make(map[string]interface{})

	for i := 0; i < fieldsCount; i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		// Make sure only exportable and addressable fields are
		// returned by Items
		if IsExportableField(field) {
			items[field.Name] = fieldValue.Interface()
		}
	}

	return items, nil
}

// SetField sets the struct's field to the provided value.
func (self *Struct) SetFieldValue(name string, value interface{}) error {
	// Fetch the field reflect.Value
	structValue := reflect.ValueOf(self.obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field '%s' in obj", name)
	}

	// If obj field value is not settable an error is thrown
	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func (self *Struct) SetFieldFloatValue(field string, source float64) (err error) {
	if fieldKind, err := self.FieldKind(field); err == nil {
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
			return fmt.Errorf("field '%T' is not a number", field)
		}
		err = self.SetFieldValue(field, apply)
	}
	return
}

func (self *Struct) ValType() (val reflect.Value, typ reflect.Type, err error) {
	objType := reflect.TypeOf(self.obj)
	if !isPointerType(objType) {
		err = errors.New("cannot reflect on non-pointer struct")
	} else {
		val = reflect.ValueOf(self.obj).Elem()
		typ = val.Type()

		if !isStructType(typ) {
			err = errors.New("cannot reflect on non-struct")
		}
	}
	return
}

// PRIVATE METHODS ================================================================================

func (self *Struct) iterate(fn func(int, reflect.Value, reflect.StructField) error) error {
	objVal, objType, err := self.ValType()
	if err != nil {
		return err
	}

	fieldsCount := objType.NumField()
	for i := 0; i < fieldsCount; i++ {
		field := objType.Field(i)
		if IsExportableField(field) {
			err = fn(i, objVal.Field(i), field)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (self *Struct) meta(v reflect.Value, f reflect.StructField, tagNames []string) (res *Field, err error) {

	ft := f.Type
	res = &Field{0, f.Name, f.Name, make([]string, 0), v.Interface(), ft} // TODO

	// iterate over tag names (reverse: 1st item overwrites 2nd etc.)
	for i := len(tagNames) - 1; i >= 0; i-- {

		// field tags
		tagName := tagNames[i]
		fieldTag := f.Tag.Get(tagName)
		tags := strings.Split(fieldTag, ",")
		if len(tags) > 1 {
			res.Tags = tags[1:]
		}

		// field label
		if len(tags) > 0 && tags[0] != "" {
			res.Label = tags[0]
		}
		res.Label = strings.ToLower(res.Label)
	}

	if res.Label == "-" {
		res = nil
	}

	return
}
