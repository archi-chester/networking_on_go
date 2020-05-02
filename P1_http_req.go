package main

import (
	"fmt"
	"net/http"
)

func main() {

	//	url	:
	//	http 3

	//	user agents

	//	net.http

	url := "http://diptera.myspecies.info:80"

	//	head
	resp, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status, " ", resp.StatusCode)

	for k, v := range resp.Header {
		fmt.Println(k, " : ", v)
	}
	//fmt.Println(resp.)

}
