package femebe

import (
	"bytes"
	"io"
	"net"
	"strings"
)

type closableBuffer struct {
	io.ReadWriter
}

func (c *closableBuffer) Close() error {
	// noop, to satisfy interface
	return nil
}

func newClosableBuffer(buf *bytes.Buffer) *closableBuffer {
	return &closableBuffer{buf}
}

// Automatically chooses between unix sockets and tcp sockets for
// listening
func autoListen(place string) (net.Listener, error) {
	if strings.Contains(place, "/") {
		return net.Listen("unix", place)
	}

	return net.Listen("tcp", place)
}

// Automatically chooses between unix sockets and tcp sockets for
// dialing.
func autoDial(place string) (net.Conn, error) {
	if strings.Contains(place, "/") {
		return net.Dial("unix", place)
	}

	return net.Dial("tcp", place)
}
