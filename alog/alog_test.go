package alog_test

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"alog/alog"
)

func TestAlog(t *testing.T) {
	threads := 3
	msgs := 50

	// Sync
	wg := sync.WaitGroup{}
	log.SetFlags(log.Lmicroseconds)

	startSync := time.Now()
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(d int) {
			defer wg.Done()
			for i := 0; i < msgs; i++ {
				log.Println(d, i)
			}
		}(i)
	}
	stopSync := time.Since(startSync)

	wg.Wait()

	time.Sleep(1 * time.Second)

	// Async
	alog := alog.New(threads)
	defer alog.Close()

	startAsync := time.Now()
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(d int) {
			defer wg.Done()
			for i := 0; i < msgs; i++ {
				alog.Info(d, i)
			}
		}(i)
	}
	stopAsync := time.Since(startAsync)

	wg.Wait()

	time.Sleep(time.Second * 1)
	fmt.Printf("\n----\nSync: %v\nAsync: %v\n----\n\n", stopSync, stopAsync)
}

func BenchmarkAlog(b *testing.B) {
	alog := alog.New(3)
	defer alog.Close()

	for i := 0; i < b.N; i++ {
        alog.Info(i)
    }
}

func BenchmarkLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
        log.Println(i)
    }
}