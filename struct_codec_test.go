package reflector

import (
	. "launchpad.net/gocheck"
	"reflect"
)

func (s *S) TestHasField_on_struct_with_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	has := refl.HasField("Dummy")
	c.Assert(has, Equals, true)
}

func (s *S) TestHasField_on_struct_pointer_with_existing_field(c *C) {
	dummyStruct := &TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	has := refl.HasField("Dummy")
	c.Assert(has, Equals, true)
}

func (s *S) TestHasField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
		Yummy: 123,
	}
	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	has := refl.HasField("Test")
	c.Assert(has, Equals, false)
}

func (s *S) TestHasField_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructReader(dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestHasField_unexported_field(c *C) {
	dummyStruct := TestStruct{
		unexported: 7890,
		Dummy:      "test",
		Yummy:      123,
	}

	refl, err := NewStructReader(dummyStruct)
	c.Assert(err, IsNil)

	has := refl.HasField("unexported")
	c.Assert(has, Equals, false)
}

func (s *S) TestNewStructCodec_from_type(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	t := reflect.TypeOf(dummyStruct)

	refl, err := NewStructReader(t)
	c.Assert(err, IsNil)

	has := refl.HasField("Dummy")
	c.Assert(has, Equals, true)
}
