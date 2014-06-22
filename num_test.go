package reflector

import (
	. "github.com/101loops/bdd"
	"math/cmplx"
)

var _ = Describe("Numbers", func() {

	Context("conversion", func() {

		var f = float64(42.0)

		It("float -> number", func() {
			var i int
			err := Float2Number(f, &i)
			Check(err, IsNil)
			Check(i, EqualsNum, 42)

			var i32 int32
			err = Float2Number(f, &i32)
			Check(err, IsNil)
			Check(i32, EqualsNum, 42)

			var i64 int64
			err = Float2Number(f, &i64)
			Check(err, IsNil)
			Check(i64, EqualsNum, 42)

			var str string
			err = Float2Number(f, &str)
			Check(err, NotNil)
		})

		It("number -> float", func() {
			f, err := Number2Float(int32(32))
			Check(err, IsNil)
			Check(f, EqualsNum, 32)

			f, err = Number2Float(int64(64))
			Check(err, IsNil)
			Check(f, EqualsNum, 64)

			var str string
			f, err = Number2Float(str)
			Check(err, NotNil)
		})
	})

	It("is number", func() {
		Check(IsNumber(42), Equals, true)
		Check(IsNumber(uint(42)), Equals, true)
		Check(IsNumber(0.4), Equals, true)
		Check(IsNumber("0.4"), Equals, false)
		Check(IsNumber(false), Equals, false)
	})

	It("is decimal number", func() {
		Check(IsDecimalNumber(0.1), Equals, true)
		Check(IsDecimalNumber(0), Equals, false)
	})

	It("is signed number", func() {
		Check(IsSignedNumber(42), Equals, true)
		Check(IsSignedNumber(uint(42)), Equals, false)
	})

	It("is unsigned number", func() {
		Check(IsUnsignedNumber(42), Equals, false)
		Check(IsUnsignedNumber(uint(42)), Equals, true)
	})

	It("is complex number", func() {
		Check(IsComplexNumber(42), Equals, false)
		Check(IsComplexNumber(cmplx.Pow(42, 1.0/3.0)), Equals, true)
	})
})
