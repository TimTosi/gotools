package queue

// -----------------------------------------------------------------------------

// Broker is the interface that wraps queue broker's methods.
type Broker interface {
	Close()
}

// Worker is the interface that wraps queue worker's methods.
type Worker interface {
	Close()
}
