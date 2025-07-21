package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "1.1.1.1")

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	// conn.Write()
}

type DNSRequest struct{}

type DNSResponse struct{}

type Resolver interface {
	GenerateTransactionID() uint16
}
