package deque

import (
	"math/rand"
	"testing"
)

func assert(ok bool) {
	if !ok {
		panic("assert fail")
	}
}

func TestNewDeque1(t *testing.T) {
	q := NewDeque()
	q.AppendRight(1)
	q.AppendRight(2)
	assert(q.ToList()[0] == 1)
	assert(q.ToList()[1] == 2)
	assert(q.PopLeft() == 1)
	assert(q.PopLeft() == 2)
}

func TestNewDeque2(t *testing.T) {
	q := NewDeque()
	for i := 0; i < 1000; i++ {
		q.AppendRight(1)
		q.AppendLeft(1)
	}
	assert(q.Len() == 2000)
	q.Clear()
}

func TestPanic(t *testing.T) {
	q := NewDeque()
	func() {
		defer func() {
			recover()
		}()
		q.PeekRight()
	}()
	func() {
		defer func() {
			recover()
		}()
		q.PeekLeft()
	}()
	func() {
		defer func() {
			recover()
		}()
		q.PopLeft()
	}()
	func() {
		defer func() {
			recover()
		}()
		q.PopRight()
	}()
}

func TestFuzz(t *testing.T) {
	q := NewDeque()
	for i := 0; i < 100000; i++ {
		n := rand.Int31n(7)
		switch n {
		case 1:
			q.AppendLeft(1)
		case 2:
			q.AppendRight(1)
		case 3:
			if q.Len() > 0 {
				q.PopLeft()
			}
		case 4:
			if q.Len() > 0 {
				q.PopRight()
			}
		case 5:
			if q.Len() > 0 {
				q.PeekLeft()
			}
		case 6:
			if q.Len() > 0 {
				q.PeekRight()
			}
		case 7:
			q.Len()
		default:
		}
	}
}
