package tcpserver

import (
	"fmt"
	"log"
	"net"
)

type TcpServer struct {
	listener net.Listener
}

func (this *TcpServer) Run(ip, port string) {
	plisterer, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))

	if err != nil {
		log.Fatalf("listen ip:%s,port:%s failed,fatal", ip, port)
	}

	this.listener = plisterer

	go func() {

		for {
			conn, errConn := this.listener.Accept()
			if errConn != nil {
				log.Print("listener accept error ,err is ", errConn)
				continue
			}

			tcpconn := &TcpConn{conn: conn}
			go tcpconn.Read()
		}

	}()
}
