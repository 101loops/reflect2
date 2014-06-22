package reflector

import (
	. "github.com/101loops/bdd"
	"reflect"
)

var _ = Describe("Struct Codec", func() {

	Context("create", func() {

		It("from struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructCodec(dummyStruct)
			Check(err, IsNil)

			has := refl.HasField("Dummy")
			Check(has, Equals, true)
		})

		It("from reflect type", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			t := reflect.TypeOf(dummyStruct)

			refl, err := NewStructCodec(t)
			Check(err, IsNil)

			has := refl.HasField("Dummy")
			Check(has, Equals, true)
		})

		It("non-struct", func() {
			dummy := "abc 123"

			_, err := NewStructCodec(dummy)
			Check(err, Contains, `reflector: is not a struct, reflect.Type or pointer to struct, but "string"`)
		})
	})

	Context("has field", func() {

		It("struct pointer", func() {
			dummyStruct := &TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructCodec(dummyStruct)
			Check(err, IsNil)

			has := refl.HasField("Dummy")
			Check(has, Equals, true)
		})

		It("non-existing", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructCodec(dummyStruct)
			Check(err, IsNil)

			has := refl.HasField("Test")
			Check(has, Equals, false)
		})

		It("unexported", func() {
			dummyStruct := TestStruct{
				unexported: 7890,
				Dummy:      "test",
				Yummy:      123,
			}

			refl, err := NewStructCodec(dummyStruct)
			Check(err, IsNil)

			has := refl.HasField("unexported")
			Check(has, Equals, false)
		})
	})

})
