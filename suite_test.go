package reflector

import (
	. "github.com/101loops/bdd"
	"testing"
)

func TestSuite(t *testing.T) {
	RunSpecs(t, "Reflector Suite")
}

type TestStruct struct {
	unexported uint64
	Dummy      string `test:"dummytag"`
	Yummy      int    `test:"yummytag"`
}
