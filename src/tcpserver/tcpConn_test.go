package tcpserver

import (
	"log"
	"testing"
)

func TestConn(t *testing.T) {

	conn := &TcpConn{}

	log.Print("conn is ", conn)
}
