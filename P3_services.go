package main

import (
	"fmt"
	"net"
)

func main() {
	//	Service

	networkType := "tcp"
	Services := "FTP"

	p, _ := net.LookupPort(networkType, Services)
	fmt.Printf("port associated with %s is %d\n", Services, p)
}
