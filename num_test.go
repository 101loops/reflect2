// +build unit

package reflector

import (
	. "launchpad.net/gocheck"
)

// TESTS ==========================================================================================

type TestData struct {
	Num   int
	Float float32
	Bool  bool
	Str   string
}

func (s *S) TestFloatNumberConversion(c *C) {
	f := float64(42.0)

	// Float -> Number

	var i int
	err := Float2Number(f, &i)
	c.Assert(err, IsNil)
	c.Assert(i, Equals, int(42))

	var i32 int32
	err = Float2Number(f, &i32)
	c.Assert(err, IsNil)
	c.Assert(i32, Equals, int32(42))

	var i64 int64
	err = Float2Number(f, &i64)
	c.Assert(err, IsNil)
	c.Assert(i64, Equals, int64(42))

	var str string
	err = Float2Number(f, &str)
	c.Assert(err, NotNil)

	// Float -> Field

	obj := &TestData{}
	refl := NewStruct(obj)

	err = refl.SetFieldFloatValue("Num", f)
	c.Assert(err, IsNil)
	c.Assert(obj.Num, Equals, int(42))

	err = refl.SetFieldFloatValue("Float", f)
	c.Assert(err, IsNil)
	c.Assert(obj.Float, Equals, float32(42))

	err = refl.SetFieldFloatValue("Bool", f)
	c.Assert(err, NotNil)

	// Number -> Float

	f, err = Number2Float(int32(32))
	c.Assert(err, IsNil)
	c.Assert(f, Equals, float64(32))

	f, err = Number2Float(int64(64))
	c.Assert(err, IsNil)
	c.Assert(f, Equals, float64(64))

	f, err = Number2Float(str)
	c.Assert(err, NotNil)
}

func (s *S) TestIsNumber(c *C) {
	c.Assert(IsNumber(42), Equals, true)
	c.Assert(IsNumber(0.4), Equals, true)
	c.Assert(IsNumber("0.4"), Equals, false)
	c.Assert(IsNumber(false), Equals, false)
}
