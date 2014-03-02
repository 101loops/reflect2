package reflector

import (
	. "launchpad.net/gocheck"
)

func (s *S) TestSetField_on_struct_with_valid_value_type(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructWriter(&dummyStruct)
	c.Assert(err, IsNil)

	err = refl.SetFieldValue("Dummy", "abc")
	c.Assert(err, IsNil)
	c.Assert(dummyStruct.Dummy, Equals, "abc")
}

func (s *S) TestSetField_on_non_pointer(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}

	_, err := NewStructWriter(dummyStruct)
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_on_non_struct(c *C) {
	dummy := "abc 123"

	_, err := NewStructWriter(&dummy)
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_non_existing_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructWriter(&dummyStruct)
	c.Assert(err, IsNil)

	err = refl.SetFieldValue("obladioblada", "life goes on")
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_invalid_value_type(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructWriter(&dummyStruct)
	c.Assert(err, IsNil)

	err = refl.SetFieldValue("Yummy", "123")
	c.Assert(err, NotNil)
}

func (s *S) TestSetField_non_exported_field(c *C) {
	dummyStruct := TestStruct{
		Dummy: "test",
	}
	refl, err := NewStructWriter(&dummyStruct)
	c.Assert(err, IsNil)

	c.Assert(refl.SetFieldValue("unexported", "fail, bitch"), NotNil)
}
