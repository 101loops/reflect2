package reflect2

import (
	"fmt"
	"reflect"
	"strings"
)

type StructCodec struct {
	t reflect.Type
}

type FieldCodec struct {
	Index int
	Name  string
	Label string
	Tags  []string
	Type  reflect.Type
}

func NewStructCodec(src interface{}) (*StructCodec, error) {
	v := reflect.ValueOf(src)
	k := v.Kind()
	if t, ok := src.(reflect.Type); ok {
		return &StructCodec{t}, nil
	} else if k == reflect.Struct {
		return &StructCodec{v.Type()}, nil
	} else if k == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		return &StructCodec{v.Elem().Type()}, nil
	}
	return nil, fmt.Errorf("reflect2: invalid entity type %q", k)
}

func (c *StructCodec) Type() reflect.Type {
	return c.t
}

func (c *StructCodec) FieldCodecs(tagNames []string) (res []*FieldCodec, err error) {
	err = c.iterate(func(i int, f reflect.StructField) error {
		code, err := codec(i, f, tagNames)
		if code != nil {
			res = append(res, code)
		}
		return err
	})
	return
}

// HasField checks if the provided field name is part of the struct.
func (c *StructCodec) HasField(name string) bool {
	field, ok := c.t.FieldByName(name)
	if !ok || !IsExportableField(field) {
		return false
	}
	return true
}

func (c *StructCodec) iterate(fn func(int, reflect.StructField) error) error {
	fieldsCount := c.t.NumField()
	for i := 0; i < fieldsCount; i++ {
		field := c.t.Field(i)
		if IsExportableField(field) {
			err := fn(i, field)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func codec(i int, f reflect.StructField, tagNames []string) (res *FieldCodec, err error) {
	t := f.Type
	res = &FieldCodec{Index: i, Name: f.Name, Label: f.Name, Tags: make([]string, 0), Type: t}

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
