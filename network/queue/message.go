package queue

import "time"

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
func (q Queue) Pop() *Message {
	// RLOCK QUEUE
	// defer
	if len(q) != 0 {
		msg := q[0]
		q = q[0:1]
		return msg
	}
	return nil
}

// TODO
func (q Queue) Push(msg *Message) {
	// WLOCK QUEUE
	// defer
	q = append(q, msg)
}

// TODO
func (q Queue) Poll(d time.Duration) {
	// launched with go
	// Specify that it is a long running routine

	// PUT LOCKING OR GOROUTINE IN PLACE !!!!!

	// if empty: wait duration d
	// else
	// pop => compare => resend OR push
}

// Resend

// TODO
func (q Queue) Discard(ID int) {
	// LOCK QUEUE
	// defer

	// Iterate on queue
	// pop if found
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

// TODO first in last out
type Queue []*Message

// NewQueue returns a new `queue.Queue`.
func NewQueue(size int) Queue { return make([]*Message, size) }
