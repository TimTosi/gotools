package queue

import zmq "github.com/pebbe/zmq4"

// -----------------------------------------------------------------------------

// Send emits the message `msg` to the first `queue.ZMQWorker` available.
func (b *ZMQBroker) Send(msg string) error {
	identity, _ := b.soc.Recv(0)
	b.soc.Send(identity, zmq.SNDMORE)
	b.soc.Recv(0)
	b.soc.Recv(0)
	b.soc.Send("", zmq.SNDMORE)
	_, err := b.soc.Send(msg, 0)
	return err
}

// Close releases resources acquired by `b.soc`.
func (b *ZMQBroker) Close() { _ = b.soc.Close() }

// -----------------------------------------------------------------------------

// ZMQBroker is a structure representing a message broker. The network
// communication stack lies on a Go implementation of the ZeroMQ library.
type ZMQBroker struct {
	q   *Queue
	soc *zmq.Socket
}

// NewZMQBroker returns a new `ZMQBroker`.
//
// NOTE: `addr` must be of the form
// - `tcp://<hostname>:<port>` for "regular" TCP networking.
// - `inproc://<name>` for in-process networking.
// - `ipc:///<tmp/filename>` for inter-process communication.
func NewZMQBroker(addr string) (*ZMQBroker, error) {
	soc, err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		return nil, err
	}
	if err := soc.Bind(addr); err != nil {
		return nil, err
	}
	return &ZMQBroker{q: NewQueue(), soc: soc}, nil
}
