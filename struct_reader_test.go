package reflector

import (
	. "github.com/101loops/bdd"
	"reflect"
)

var _ = Describe("Struct Codec Reader", func() {

	Context("field value", func() {

		It("struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}

			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			value, err := refl.FieldValue("Dummy")
			Check(err, IsNil)
			Check(value, Equals, "test")
		})

		It("struct pointer", func() {
			dummyStruct := &TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			value, err := refl.FieldValue("Dummy")
			Check(err, IsNil)
			Check(value, Equals, "test")
		})

		It("non-existing", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			_, err = refl.FieldValue("obladioblada")
			Check(err, NotNil)
		})

		It("unexported", func() {
			dummyStruct := TestStruct{
				unexported: 12345,
				Dummy:      "test",
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			Check(func() {
				refl.FieldValue("unexported")
			}, Panics)
		})
	})

	Context("field kind", func() {

		It("struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			kind, err := refl.FieldKind("Dummy")
			Check(err, IsNil)
			Check(kind, Equals, reflect.String)

			kind, err = refl.FieldKind("Yummy")
			Check(err, IsNil)
			Check(kind, Equals, reflect.Int)
		})

		It("struct pointer", func() {
			dummyStruct := &TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			kind, err := refl.FieldKind("Dummy")
			Check(err, IsNil)
			Check(kind, Equals, reflect.String)

			kind, err = refl.FieldKind("Yummy")
			Check(err, IsNil)
			Check(kind, Equals, reflect.Int)
		})

		It("non-existing", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			_, err = refl.FieldKind("obladioblada")
			Check(err, NotNil)
		})
	})

	Context("field tag", func() {

		It("struct", func() {
			dummyStruct := TestStruct{}

			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			tag, err := refl.FieldTag("Dummy", "test")
			Check(err, IsNil)
			Check(tag, Equals, "dummytag")

			tag, err = refl.FieldTag("Yummy", "test")
			Check(err, IsNil)
			Check(tag, Equals, "yummytag")
		})

		It("struct pointer", func() {
			dummyStruct := &TestStruct{}

			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			tag, err := refl.FieldTag("Dummy", "test")
			Check(err, IsNil)
			Check(tag, Equals, "dummytag")

			tag, err = refl.FieldTag("Yummy", "test")
			Check(err, IsNil)
			Check(tag, Equals, "yummytag")
		})

		It("non-existing", func() {
			dummyStruct := TestStruct{}

			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			_, err = refl.FieldTag("obladioblada", "test")
			Check(err, NotNil)
		})

		It("unexported", func() {
			dummyStruct := TestStruct{
				unexported: 12345,
				Dummy:      "test",
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			_, err = refl.FieldTag("unexported", "test")
			Check(err, NotNil)
		})
	})

	Context("field names", func() {

		It("struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			fields, err := refl.FieldNames()
			Check(err, IsNil)
			Check(fields, Equals, []string{"Dummy", "Yummy"})
		})

		It("struct pointer", func() {
			dummyStruct := &TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			fields, err := refl.FieldNames()
			Check(err, IsNil)
			Check(fields, Equals, []string{"Dummy", "Yummy"})
		})

		It("non-exported", func() {
			dummyStruct := TestStruct{
				unexported: 6789,
				Dummy:      "test",
				Yummy:      123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			fields, err := refl.FieldNames()
			Check(err, IsNil)
			Check(fields, Equals, []string{"Dummy", "Yummy"})
		})
	})

	Context("tags", func() {

		It("struct", func() {
			dummyStruct := TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			tags, err := refl.Tags("test")
			Check(err, IsNil)
			Check(tags, Equals, map[string]string{
				"Dummy": "dummytag",
				"Yummy": "yummytag",
			})
		})

		It("non-exported", func() {
			dummyStruct := &TestStruct{
				Dummy: "test",
				Yummy: 123,
			}
			refl, err := NewStructReader(dummyStruct)
			Check(err, IsNil)

			tags, err := refl.Tags("test")
			Check(err, IsNil)
			Check(tags, Equals, map[string]string{
				"Dummy": "dummytag",
				"Yummy": "yummytag",
			})
		})
	})

	It("key val: struct", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
			Yummy: 123,
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		tags, err := refl.KeyVal()
		Check(err, IsNil)
		Check(tags, Equals, map[string]interface{}{
			"Dummy": "test",
			"Yummy": 123,
		})
	})
})
