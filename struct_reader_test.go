package reflector

import (
	. "github.com/101loops/bdd"
	"reflect"
)

var _ = Describe("Struct Codec Reader", func() {

	It("create from non-struct", func() {
		dummy := "abc 123"

		_, err := NewStructReader(dummy)
		Check(err, NotNil)
	})

	It("get field: struct", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}

		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		value, err := refl.FieldValue("Dummy")
		Check(err, IsNil)
		Check(value, Equals, "test")
	})

	It("get field: struct pointer", func() {
		dummyStruct := &TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		value, err := refl.FieldValue("Dummy")
		Check(err, IsNil)
		Check(value, Equals, "test")
	})

	It("get field: non-existing", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		_, err = refl.FieldValue("obladioblada")
		Check(err, NotNil)
	})

	It("get field: unexported", func() {
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

	It("field kind: struct", func() {
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

	It("field kind: struct pointer", func() {
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

	It("field kind: non-existing", func() {
		dummyStruct := TestStruct{
			Dummy: "test",
			Yummy: 123,
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		_, err = refl.FieldKind("obladioblada")
		Check(err, NotNil)
	})

	It("field tag: struct", func() {
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

	It("field tag: struct pointer", func() {
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

	It("field tag: non-existing", func() {
		dummyStruct := TestStruct{}

		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		_, err = refl.FieldTag("obladioblada", "test")
		Check(err, NotNil)
	})

	It("field tag: unexported", func() {
		dummyStruct := TestStruct{
			unexported: 12345,
			Dummy:      "test",
		}
		refl, err := NewStructReader(dummyStruct)
		Check(err, IsNil)

		_, err = refl.FieldTag("unexported", "test")
		Check(err, NotNil)
	})

	It("field names: struct", func() {
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

	It("field names: struct pointer", func() {
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

	It("field names: non-exported", func() {
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

	It("tags: struct", func() {
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

	It("tags: non-exported", func() {
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
