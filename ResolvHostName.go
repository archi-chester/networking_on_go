package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	Host := "www.google.com"
	ServerAddress, err := net.ResolveIPAddr("ip",Host)
	if err != nil {
		fmt.Println("[-] Couldnt Resolve")
		os.Exit(1)
	} else {
		fmt.Println(ServerAddress.String())
	}
}
