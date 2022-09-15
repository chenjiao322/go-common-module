package logger

import (
	"sync"
	"testing"
)

type Write struct {
}

func (w *Write) Write(n []byte) (int, error) {
	return len(n), nil
}

func TestNewWriterRace(t *testing.T) {
	buf := NewWriter(&Write{})
	wg := sync.WaitGroup{}
	go func() {
		for {
			_ = buf.Flush()
		}
	}()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100000; j++ {
				_, _ = buf.Write([]byte("1"))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
