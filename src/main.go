package main

import (
	"fmt"

	pblogin "QueueService/common/pb/login"
)

func main() {

	req := &pblogin.LoginRequest{}
	req.Username = "123"
	req.Password = "321"
	fmt.Print("Queue Service start,req is ", req)
}
