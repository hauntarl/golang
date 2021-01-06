package main

import (
	"fmt"
	"time"
)

// these represent severity levels of a log message
const (
	logI = "INFO"
	logW = "WARNING"
	logE = "ERROR"
)

// a structure which defines the components a log has
type log struct {
	time     time.Time
	severity string
	message  string
}

// new is constructor of log structure
func newLog(severity, message string) log {
	return log{
		time:     time.Now(),
		severity: severity,
		message:  message,
	}
}

// defining this function, log implicity implements the built-in Stringer
// interface, which is used by functions like fmt.Println() to get the string
// representation of any object
func (l log) String() string {
	return fmt.Sprintf(
		"%v - [%v]%v",
		l.time.Format("2006-01-02T15:04:05"),
		l.severity,
		l.message,
	)
}

// logger func listens to the log channel, and writes them onto the console
func logger(ch chan log) {
	for entry := range ch {
		fmt.Println(entry)
	}
}

func main() {
	var ch = make(chan log, 50) // buffered channel for logging

	go logger(ch) // creates receiver for log channel
	defer func() {
		fmt.Println("Closing log channel")
		close(ch) // one way of gracefully close the channel
	}()

	ch <- newLog(logI, "App is starting")
	time.Sleep(time.Second)

	ch <- newLog(logW, "Something about to go down")
	time.Sleep(time.Second)

	ch <- newLog(logE, "Something went down")
	time.Sleep(time.Second)
}
