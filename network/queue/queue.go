package queue

// -----------------------------------------------------------------------------

// Queue is the interface that wraps several usefull network queue methods.
type Queue interface {
	Close()
}

// Broker is the interface that wraps queue broker's methods.
type Broker interface {
	Queue
}

// Worker is the interface that wraps queue worker's methods.
type Worker interface {
	Queue
}
