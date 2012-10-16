package femebe

import (
	"bytes"
	"net"
	"testing"
)

//Benchmark for testing how long the send operation on a MessageStream takes when used soly in
//memory.
func BenchmarkEchoSendMem(b *testing.B) {
	b.StopTimer()
	var ping Message
	var pong Message

	ping.InitFromBytes('i', []byte("ftest"))
	underBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	buf := newClosableBuffer(underBuf)
	ms := NewServerMessageStream("echo", buf)

	for i := 0; i < b.N; i++ {
		underBuf.Reset()
		b.StartTimer()
		for j := 0; j < 1000; j++ {
			ms.Send(&m)
		}
		b.StopTimer()
		for j := 0; j < 1000; j++ {
			ms.Next(&pong)
		}
	}
}

//Benchmark for testing how long the send operation on a MessageStream takes when used soly in
//memory.
func BenchmarkEchoNextNet(b *testing.B) {
	b.StopTimer()
	var ping Message
	var pong Message

	ping.InitFromBytes('i', []byte("ftest"))
	underBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	buf := newClosableBuffer(underBuf)
	ms := NewServerMessageStream("echo", buf)

	for i := 0; i < b.N; i++ {
		underBuf.Reset()
		for j := 0; j < 1000; j++ {
			ms.Send(&m)
		}
		b.StartTimer()
		for j := 0; j < 1000; j++ {
			ms.Next(&pong)
		}
		b.StopTimer()
	}
}

func initNet() (client *net.TCPConn, server *net.TCPListener) {
	server, serr := autoDial(":2222")
	if serr != nil {
		panic(serr.Error())
	}
	client, cerr := autoDial("localhost:2222")
	if cerr != nil {
		panic(serr.Error())
	}
}

//Test function to make sure that everything is working before benching
func TestEchoNet(t *testing.T) {
	var ping Message
	var pong Message

	serve, err = net.Listen("tcp", ":2222")
	if err != nil {
		t.Fail(err)
	}

	ping.InitFromBytes('i', []byte("ftest"))

	buf := newClosableBuffer(bytes.NewBuffer(make([]byte, 0, 1024)))

	ms := NewServerMessageStream("echo", buf)
	ms.Send(&m)
	t.Logf("%v", buf)

	ms.Next(&pong)

	rest, _ := pong.Force()
	t.Logf("Type:%c, bytes:%s,", pong.MsgType(), rest)
}

//Benchmark for testing how long the send operation on a MessageStream takes when used soly in
//memory.
func BenchmarkEchoNextMem(b *testing.B) {
	b.StopTimer()
	var ping Message
	var pong Message

	ping.InitFromBytes('i', []byte("ftest"))
	underBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	buf := newClosableBuffer(underBuf)
	ms := NewServerMessageStream("echo", buf)

	for i := 0; i < b.N; i++ {
		underBuf.Reset()
		for j := 0; j < 1000; j++ {
			ms.Send(&m)
		}
		b.StartTimer()
		for j := 0; j < 1000; j++ {
			ms.Next(&pong)
		}
		b.StopTimer()
	}
}

//Test function to make sure that everything is working before benching
func TestEcho(t *testing.T) {
	var ping Message
	ping.InitFromBytes('i', []byte("ftest"))
	var pong Message

	buf := newClosableBuffer(bytes.NewBuffer(make([]byte, 0, 1024)))

	ms := NewServerMessageStream("echo", buf)
	ms.Send(&m)
	t.Logf("%v", buf)

	ms.Next(&pong)

	rest, _ := pong.Force()
	t.Logf("Type:%c, bytes:%s,", pong.MsgType(), rest)
}
