package types

import (
	"encoding/json"
	"strconv"
)

type RingBuffer[T any] struct {
	buf []T

	base int

	head int
	tail int
}

func NewRingBuffer[T any](base int, capacity ...int) *RingBuffer[T] {
	c := 16
	if len(capacity) > 0 && capacity[0] > 0 {
		c = capacity[0]
	}

	return &RingBuffer[T]{
		base: base,
		buf:  make([]T, c),
		head: 0,
		tail: 0,
	}
}

func (b *RingBuffer[T]) Capacity() int {
	return cap(b.buf)
}

func (b *RingBuffer[T]) Length() int {
	switch {
	case b.head >= b.tail:
		return b.head - b.tail
	case b.head < b.tail:
		return len(b.buf) + b.head - b.tail
	default:
		return 0
	}
}

func (b *RingBuffer[T]) Push(value T) {
	switch {
	case b.head >= b.tail:
		//  0   1   2   3
		//  t       h      : 2 elements
		size := b.head - b.tail
		if size == len(b.buf)-1 { // time to grow
			newBuf := make([]T, 2*len(b.buf))
			copy(newBuf, b.buf)
			b.buf = newBuf
		}
		b.buf[b.head] = value
		b.head++
		if b.head == len(b.buf) {
			b.head = 0
		}
	case b.head < b.tail:
		//  0   1   2   3
		//  h           t  : 1 element
		size := len(b.buf) + b.head - b.tail
		if size == len(b.buf)-1 { // time to grow
			newBuf := make([]T, 2*len(b.buf))
			copy(newBuf, b.buf[b.tail:])
			copy(newBuf[len(b.buf)-b.tail:], b.buf[:b.head])
			b.buf = newBuf
			b.tail = 0
			b.head = size
		}
		b.buf[b.head] = value
		b.head++
	}
}

// Pop removes an element from the tail of the buffer and returns its value
func (b *RingBuffer[T]) Pop() (T, bool) {
	if b.head == b.tail {
		var res T // nil value
		return res, false
	}

	v := b.buf[b.tail]
	b.tail++
	if b.tail == len(b.buf) {
		b.tail = 0
	}

	b.base++

	return v, true
}

// Head returns the value at the head of the buffer
func (b *RingBuffer[T]) Head() (T, bool) {
	if b.head == b.tail {
		var res T // nil value
		return res, false
	}

	head := b.head - 1
	if head < 0 {
		head += len(b.buf)
	}

	return b.buf[head], true
}

// Pick removes an element from the head of the buffer and returns its value
func (b *RingBuffer[T]) Pick() (T, bool) {
	if b.head == b.tail {
		var res T // nil value
		return res, false
	}

	b.head--
	if b.head < 0 {
		b.head += len(b.buf)
	}

	return b.buf[b.head], true
}

func (b *RingBuffer[T]) Forget(count int) {
	if count >= b.Length() {
		b.head = 0
		b.tail = 0
		return
	}

	b.head -= count
	if b.head < 0 {
		b.head += len(b.buf)
	}
}

func (b *RingBuffer[T]) At(idx int) (T, bool) {
	if idx < b.base {
		var res T // nil value
		return res, false
	}

	offset := idx - b.base

	if offset > b.Length()-1 {
		var res T // nil value
		return res, false
	}

	index := b.tail + offset
	if index >= len(b.buf) {
		index -= len(b.buf)
	}

	return b.buf[index], true
}

func (b *RingBuffer[T]) MarshalJSON() ([]byte, error) {
	var buf []T

	switch {
	case b.head >= b.tail:
		//  0   1   2   3
		//  t       h      : 2 elements
		buf = b.buf[b.tail:b.head]

	case b.head < b.tail:
		//  0   1   2   3
		//  h           t  : 1 element
		buf = b.buf[b.tail:len(b.buf)]
		buf = append(buf, b.buf[:b.head]...)
	}

	bytes, err := json.Marshal(buf)
	if err != nil {
		return nil, err
	}

	return append(append([]byte(`{"base":`+strconv.Itoa(b.base)+`,"buf":`), bytes...), byte('}')), nil
}

func (b *RingBuffer[T]) UnmarshalJSON(data []byte) error {
	raw := &struct {
		Base int
		Buf  json.RawMessage
	}{}
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}

	buf := make([]T, 0)
	if err := json.Unmarshal(raw.Buf, &buf); err != nil {
		return err
	}
	var _nil T
	buf = append(buf, _nil)

	b.base = raw.Base
	b.buf = buf
	b.head = len(buf) - 1

	return nil
}
