package reflector

import (
	. "launchpad.net/gocheck"
	"reflect"
)

func (s *S) TestGetField_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	value, err := refl.FieldValue("Dummy")
	c.Assert(err, IsNil)
	c.Assert(value, Equals, "test")
}

func (s *S) TestGetField_on_struct_pointer(c *C) {
	dummyStruct := &TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	value, err := refl.FieldValue("Dummy")
	c.Assert(err, IsNil)
	c.Assert(value, Equals, "test")
}

func (s *S) TestGetField_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestGetField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	_, err = refl.FieldValue("obladioblada")
	c.Assert(err, NotNil)
}

func (s *S) TestGetField_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	c.Assert(func() {
		refl.FieldValue("unexported")
	}, PanicMatches, ".*unexported field.*")
}

func (s *S) TestFieldKind_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	kind, err := refl.FieldKind("Dummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.String)

	kind, err = refl.FieldKind("Yummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.Int)
}

func (s *S) TestFieldKind_on_struct_pointer(c *C) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	kind, err := refl.FieldKind("Dummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.String)

	kind, err = refl.FieldKind("Yummy")
	c.Assert(err, IsNil)
	c.Assert(kind, Equals, reflect.Int)
}

func (s *S) TestFieldKind_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestFieldKind_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	_, err = refl.FieldKind("obladioblada")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_on_struct(c *C) {
	dummyStruct := TestStruct{}

	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	tag, err := refl.FieldTag("Dummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "dummytag")

	tag, err = refl.FieldTag("Yummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "yummytag")
}

func (s *S) TestFieldTag_on_struct_pointer(c *C) {
	dummyStruct := &TestStruct{}

	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	tag, err := refl.FieldTag("Dummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "dummytag")

	tag, err = refl.FieldTag("Yummy", "test")
	c.Assert(err, IsNil)
	c.Assert(tag, Equals, "yummytag")
}

func (s *S) TestFieldTag_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_non_existing_field(c *C) {
	dummyStruct := TestStruct{}

	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	_, err = refl.FieldTag("obladioblada", "test")
	c.Assert(err, NotNil)
}

func (s *S) TestFieldTag_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 12345,
		Dummy:      "test",
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	_, err = refl.FieldTag("unexported", "test")
	c.Assert(err, NotNil)
}

func (s *S) TestFields_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	fields, err := refl.FieldNames()
	c.Assert(err, IsNil)
	c.Assert(fields, DeepEquals, []string{"Dummy", "Yummy"})
}

func (s *S) TestFields_on_struct_pointer(c *C) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	fields, err := refl.FieldNames()
	c.Assert(err, IsNil)
	c.Assert(fields, DeepEquals, []string{"Dummy", "Yummy"})
}

func (s *S) TestFields_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestFields_with_non_exported_fields(c *C) {
	dummyStruct := TestStruct{
		unexported: 6789,
		Dummy:      "test",
		Yummy:      123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	fields, err := refl.FieldNames()
	c.Assert(err, IsNil)
	c.Assert(fields, DeepEquals, []string{"Dummy", "Yummy"})
}

func (s *S) TestTags_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	tags, err := refl.Tags("test")
	c.Assert(err, IsNil)
	c.Assert(tags, DeepEquals, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func (s *S) TestTags_on_struct_pointer(c *C) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	tags, err := refl.Tags("test")
	c.Assert(err, IsNil)
	c.Assert(tags, DeepEquals, map[string]string{
		"Dummy": "dummytag",
		"Yummy": "yummytag",
	})
}

func (s *S) TestTags_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestItems_on_struct(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	tags, err := refl.KeyVal()
	c.Assert(err, IsNil)
	c.Assert(tags, DeepEquals, map[string]interface{}{
		"Dummy": "test",
		"Yummy": 123,
	})
}

func (s *S) TestItems_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}
