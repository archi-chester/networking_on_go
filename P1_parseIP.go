package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//	Net.ParseIP()

	if len(os.Args) != 2 {
		fmt.Println("[-] Invalid IP")
		os.Exit(1)
	}

	IP := os.Args[1]

	address := net.ParseIP(IP)

	fmt.Println(address)
	fmt.Printf("%T\n", address)

}
