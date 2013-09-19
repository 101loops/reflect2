package reflector

import (
	"fmt"
	"reflect"
)

type StructReader struct {
	*StructCodec
	obj interface{}
	val reflect.Value
}

type Field struct {
	*FieldCodec
	Data interface{}
}

// FACTORY ========================================================================================

func NewStructReader(obj interface{}) (*StructReader, error) {
	codec, err := NewStructCodec(obj)
	if err != nil {
		return nil, err
	}

	// dereference pointer
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return &StructReader{codec, obj, v}, nil
}

// PUBLIC METHODS =================================================================================

// Fields iterates over the struct's fields and calls the provided function with each field.
func (self *StructReader) Iterate(tagNames []string, fn func(*Field) error) error {
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
func (self *StructReader) Fields(tagNames []string) (res []*Field, err error) {
	res = make([]*Field, 0)
	err = self.iterate(func(i int, f reflect.StructField) error {
		fld, err := field(i, self.val.Field(i), f, tagNames)
		if fld != nil {
			res = append(res, fld)
		}
		return err
	})
	return
}

// FieldNames returns the struct fields name list.
func (self *StructReader) FieldNames() ([]string, error) {
	var fields []string
	err := self.iterate(func(_ int, f reflect.StructField) error {
		fields = append(fields, f.Name)
		return nil
	})
	return fields, err
}

// Tags lists the struct tag fields.
func (self *StructReader) Tags(key string) (map[string]string, error) {
	tags := make(map[string]string)
	err := self.iterate(func(_ int, f reflect.StructField) error {
		tags[f.Name] = f.Tag.Get(key)
		return nil
	})
	return tags, err
}

// FieldValue returns the value of the struct field.
func (self *StructReader) FieldValue(name string) (interface{}, error) {
	field := self.val.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("reflector: no such field '%s' in obj", name)
	}
	return field.Interface(), nil
}

// FieldType returns the kind of the struct field.
func (self *StructReader) FieldType(name string) (reflect.Type, error) {
	field := self.val.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("reflector: no such field '%s' in obj", name)
	}
	return field.Type(), nil
}

// FieldKind returns the kind of the struct field.
func (self *StructReader) FieldKind(name string) (reflect.Kind, error) {
	typ, err := self.FieldType(name)
	if err != nil {
		return reflect.Invalid, err
	}
	return typ.Kind(), nil
}

// FieldTag returns the struct's field tag value.
func (self *StructReader) FieldTag(name, tagKey string) (string, error) {
	field, ok := self.Type().FieldByName(name)
	if !ok {
		return "", fmt.Errorf("reflector: no such field '%s' in obj", name)
	}

	if !IsExportableField(field) {
		return "", fmt.Errorf("reflector: cannot access non-exported field")
	}

	return field.Tag.Get(tagKey), nil
}

// KeyVal returns the field/value struct pairs as a map.
func (self *StructReader) KeyVal() (map[string]interface{}, error) {
	fieldsCount := self.Type().NumField()

	items := make(map[string]interface{})

	for i := 0; i < fieldsCount; i++ {
		field := self.Type().Field(i)
		fieldValue := self.val.Field(i)

		// Make sure only exportable and addressable fields are returned
		if IsExportableField(field) {
			items[field.Name] = fieldValue.Interface()
		}
	}

	return items, nil
}

// HELPERS ========================================================================================

func field(i int, v reflect.Value, f reflect.StructField, tagNames []string) (*Field, error) {
	codec, err := codec(i, f, tagNames)
	if codec == nil || err != nil {
		return nil, err
	}
	return &Field{codec, v.Interface()}, nil
}
