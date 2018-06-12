package main

import (
	"fmt"
	"time"
	"sync"
	"os"
	"os/signal"
	"syscall"
)

func consumer(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func main() {
	stop := make(chan bool)

	var wg sync.WaitGroup

	// Spawn example consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer(stop)
		}(stop)
	}

	waitForSignal()

	close(stop)

	fmt.Println("stopping all jobs!")

	wg.Wait()

	fmt.Println("main process exit!")
}
func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println("wait for signals!")
	<-sigs
}