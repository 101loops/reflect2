package reflect2

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

// Fields iterates over the struct's fields and calls the provided function with each field.
func (reader *StructReader) Iterate(tagNames []string, fn func(*Field) error) error {
	fields, err := reader.Fields(tagNames)
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
func (reader *StructReader) Fields(tagNames []string) (res []*Field, err error) {
	err = reader.iterate(func(i int, f reflect.StructField) error {
		fld, err := field(i, reader.val.Field(i), f, tagNames)
		if fld != nil {
			res = append(res, fld)
		}
		return err
	})
	return
}

// FieldNames returns the struct fields name list.
func (reader *StructReader) FieldNames() ([]string, error) {
	var fields []string
	err := reader.iterate(func(_ int, f reflect.StructField) error {
		fields = append(fields, f.Name)
		return nil
	})
	return fields, err
}

// Tags lists the struct tag fields.
func (reader *StructReader) Tags(key string) (map[string]string, error) {
	tags := make(map[string]string)
	err := reader.iterate(func(_ int, f reflect.StructField) error {
		tags[f.Name] = f.Tag.Get(key)
		return nil
	})
	return tags, err
}

// FieldValue returns the value of the struct field.
func (reader *StructReader) FieldValue(name string) (interface{}, error) {
	field := reader.val.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("reflector: no such field '%s' in obj", name)
	}
	return field.Interface(), nil
}

// FieldType returns the kind of the struct field.
func (reader *StructReader) FieldType(name string) (reflect.Type, error) {
	field := reader.val.FieldByName(name)
	if !field.IsValid() {
		return nil, fmt.Errorf("reflector: no such field '%s' in obj", name)
	}
	return field.Type(), nil
}

// FieldKind returns the kind of the struct field.
func (reader *StructReader) FieldKind(name string) (reflect.Kind, error) {
	typ, err := reader.FieldType(name)
	if err != nil {
		return reflect.Invalid, err
	}
	return typ.Kind(), nil
}

// FieldTag returns the struct's field tag value.
func (reader *StructReader) FieldTag(name, tagKey string) (string, error) {
	field, ok := reader.Type().FieldByName(name)
	if !ok {
		return "", fmt.Errorf("reflector: no such field '%s' in obj", name)
	}

	if !IsExportableField(field) {
		return "", fmt.Errorf("reflector: cannot access non-exported field")
	}

	return field.Tag.Get(tagKey), nil
}

// KeyVal returns the field/value struct pairs as a map.
func (reader *StructReader) KeyVal() (map[string]interface{}, error) {
	fieldsCount := reader.Type().NumField()

	items := make(map[string]interface{})

	for i := 0; i < fieldsCount; i++ {
		field := reader.Type().Field(i)
		fieldValue := reader.val.Field(i)

		// Make sure only exportable and addressable fields are returned
		if IsExportableField(field) {
			items[field.Name] = fieldValue.Interface()
		}
	}

	return items, nil
}

func field(i int, v reflect.Value, f reflect.StructField, tagNames []string) (*Field, error) {
	codec, err := codec(i, f, tagNames)
	if codec == nil || err != nil {
		return nil, err
	}
	return &Field{codec, v.Interface()}, nil
}
