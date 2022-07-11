package ringbuf

const (
	defaultBufSize = 1024
)

type RingBuf struct {
	readIndex  int
	writeIndex int
	cap        int

	buf []byte
}

func NewRingBuf() *RingBuf {
	return &RingBuf{
		readIndex:  0,
		writeIndex: 0,
		cap:        defaultBufSize,
		buf:        make([]byte, defaultBufSize, defaultBufSize),
	}
}

func (buf *RingBuf) Write(data []byte) {

	nWriteSize := len(data)
	if nWriteSize <= 0 {
		return
	}

	reservedSize := (buf.cap - buf.writeIndex) + buf.readIndex
	if reservedSize >= nWriteSize {
		if (buf.cap - buf.writeIndex) < nWriteSize {
			copy(buf.buf[0:], buf.buf[buf.readIndex:buf.writeIndex])
			buf.writeIndex = buf.writeIndex - buf.readIndex
			buf.readIndex = 0
		}

		copy(buf.buf[buf.writeIndex:], data[0:])
		buf.writeIndex += nWriteSize

	} else {
		tempBuf := make([]byte, 2*buf.cap, 2*buf.cap)
		buf.cap = 2 * buf.cap
		copy(tempBuf[0:], tempBuf[buf.readIndex:buf.writeIndex])
		dataSize := buf.DataSize()
		buf.buf = tempBuf
		buf.readIndex = 0
		buf.writeIndex = dataSize
		copy(buf.buf[buf.writeIndex:], data[0:])
		buf.writeIndex += nWriteSize
	}
}

func (buf *RingBuf) Read(nSize int) ([]byte, error) {

	dataSize := buf.DataSize()

	readSize := 0
	if dataSize >= nSize {
		readSize = nSize
	} else {
		readSize = dataSize
	}

	temBuf := make([]byte, readSize, readSize)
	copy(temBuf[0:readSize], buf.buf[buf.readIndex:buf.readIndex+readSize])
	buf.readIndex += readSize
	return temBuf, nil
}

func (buf *RingBuf) ReadAll() ([]byte, error) {

	dataSize := buf.DataSize()
	return buf.Read(dataSize)
}

func (buf *RingBuf) DataSize() int {
	return buf.writeIndex - buf.readIndex
}
