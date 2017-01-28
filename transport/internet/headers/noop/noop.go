package noop

import (
	"context"
	"net"

	"v2ray.com/core/common"
)

type NoOpHeader struct{}

func (v NoOpHeader) Size() int {
	return 0
}
func (v NoOpHeader) Write([]byte) (int, error) {
	return 0, nil
}

func NewNoOpHeader(context.Context, interface{}) (interface{}, error) {
	return NoOpHeader{}, nil
}

type NoOpConnectionHeader struct{}

func (NoOpConnectionHeader) Client(conn net.Conn) net.Conn {
	return conn
}

func (NoOpConnectionHeader) Server(conn net.Conn) net.Conn {
	return conn
}

func NewNoOpConnectionHeader(context.Context, interface{}) (interface{}, error) {
	return NoOpConnectionHeader{}, nil
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), NewNoOpHeader))
	common.Must(common.RegisterConfig((*ConnectionConfig)(nil), NewNoOpConnectionHeader))
}
