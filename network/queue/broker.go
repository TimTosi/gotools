package queue

import zmq "github.com/pebbe/zmq4"

// -----------------------------------------------------------------------------

// TODO
func (b *ZMQBroker) Send(msg string) error {
	identity, _ := b.soc.Recv(0)
	b.soc.Send(identity, zmq.SNDMORE)
	b.soc.Recv(0) //  Envelope delimiter
	b.soc.Recv(0) //  Response from worker
	b.soc.Send("", zmq.SNDMORE)
	_, err := b.soc.Send(msg, 0)
	return err
}

// TODO
func (b *ZMQBroker) Close() {
	b.Close()
	// Runtine.Finalizer ?
}

// -----------------------------------------------------------------------------

// TODO
type ZMQBroker struct {
	Broker

	soc *zmq.Socket
}

// NewZMQBroker returns a new `ZMQBroker`.
//
// NOTE: Address formatting
// - tcp://hostname:port sockets let us do "regular" TCP networking
// - inproc://name sockets let us do in-process networking with the same
// code we'd use for TCP networking.
// - ipc:///tmp/filename sockets use UNIX domain sockets for inter-process
// communication.
func NewZMQBroker(addr string) (*ZMQBroker, error) {
	soc, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return nil, err
	}
	if err := soc.Bind(addr); err != nil {
		return nil, err
	}
	return &ZMQBroker{soc: soc}, nil
}
