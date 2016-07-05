package queue

import (
	"sync"
	"time"
)

// -----------------------------------------------------------------------------

// TODO
func (m *Message) TimeoutReached(d time.Duration) bool {
	if elapsed := time.Since(m.Timeout); elapsed >= d {
		return true
	}
	return false
}

// -----------------------------------------------------------------------------

// TODO
func (q Queue) Shift() *Message {
	q.Lock()
	defer q.Unlock()

	if len(q.msgs) != 0 {
		msg := q.msgs[0]
		q.msgs = q.msgs[1:]
		return msg
	}
	return nil
}

// TODO
func (q Queue) Push(msg *Message) {
	q.Lock()
	defer q.Unlock()

	q.msgs = append(q.msgs, msg)
}

// TODO
// launched with go
// Specify that it is a long running routine
func (q Queue) Poll(IDChan <-chan int, msgChan chan<- *Message, d time.Duration) {
	for {
		select {
		case ID := <-IDChan:
			q.Discard(ID)
		case <-time.After(d):
			msgChan <- q.Purge(d)
		}
	}

	// if empty: wait duration d
	// else
	// pop => compare => resend OR push
}

// TODO RESEND !!!!

// TODO
// NOTE: No leak
func (q Queue) Discard(ID int) bool {
	q.Lock()
	defer q.Unlock()

	for i := 0; i < len(q.msgs); i++ {
		if q.msgs[i].ID == ID {
			copy(q.msgs[i:], q.msgs[i+1:])
			q.msgs[len(q.msgs)-1] = nil
			q.msgs = q.msgs[:len(q.msgs)-1]
			return true
		}
	}
	return false
}

// TODO
// NOTE: No leak
// TODO BETTER LOCKING HERE
func (q Queue) Purge(d time.Duration) *Message {
	q.Lock()
	defer q.Unlock()

	for i := 0; i < len(q.msgs); i++ {
		if q.msgs[i].TimeoutReached(d) == true {
			msg := q.msgs[i]
			copy(q.msgs[i:], q.msgs[i+1:])
			q.msgs[len(q.msgs)-1] = nil
			q.msgs = q.msgs[:len(q.msgs)-1]
			return msg
		}
	}
	return nil
}

// -----------------------------------------------------------------------------

// TODO
type Message struct {
	ID      int
	Msg     string
	Timeout time.Time
}

// NewMessage returns a new `queue.Message`.
func NewMessage(ID int, msg string) *Message { return &Message{ID: ID, Msg: msg, Timeout: time.Now()} }

// TODO + explanation first in last out container
type Queue struct {
	*sync.RWMutex

	msgs []*Message
}

// NewQueue returns a new `queue.Queue`.
func NewQueue(size int) Queue { return Queue{RWMutex: &sync.RWMutex{}, msgs: make([]*Message, size)} }
