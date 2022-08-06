package main

import (
	"alog/alog"
	"flag"
	"log"
	"strconv"
	"time"
)

const (
	DEFAULT_THREADS = 1
	DEFAULT_MSGS = 10
)
func main() {
	var threads, msgs int = DEFAULT_THREADS, DEFAULT_MSGS
	
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		arg1, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("First parameter must be int!")
		}

		arg2, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Second parameter must be int!")
		}

		threads = arg1
		msgs = arg2
	}

	alog := alog.New(threads)
	defer alog.Close()

	for i := 0; i < threads; i++ {
		for j := 0; j < msgs; j++ {
			alog.Info(i, j)
		}
	}

	time.Sleep(1*time.Second)
}