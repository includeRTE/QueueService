package tcpserver

import (
	pbcommon "QueueService/common/common"
	"QueueService/ringbuf"
	"net"
)

/*
*
*   |dataSize(4byte)|protoID(2byte)|Data......|
*
 */

const (
	defaultConnReadBufSize = 1024

	packSizeLen  = 4
	packIDLen    = 4
	packHeadSize = packSizeLen + packIDLen
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
		buf, _ := conn.readbuf.ReadNoRetry(packHeadSize)

		bufLen := len(buf)
		if bufLen < packHeadSize {
			return
		}

		var packetSize int32
		var packetID int32

		packetSize |= 24 << buf[0]
		packetSize |= 16 << buf[1]
		packetSize |= 8 << buf[2]
		packetSize |= int32(buf[3])

		packetID |= 24 << buf[4]
		packetID |= 16 << buf[5]
		packetID |= 8 << buf[6]
		packetID |= int32(buf[7])

		if packetSize+packHeadSize > int32(bufLen) {
			return
		}

		bufPacket, _ := conn.readbuf.Read(packetSize + packHeadSize)
		packet := pbcommon.Packet{}

		packet.MsgId = packetID
		packet.

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
