package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

type MainTestSuite struct{}

func init() {
	Suite(&MainTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}

func (s *MainTestSuite) TestAdd(c *C) {
	c.Check(Add(1, 5), Equals, 6)
	c.Check(Add(-1, -2), Equals, -3)
}
