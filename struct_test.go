// +build unit

package reflector

import (
	. "launchpad.net/gocheck"
	"reflect"
)

type TestStruct struct {
	unexported uint64
	Dummy      string `test:"dummytag"`
	Yummy      int    `test:"yummytag"`
}

// TESTS ==========================================================================================

func (s *S) TestGetField_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	refl := NewStruct(&dummyStruct)
	value, err := refl.FieldValue("Dummy")
	c.Assert(err, IsNil)
	c.Assert(value, Equals, "test")
}

func (s *S) TestGetField_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.FieldValue("Dummy")
	c.Assert(err, NotNil)
}

func (s *S) TestGetField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl := NewStruct(&dummyStruct)

	_, err := refl.FieldValue("obladioblada")
	c.Assert(err, NotNil)
}

func (s *S) TestGetField_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}
	refl := NewStruct(&dummyStruct)

	c.Assert(func() {
		refl.FieldValue("unexported")
	}, PanicMatches, ".*unexported field.*")
}

func (s *S) TestFieldKind_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	kind, err := refl.FieldKind("Dummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.String)

	kind, err = refl.FieldKind("Yummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.Int)
}

func (s *S) TestFieldKind_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.FieldKind("Dummy")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldKind_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	_, err := refl.FieldKind("obladioblada")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_on_struct(c *C) {
	dummyStruct := TestStruct{}
	refl := NewStruct(&dummyStruct)

	tag, err := refl.FieldTag("Dummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "dummytag")

	tag, err = refl.FieldTag("Yummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "yummytag")
}

func (s *S) TestFieldTag_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.FieldTag("Dummy", "test")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_non_existing_field(c *C) {
	dummyStruct := TestStruct{}
	refl := NewStruct(&dummyStruct)

	_, err := refl.FieldTag("obladioblada", "test")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}
	refl := NewStruct(&dummyStruct)

	_, err := refl.FieldTag("unexported", "test")
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_on_struct_with_valid_value_type(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl := NewStruct(&dummyStruct)

	err := refl.SetFieldValue("Dummy", "abc")
	c.Assert(err, IsNil)
	c.Assert(dummyStruct.Dummy, Equals, "abc")
}

// func (s *S) TestSetField_on_non_struct(c *C) {
//     dummy := "abc 123"

//     err := SetFieldValue(&dummy, "Dummy", "abc")
//     c.Assert(err, NotNil)
// }

func (s *S) TestSetField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl := NewStruct(&dummyStruct)

	err := refl.SetFieldValue("obladioblada", "life goes on")
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_invalid_value_type(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl := NewStruct(&dummyStruct)

	err := refl.SetFieldValue("Yummy", "123")
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_non_exported_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl := NewStruct(&dummyStruct)

	c.Assert(refl.SetFieldValue("unexported", "fail, bitch"), NotNil)
}

func (s *S) TestFields_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	fields, err := refl.FieldNames()
	c.Assert(err, IsNil)
	c.Assert(fields, DeepEquals, []string{"Dummy", "Yummy"})
}

func (s *S) TestFields_on_non_struct(c *C) {
	dummy := "abc 123"

	refl := NewStruct(&dummy)
	_, err := refl.FieldNames()
	c.Assert(err, NotNil)
}

func (s *S) TestFields_with_non_exported_fields(c *C) {
	dummyStruct := TestStruct{
		unexported: 6789,
		Dummy:      "test",
		Yummy:      123,
	}
	refl := NewStruct(&dummyStruct)

	fields, err := refl.FieldNames()
	c.Assert(err, IsNil)
	c.Assert(fields, DeepEquals, []string{"Dummy", "Yummy"})
}

func (s *S) TestHasField_on_struct_with_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	has, err := refl.HasField("Dummy")
	c.Assert(err, IsNil)
	c.Assert(has, Equals, true)
}

func (s *S) TestHasField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	has, err := refl.HasField("Test")
	c.Assert(err, IsNil)
	c.Assert(has, Equals, false)
}

func (s *S) TestHasField_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.HasField("Test")
	c.Assert(err, NotNil)
}

func (s *S) TestHasField_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 7890,
		Dummy:      "test",
		Yummy:      123,
	}
	refl := NewStruct(&dummyStruct)

	has, err := refl.HasField("unexported")
	c.Assert(err, IsNil)
	c.Assert(has, Equals, false)
}

func (s *S) TestTags_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	tags, err := refl.Tags("test")
	c.Assert(err, IsNil)
	c.Assert(tags, DeepEquals, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func (s *S) TestTags_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.Tags("test")
	c.Assert(err, NotNil)
}

func (s *S) TestItems_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl := NewStruct(&dummyStruct)

	tags, err := refl.KeyVal()
	c.Assert(err, IsNil)
	c.Assert(tags, DeepEquals, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func (s *S) TestItems_on_non_struct(c *C) {
	dummy := "abc 123"
	refl := NewStruct(&dummy)

	_, err := refl.KeyVal()
	c.Assert(err, NotNil)
}
