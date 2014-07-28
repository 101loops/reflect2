package reflector

import (
	"math/cmplx"
	. "github.com/101loops/bdd"
)

var _ = Describe("Numbers", func() {

	Context("conversion", func() {

		var f = float64(42.0)

		It("float -> number", func() {
			var i int
			err := Float2Number(f, &i)
			Check(err, IsNil)
			Check(i, EqualsNum, 42)

			var i8 int8
			err = Float2Number(f, &i8)
			Check(err, IsNil)
			Check(i8, EqualsNum, 42)

			var i16 int16
			err = Float2Number(f, &i16)
			Check(err, IsNil)
			Check(i16, EqualsNum, 42)

			var i32 int32
			err = Float2Number(f, &i32)
			Check(err, IsNil)
			Check(i32, EqualsNum, 42)

			var i64 int64
			err = Float2Number(f, &i64)
			Check(err, IsNil)
			Check(i64, EqualsNum, 42)

			var ui uint
			err = Float2Number(f, &ui)
			Check(err, IsNil)
			Check(ui, EqualsNum, 42)

			var ui8 uint8
			err = Float2Number(f, &ui8)
			Check(err, IsNil)
			Check(ui8, EqualsNum, 42)

			var ui16 uint16
			err = Float2Number(f, &ui16)
			Check(err, IsNil)
			Check(ui16, EqualsNum, 42)

			var ui32 uint32
			err = Float2Number(f, &ui32)
			Check(err, IsNil)
			Check(ui32, EqualsNum, 42)

			var ui64 uint64
			err = Float2Number(f, &ui64)
			Check(err, IsNil)
			Check(ui64, EqualsNum, 42)

			var f32 float32
			err = Float2Number(f, &f32)
			Check(err, IsNil)
			Check(f32, EqualsNum, 42)

			var f64 float64
			err = Float2Number(f, &f64)
			Check(err, IsNil)
			Check(f64, EqualsNum, 42)

			var str string
			err = Float2Number(f, &str)
			Check(err, Contains, "reflector: dst is not a number, but *string")
		})

		It("number -> float", func() {
			fixture := []interface{}{
				int(42), int8(42), int16(42), int32(42), int64(42),
				uint(42), uint8(42), uint16(42), uint32(42), uint64(42),
				float32(42), float64(42),
			}

			for _, num := range fixture {
				f, err := Number2Float(num)
				Check(err, IsNil)
				Check(f, EqualsNum, 42)
			}

			var str string
			_, err := Number2Float(str)
			Check(err, Contains, "reflector: src is not a number, but string")
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
