package queue

import zmq "github.com/pebbe/zmq4"

// -----------------------------------------------------------------------------

// Receive returns a `string` containing a message received on `w.soc` or
// an error.
func (w *ZMQWorker) Receive() (msg string, err error) {
	w.soc.Send("", zmq.SNDMORE)
	w.soc.Send("Hi Boss", 0) // Change msg and send last message ID received

	if _, err := w.soc.Recv(0); err != nil {
		return "", err
	}
	return w.soc.Recv(0)
}

// Close releases resources acquired by `w.soc`.
func (w *ZMQWorker) Close() { _ = w.soc.Close() }

// Identity returns a `string` containing `w.soc`s identity or an `error`.
func (w *ZMQWorker) Identity() (string, error) { return w.soc.GetIdentity() }

// -----------------------------------------------------------------------------

// ZMQWorker is a structure representing a worker process. The network
// communication stack lies on a Go implementation of the ZeroMQ library.
type ZMQWorker struct {
	soc *zmq.Socket
}

// NewZMQWorker returns a new `ZMQWorker`.
//
// NOTE: `id` should be unique.
// NOTE: `addr` must be of the following form
// - `tcp://<hostname>:<port>` for "regular" TCP networking.
// - `inproc://<name>` for in-process networking.
// - `ipc:///<tmp/filename>` for inter-process communication.
func NewZMQWorker(addr, id string) (*ZMQWorker, error) {
	soc, err := zmq.NewSocket(zmq.DEALER)
	if err != nil {
		return nil, err
	}
	if err := soc.SetIdentity(id); err != nil {
		return nil, err
	}
	if err := soc.Connect(addr); err != nil {
		return nil, err
	}
	return &ZMQWorker{soc: soc}, nil
}
