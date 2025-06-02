package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type syncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	var s syncedBuffer
	var wg sync.WaitGroup

	wg.Add(2)

	// first goRoutine: holds lock longer
	go func() {
		defer wg.Done()
		s.lock.Lock()
		fmt.Println("[goRoutien 1] acquired lock  , writing")
		s.buffer.WriteString("Hello from goRoutine 1!\n")
		time.Sleep(5 * time.Second) // Simulate long operation
		s.lock.Unlock()
		fmt.Println("[goRoutien 1] released lock")
	}()

	// delay second goRoutine slightly so it hits the lock while the first is holding the lock
	go func() {
		defer wg.Done()
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("[goRoutien 2] trying to acquire lock...")
		s.lock.Lock()
		fmt.Println("[goRoutine 2] acquired lock , writing...")
		s.buffer.WriteString("hello from go routine 2 \n")
		time.Sleep(2 * time.Second)
		s.lock.Unlock()
		fmt.Println("[go routine 2 ] released lock")
	}()

	wg.Wait()
	time.Sleep(3 * time.Second)
	fmt.Println("final buffer content: \n" + s.buffer.String())
}
