package reflector

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Struct Codec Writer", func() {

	It("set field: struct", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructWriter(&dummyStruct)
		Check(err, IsNil)

		err = refl.SetFieldValue("Dummy", "abc")
		Check(err, IsNil)
		Check(dummyStruct.Dummy, Equals, "abc")
	})

	It("set field: non-pointer", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}

		_, err := NewStructWriter(dummyStruct)
		Check(err, NotNil)
	})

	It("set field: non-struct", func() {
		dummy := "abc 123"

		_, err := NewStructWriter(&dummy)
		Check(err, NotNil)
	})

	It("set field: non-existing", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructWriter(&dummyStruct)
		Check(err, IsNil)

		err = refl.SetFieldValue("obladioblada", "life goes on")
		Check(err, NotNil)
	})

	It("set field: invalid value type", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructWriter(&dummyStruct)
		Check(err, IsNil)

		err = refl.SetFieldValue("Yummy", "123")
		Check(err, NotNil)
	})

	It("set field: non-exported", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructWriter(&dummyStruct)
		Check(err, IsNil)

		Check(refl.SetFieldValue("unexported", "fail, bitch"), NotNil)
	})

})
