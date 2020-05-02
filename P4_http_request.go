package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	address := "www.google.com:80"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	req := "HEAD / HTTP/1.0\r\n\r\n"

	_, err = conn.Write([]byte(req))
	if err != nil {
		panic(err)
	}

	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp))


}
