package queue

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

// -----------------------------------------------------------------------------

// Send emits the message `msg` to the first `queue.ZMQWorker` available.
func (b *ZMQBroker) Send(msg string) error {
	identity, _ := b.soc.Recv(0)
	b.soc.Send(identity, zmq.SNDMORE)
	s1, err1 := b.soc.Recv(0)
	fmt.Printf("PRINT 1: %s -- %v\n", s1, err1) // ERASE
	s2, err2 := b.soc.Recv(0)
	fmt.Printf("PRINT 2: %s -- %v\n", s2, err2) // ERASE
	// b.msgIDChan <- ID received
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
	q         *Queue
	soc       *zmq.Socket
	msgIDChan chan int
}

// NewZMQBroker returns a new `ZMQBroker`.
//
// NOTE: `addr` must be of the following form
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
	return &ZMQBroker{msgIDChan: make(chan int), q: NewQueue(), soc: soc}, nil
}

// Run launches the broker and coordinates the message queue `b.q`.
//
// NOTE: This function is an infinite loop.
func (b *ZMQBroker) Run(d time.Duration, workChan chan *Message) {
	emitAgainChan := make(chan *Message, 0)
	go b.q.Poll(b.msgIDChan, emitAgainChan, d)

	for {
		select {
		case msg := <-emitAgainChan:
			b.Send(msg.ToString()) // HANDLE ERROR
			b.q.Push(msg.Copy())
		default:
			select {
			case msg := <-workChan:
				b.Send(msg.ToString()) // HANDLE ERROR
				b.q.Push(msg)
			default:
			}
		}
	}
}
