package queue

import "time"

// -----------------------------------------------------------------------------

// TimeoutReached returns `true` if the time elapsed since `m.Timeout` is
// greater or equal to `d`. It returns `false` otherwise.
func (m *Message) TimeoutReached(d time.Duration) bool {
	if elapsed := time.Since(m.Timeout); elapsed >= d {
		return true
	}
	return false
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
