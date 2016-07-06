package queue

import (
	"testing"
	"time"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

func TestMessage_TimeoutReached_true(t *testing.T) {
	m := NewMessage(1, "Ok")
	time.Sleep(1 * time.Second)
	ensure.True(t, m.TimeoutReached(1*time.Second))
}

func TestMessage_TimeoutReached_false(t *testing.T) {
	m := NewMessage(1, "Ok")
	ensure.False(t, m.TimeoutReached(1*time.Second))
}
