package closeable

import "bytes"

func NewStringBuffer(buff string) *Buffer {
	return &Buffer{bytes.NewBufferString(buff)}
}

func NewBytesBuffer(buff []byte) *Buffer {
	return &Buffer{bytes.NewBuffer(buff)}
}

func (buff *Buffer) Close() error {
	return nil
}

type Buffer struct {
	*bytes.Buffer
}
