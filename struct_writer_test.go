package reflector

import (
	. "github.com/101loops/bdd"
)

var _ = Describe("Struct Codec Writer", func() {

	Context("create", func() {

		It("from non-pointer", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}

			_, err := NewStructWriter(dummyStruct)
			Check(err, Contains, "reflector: writer requires pointer to struct")
		})

		It("from non-struct", func() {
			dummy := "abc 123"

			_, err := NewStructWriter(&dummy)
			Check(err, Contains, `reflector: is not a struct, reflect.Type or pointer to struct, but "ptr"`)
		})
	})

	Context("set field value", func() {

		It("struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructWriter(&dummyStruct)
			Check(err, IsNil)

			err = refl.SetFieldValue("Dummy", "abc")
			Check(err, IsNil)
			Check(dummyStruct.Dummy, Equals, "abc")
		})

		It("non-existing", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructWriter(&dummyStruct)
			Check(err, IsNil)

			err = refl.SetFieldValue("obladioblada", "life goes on")
			Check(err, NotNil)
		})

		It("invalid value type", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructWriter(&dummyStruct)
			Check(err, IsNil)

			err = refl.SetFieldValue("Yummy", "123")
			Check(err, NotNil)
		})

		It("non-exported", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructWriter(&dummyStruct)
			Check(err, IsNil)

			Check(refl.SetFieldValue("unexported", "fail, bitch"), NotNil)
		})
	})

	It("set field to float", func() {
		type TestData struct {
			Num   int
			Float float32
			Bool  bool
			Str   string
		}

		var f = float64(42.0)
		obj := &TestData{}

		refl, err := NewStructWriter(obj)
		Check(err, IsNil)

		err = refl.SetFieldFloatValue("Num", f)
		Check(err, IsNil)
		Check(obj.Num, EqualsNum, 42)

		err = refl.SetFieldFloatValue("Float", f)
		Check(err, IsNil)
		Check(obj.Float, EqualsNum, 42)

		err = refl.SetFieldFloatValue("Bool", f)
		Check(err, NotNil)
	})

})
