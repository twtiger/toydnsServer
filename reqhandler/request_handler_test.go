package reqhandler

import (
	"bytes"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type RequestHandlerSuite struct{}

var _ = Suite(&RequestHandlerSuite{})

func (s *RequestHandlerSuite) TestDataFromUDPConnection(c *C) {
	udpPacket := []byte("TEST_CONNECTION")
	m := &mockUDPConn{data: udpPacket}
	output, err := readRequest(m)

	c.Assert(err, IsNil)
	c.Assert(true, Equals, bytes.Equal(udpPacket, output[:len(udpPacket)]))
}

func (s *RequestHandlerSuite) TestDataReadUDPConnectionHasUDPPacketSize(c *C) {
	m := &mockUDPConn{}
	output, err := readRequest(m)

	c.Assert(err, IsNil)
	c.Assert(512, Equals, len(output))
}