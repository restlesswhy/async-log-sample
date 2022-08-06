package alog

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"
)

var ErrClosed = errors.New("alog closed")

type Alog interface {
	Info(v ...any)
	io.Closer
}

type alog struct {
	log    *log.Logger
	close  chan struct{}
	infoCh chan []any

	wg sync.WaitGroup
}

func New(workers int) Alog {
	alog := &alog{
		log:    log.New(os.Stdout, "", log.LstdFlags),
		close:  make(chan struct{}),
		infoCh: make(chan []any),
	}

	alog.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go alog.run()
	}

	return alog
}

func (a *alog) run() {
	defer a.wg.Done()

main:
	for {
		select {
		case <-a.close:
			break main

		case info := <-a.infoCh:
			go func (info ...any)  {
				a.log.Println(info...)
			}(info...)
		}
	}
}

func (a *alog) Info(v ...any) {
	a.infoCh <- v
}

func (a *alog) Close() error {
	select {
	case <-a.close:
		return ErrClosed

	default:
		close(a.close)
		a.wg.Wait()
		return nil
	}
}