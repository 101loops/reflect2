// +build unit

package reflector

import (
	. "launchpad.net/gocheck"
	"time"
)

type MyString string

// TESTS ==========================================================================================

func (s *S) TestDefault(c *C) {

	c.Assert(IsDefault(true), Equals, false)
	c.Assert(IsDefault(false), Equals, true)

	c.Assert(IsDefault(5), Equals, false)
	c.Assert(IsDefault(0), Equals, true)
	c.Assert(IsDefault(int32(0)), Equals, true)
	c.Assert(IsDefault(int64(0)), Equals, true)

	c.Assert(IsDefault([]byte("ABC")), Equals, false)
	c.Assert(IsDefault([]byte("")), Equals, true)

	c.Assert(IsDefault("test"), Equals, false)
	c.Assert(IsDefault(""), Equals, true)
	//c.Assert(IsDefault(MyString("")), Equals, true) TODO?

	c.Assert(IsDefault(1.0), Equals, false)
	c.Assert(IsDefault(0.0), Equals, true)
	c.Assert(IsDefault(float32(0.0)), Equals, true)
	c.Assert(IsDefault(float64(0.0)), Equals, true)

	c.Assert(IsDefault(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)), Equals, false)
	c.Assert(IsDefault(defaultTime), Equals, true)
}
