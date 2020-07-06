package circular

import (
	"errors"
	"io"
)

type Buffer struct {
	data []byte
	currentCapacity int
	oldestValueIndex int
	newestValueIndex int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size)}
}

func (b *Buffer) ReadByte() (res byte, err error) {
	if b.currentCapacity == 0 {
		err = io.EOF
		return
	}
	res = b.data[b.oldestValueIndex]
	b.oldestValueIndex = (b.oldestValueIndex+1) % len(b.data)
	b.currentCapacity --
	return
}

func (b *Buffer) WriteByte(c byte) error {
	if b.currentCapacity == len(b.data) {
		return errors.New("buffer full")
	}
	b.data[b.newestValueIndex] = c
	b.newestValueIndex = (b.newestValueIndex+1) % len(b.data)
	b.currentCapacity ++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		b.data[b.oldestValueIndex] = c
		b.oldestValueIndex = (b.oldestValueIndex + 1) % len(b.data)
		return
	}
}

func (b *Buffer) Reset() {
	b.data = make([]byte, len(b.data))
	b.oldestValueIndex = 0
	b.newestValueIndex = 0
	b.currentCapacity = 0
}