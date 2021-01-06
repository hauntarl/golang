package main

import (
	"fmt"
	"time"
)

const (
	logI = "INFO"
	logW = "WARNING"
	logE = "ERROR"
)

type log struct {
	time     time.Time
	severity string
	message  string
}

func newLog(severity, message string) log {
	return log{
		time:     time.Now(),
		severity: severity,
		message:  message,
	}
}

func (l log) String() string {
	return fmt.Sprintf(
		"%v - [%v]%v",
		l.time.Format("2006-01-02T15:04:05"),
		l.severity,
		l.message,
	)
}

func main() {
	var (
		ch     = make(chan log, 50)
		signal = make(chan struct{})
		// acts as a signal only channel, empty struct = 0 memory allocation
	)

	go logger(ch, signal)

	ch <- newLog(logI, "App is starting")
	time.Sleep(time.Second)

	ch <- newLog(logW, "Something about to go down")
	time.Sleep(time.Second)

	ch <- newLog(logE, "Something went down")
	time.Sleep(time.Second)

	signal <- struct{}{}
	time.Sleep(time.Second)
	// you can try sending data on closed channel, go-runtime will panic
}

func logger(ch chan log, signal chan struct{}) {
	defer fmt.Println("Closing log channel")
	defer fmt.Println("Closing signal channel")
	defer close(ch)
	defer close(signal)
	for {
		select {
		case entry := <-ch:
			fmt.Println(entry)
		case _ = <-signal:
			fmt.Println("Signal to wrap things up...")
			return
		}
		// entire select statement is blocked until a message is received on
		// one of the case
	}
}
