package logger

import (
	"io"
	"sync"
)

const defaultBufSize = 4096

type Writer struct {
	err  error
	buf  []byte
	n    int
	wr   io.Writer
	lock sync.Mutex
}

func NewWriter(w io.Writer) *Writer {
	size := defaultBufSize
	b, ok := w.(*Writer)
	if ok && len(b.buf) >= size {
		return b
	}
	return &Writer{
		buf: make([]byte, size),
		wr:  w,
	}
}

func (b *Writer) Flush() error {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.flush()
}

func (b *Writer) flush() error {
	if b.err != nil {
		return b.err
	}
	if b.n == 0 {
		return nil
	}
	n, err := b.wr.Write(b.buf[0:b.n])
	if n < b.n && err == nil {
		err = io.ErrShortWrite
	}
	if err != nil {
		if n > 0 && n < b.n {
			copy(b.buf[0:b.n-n], b.buf[n:b.n])
		}
		b.n -= n
		b.err = err
		return err
	}
	b.n = 0
	return nil
}
func (b *Writer) Write(p []byte) (nn int, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	return b.write(p)
}

func (b *Writer) write(p []byte) (nn int, err error) {
	for len(p) > len(b.buf)-b.n && b.err == nil {
		var n int
		if b.n == 0 {
			n, b.err = b.wr.Write(p)
		} else {
			n = copy(b.buf[b.n:], p)
			b.n += n
			_ = b.flush()
		}
		nn += n
		p = p[n:]
	}
	if b.err != nil {
		return nn, b.err
	}
	n := copy(b.buf[b.n:], p)
	b.n += n
	nn += n
	return nn, nil
}
