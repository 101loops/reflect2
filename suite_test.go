package reflector

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
}

var _ = Suite(&S{})

type TestStruct struct {
	unexported uint64
	Dummy      string `test:"dummytag"`
	Yummy      int    `test:"yummytag"`
}
