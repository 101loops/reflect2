package reflect2

import (
	. "github.com/101loops/bdd"
)

type TestData struct {
	Num   int
	Float float32
	Bool  bool
	Str   string
}

var _ = Describe("Numbers", func() {

	It("conversion", func() {
		f := float64(42.0)

		// Float -> Number

		var i int
		err := Float2Number(f, &i)
		Check(err, IsNil)
		Check(i, Equals, int(42))

		var i32 int32
		err = Float2Number(f, &i32)
		Check(err, IsNil)
		Check(i32, Equals, int32(42))

		var i64 int64
		err = Float2Number(f, &i64)
		Check(err, IsNil)
		Check(i64, Equals, int64(42))

		var str string
		err = Float2Number(f, &str)
		Check(err, NotNil)

		// Float -> Field

		obj := &TestData{}
		refl, err := NewStructWriter(obj)
		Check(err, IsNil)

		err = refl.SetFieldFloatValue("Num", f)
		Check(err, IsNil)
		Check(obj.Num, Equals, int(42))

		err = refl.SetFieldFloatValue("Float", f)
		Check(err, IsNil)
		Check(obj.Float, Equals, float32(42))

		err = refl.SetFieldFloatValue("Bool", f)
		Check(err, NotNil)

		// Number -> Float

		f, err = Number2Float(int32(32))
		Check(err, IsNil)
		Check(f, Equals, float64(32))

		f, err = Number2Float(int64(64))
		Check(err, IsNil)
		Check(f, Equals, float64(64))

		f, err = Number2Float(str)
		Check(err, NotNil)
	})

	It("check", func() {
		Check(IsNumber(42), Equals, true)
		Check(IsNumber(0.4), Equals, true)
		Check(IsNumber("0.4"), Equals, false)
		Check(IsNumber(false), Equals, false)

		Check(IsDecimalNumber(0.1), Equals, true)
		Check(IsDecimalNumber(0), Equals, false)

		Check(IsSignedNumber(42), Equals, true)
		Check(IsSignedNumber(uint(42)), Equals, false)

		Check(IsUnsignedNumber(42), Equals, false)
		Check(IsUnsignedNumber(uint(42)), Equals, true)
	})
})
