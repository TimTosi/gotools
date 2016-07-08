package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/timtosi/gotools/network/queue"
)

// -----------------------------------------------------------------------------

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments.\n")
	}
	wq, err := queue.NewZMQWorker(fmt.Sprintf("tcp://127.0.0.1:%s", os.Args[1]), "1")
	if err != nil {
		log.Fatal(err)
	}
	wq2, err := queue.NewZMQWorker(fmt.Sprintf("tcp://127.0.0.1:%s", os.Args[1]), "2")
	if err != nil {
		log.Fatal(err)
	}
	wq3, err := queue.NewZMQWorker(fmt.Sprintf("tcp://127.0.0.1:%s", os.Args[1]), "3")
	if err != nil {
		log.Fatal(err)
	}

	do := func(w *queue.ZMQWorker) {
		for {
			fmt.Printf("OK\n")
			workerID, _ := w.Identity()
			if msg, err := w.Receive(); err == nil {
				fmt.Printf("Woker %s - Message Received: %v\n", workerID, msg)
			} else {
				fmt.Printf("Worker %s - BUG HERE %v\n", workerID, err)
			}
		}
	}

	go do(wq)
	go do(wq2)
	go do(wq3)

	for {
		time.Sleep(10 * time.Millisecond)
	}
}
