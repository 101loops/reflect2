package reflector

import (
	. "github.com/101loops/bdd"
	"time"
)

var (
	defaultTime time.Time
)

type MyString string

var _ = Describe("Utility", func() {
	It("is default", func() {
		Check(IsDefault(true), Equals, false)
		Check(IsDefault(false), Equals, true)

		Check(IsDefault(5), Equals, false)
		Check(IsDefault(0), Equals, true)
		Check(IsDefault(int32(0)), Equals, true)
		Check(IsDefault(int64(0)), Equals, true)

		Check(IsDefault([]byte("ABC")), Equals, false)
		Check(IsDefault([]byte("")), Equals, true)

		Check(IsDefault("test"), Equals, false)
		Check(IsDefault(""), Equals, true)
		//Check(IsDefault(MyString("")), Equals, true) TODO?

		Check(IsDefault(1.0), Equals, false)
		Check(IsDefault(0.0), Equals, true)
		Check(IsDefault(float32(0.0)), Equals, true)
		Check(IsDefault(float64(0.0)), Equals, true)

		Check(IsDefault(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)), Equals, false)
		Check(IsDefault(defaultTime), Equals, true)
	})

	It("is pointer", func() {
		i := 5
		Check(IsPointer(i), Equals, false)
		Check(IsPointer(&i), Equals, true)
	})

	It("is struct", func() {
		i := 5
		type Test struct{}
		Check(IsStruct(i), Equals, false)
		Check(IsStruct(Test{}), Equals, true)
		Check(IsStruct(&Test{}), Equals, false)
	})
})
