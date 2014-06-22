package reflector

import . "github.com/101loops/bdd"

var _ = Describe("Utility", func() {

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
