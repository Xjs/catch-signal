package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	defer func() {
		log.Println("really bye")
	}()
	log.Println("Listening...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c
	log.Println("Got signal:", s)
	time.Sleep(5 * time.Second)
	log.Println("byebye")
}
