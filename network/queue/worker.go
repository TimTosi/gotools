package queue

import zmq "github.com/pebbe/zmq4"

// -----------------------------------------------------------------------------

// TODO
func (w *ZMQWorker) Receive() (msg string, err error) {
	w.soc.Send("", zmq.SNDMORE)
	w.soc.Send("Hi Boss", 0) // Change msg

	if _, err := w.soc.Recv(0); err != nil {
		return "", nil
	}
	return w.soc.Recv(0)
}

// TODO
func (w *ZMQWorker) Close() {
	w.Close()
}

// TODO
func (w *ZMQWorker) Identity() (string, error) { return w.soc.GetIdentity() }

// -----------------------------------------------------------------------------

// TODO
type ZMQWorker struct {
	Worker

	soc *zmq.Socket
}

// NewZMQWorker returns a new `ZMQWorker`.
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
