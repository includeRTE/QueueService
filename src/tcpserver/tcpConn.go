package tcpserver

import (
	"QueueService/src/ringbuf"
	"net"
)

/*
*
*   |dataSize(4byte)|protoID(2byte)|Data......|
*
 */

const (
	defaultConnReadBufSize = 1024
	packHeadSize           = 6
)

type TcpConn struct {
	conn     net.Conn
	writebuf *ringbuf.RingBuf
	readbuf  *ringbuf.RingBuf
}

func NewTcpConn(tcpconn net.Conn) *TcpConn {

	return &TcpConn{
		conn:     tcpconn,
		writebuf: ringbuf.NewRingBuf(),
		readbuf:  ringbuf.NewRingBuf(),
	}
}

func (conn *TcpConn) ParsePacket() {

	dataSize := conn.GetReadBufDataSize()
	if dataSize >= packHeadSize {

	}
}

func (conn *TcpConn) Read() {

	for {
		buf := make([]byte, 0, defaultConnReadBufSize)
		readSize, err := conn.conn.Read(buf)
		if err != nil || readSize <= 0 {
			break
		}

		conn.readbuf.Write(buf)
		conn.ParsePacket()
	}
}

func (conn *TcpConn) Write(data []byte) error {
	conn.writebuf.Write(data)
	return nil
}

func (conn *TcpConn) GetReadBufDataSize() int {
	return conn.readbuf.DataSize()
}
