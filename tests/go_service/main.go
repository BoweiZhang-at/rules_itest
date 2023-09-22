package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"
)

var fibSink int

func main() {
	sleepTime := flag.Duration("sleep-time", 0, "How long to sleep before binding the port")
	busyWaitTime := flag.Duration("busy-time", 0, "How long to busy-wait before binding the port")
	port := flag.Int("port", 0, "Port to bind")

	flag.Parse()

	log.Println("started")
	time.Sleep(*sleepTime)
	log.Println("done sleeping")

	finishBusyWait := time.Now().Add(*busyWaitTime)
	for time.Now().Before(finishBusyWait) {
		fibSink = fib(10)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.ListenAndServe("127.0.0.1:"+strconv.Itoa(*port), nil)
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
