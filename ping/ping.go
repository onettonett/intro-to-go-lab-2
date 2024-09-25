package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func foo(channel chan string) {
	// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
	for {
		channel <- "ping"
		fmt.Println("Foo has sent ping.")
		time.Sleep(time.Millisecond * 100)
		receive := <-channel
		fmt.Println("Foo has received: " + receive)
	}

}

func bar(channel chan string) {
	// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
	for {
		receive := <-channel
		fmt.Println("Bar has received: " + receive)
		time.Sleep(time.Millisecond * 100)
		channel <- "pong"
		fmt.Println("Bar has sent: " + receive)
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	channel := make(chan string)
	for {
		go foo(channel) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
		go bar(channel)
	}

	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
