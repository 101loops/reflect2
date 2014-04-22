package reflect2

import (
	. "github.com/101loops/bdd"
	"reflect"
)

var _ = Describe("Struct Codec", func() {

	It("has field: struct", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
			Yummy: 123,
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		has := refl.HasField("Dummy")
		Check(has, Equals, true)
	})

	It("has field: struct pointer", func() {
		dummyStruct := &TestStruct{
			Dummy: "test",
			Yummy: 123,
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		has := refl.HasField("Dummy")
		Check(has, Equals, true)
	})

	It("has field: non-existing", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
			Yummy: 123,
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		has := refl.HasField("Test")
		Check(has, Equals, false)
	})

	It("has field: non-struct", func() {
		dummy := "abc 123"

		_, err := NewStructReader(dummy)
		Check(err, NotNil)
	})

	It("has field: unexported", func() {
		dummyStruct := TestStruct{
			unexported: 7890,
			Dummy:      "test",
			Yummy:      123,
		}

		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		has := refl.HasField("unexported")
		Check(has, Equals, false)
	})

	It("codec from type", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		t := reflect.TypeOf(dummyStruct)

		refl, err := NewStructReader(t)
		Check(err, IsNil)

		has := refl.HasField("Dummy")
		Check(has, Equals, true)
	})
})
