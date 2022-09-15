package deque

//
// 在 https://github.com/eapache/queue/blob/master/queue.go 中参考了大量代码
// 修改成 双端都可以pop和append

type Deque struct {
	buf   []byte
	head  int
	tail  int
	count int
}

const InitCap = 1 << 4

func NewDeque() *Deque {
	return &Deque{
		buf:   make([]byte, InitCap),
		head:  0,
		tail:  0,
		count: 0,
	}
}

func (q *Deque) tryResize() {
	bufSize := len(q.buf)
	toSmall := q.count == bufSize
	toBig := bufSize > InitCap && q.count<<2 == bufSize
	if toSmall || toBig {
		newBuf := make([]byte, q.count<<1)
		if q.tail > q.head {
			copy(newBuf, q.buf[q.head:q.tail])
		} else {
			n := copy(newBuf, q.buf[q.head:])
			copy(newBuf[n:], q.buf[:q.tail])
		}
		q.head = 0
		q.tail = q.count
		q.buf = newBuf
	}
}

func (q *Deque) AppendLeft(c byte) {
	q.tryResize()
	if q.head == 0 {
		q.head = len(q.buf)
	}
	q.head = (q.head - 1) & q.Mask()
	q.buf[q.tail] = c
	q.count += 1
}

func (q *Deque) AppendRight(c byte) {
	q.tryResize()
	q.buf[q.tail] = c
	q.tail = (q.tail + 1) & q.Mask()
	q.count += 1
}

func (q *Deque) PopLeft() byte {
	if q.Len() == 0 {
		panic("pop empty queue")
	}
	out := q.buf[q.head]
	q.head = (q.head + 1) & q.Mask()
	q.count -= 1
	q.tryResize()
	return out
}

func (q *Deque) PopRight() byte {
	if q.Len() == 0 {
		panic("pop empty queue")
	}
	q.tail = (q.tail - 1) & q.Mask()
	q.count -= 1
	out := q.buf[q.tail]
	q.tryResize()
	return out
}

func (q *Deque) PeekLeft() byte {
	if q.Len() == 0 {
		panic("peek empty queue")
	}
	return q.buf[q.head]
}

func (q *Deque) PeekRight() byte {
	if q.Len() == 0 {
		panic("peek empty queue")
	}
	if q.tail == 0 {
		return q.buf[len(q.buf)-1]
	}
	return q.buf[q.tail-1]
}

func (q *Deque) Len() int {
	return q.count
}

func (q *Deque) Mask() int {
	return len(q.buf) - 1
}

func (q *Deque) ToList() []byte {
	cur := q.head
	out := make([]byte, q.count)
	idx := 0
	for ; cur != q.tail; cur, idx = (cur+1)&q.Mask(), idx+1 {
		out[idx] = q.buf[cur]
	}
	return out
}

func (q *Deque) Clear() {
	q.buf = make([]byte, InitCap)
	q.head = 0
	q.tail = 0
	q.count = 0
}
