package queue

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------------------

// TimeoutReached returns `true` if the time elapsed since `m.Timeout` is
// greater or equal to `d`. It returns `false` otherwise.
func (m *Message) TimeoutReached(d time.Duration) bool {
	if elapsed := time.Since(m.Timeout); elapsed >= d {
		return true
	}
	return false
}

// Copy returns a new copy of `m` with updated `m.Timeout`.
func (m *Message) Copy() *Message { return NewMessage(m.ID, m.Msg) }

// ToString returns the `string` representation of `m`.
func (m *Message) ToString() string {
	return fmt.Sprintf("{ID: %d, Msg: \"%s\"}", m.ID, m.Msg)
}

// -----------------------------------------------------------------------------

// Message is a structure representing messages sent and buffered between
// `queue.ZMQBroker` and `queue.ZMQWorker`.
type Message struct {
	ID      int
	Msg     string
	Timeout time.Time
}

// NewMessage returns a new `queue.Message`.
func NewMessage(ID int, msg string) *Message { return &Message{ID: ID, Msg: msg, Timeout: time.Now()} }
