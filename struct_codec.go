package reflector

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

// FACTORY ========================================================================================

func NewStructCodec(obj interface{}) (*StructCodec, error) {
	v := reflect.ValueOf(obj)
	k := v.Kind()
	if t, ok := obj.(reflect.Type); ok {
		return &StructCodec{t}, nil
	} else if k == reflect.Struct {
		return &StructCodec{v.Type()}, nil
	} else if k == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		return &StructCodec{v.Elem().Type()}, nil
	}
	return nil, fmt.Errorf("reflector: invalid entity type '%v'", k)
}

// PUBLIC METHODS =================================================================================

func (self *StructCodec) Type() reflect.Type {
	return self.t
}

func (self *StructCodec) FieldCodecs(tagNames []string) (res []*FieldCodec, err error) {
	res = make([]*FieldCodec, 0)
	err = self.iterate(func(i int, f reflect.StructField) error {
		code, err := codec(i, f, tagNames)
		if code != nil {
			res = append(res, code)
		}
		return err
	})
	return
}

// HasField checks if the provided field name is part of the struct.
func (self *StructCodec) HasField(name string) bool {
	field, ok := self.t.FieldByName(name)
	if !ok || !IsExportableField(field) {
		return false
	}
	return true
}

// PRIVATE METHODS ================================================================================

func (self *StructCodec) iterate(fn func(int, reflect.StructField) error) error {
	fieldsCount := self.t.NumField()
	for i := 0; i < fieldsCount; i++ {
		field := self.t.Field(i)
		if IsExportableField(field) {
			err := fn(i, field)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// HELPERS ========================================================================================

func codec(i int, f reflect.StructField, tagNames []string) (res *FieldCodec, err error) {
	t := f.Type
	res = &FieldCodec{i, f.Name, f.Name, make([]string, 0), t}

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
