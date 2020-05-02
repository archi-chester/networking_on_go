package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	url, err := url.Parse("http://diptera.myspecies.info:80")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url.String(), nil)

	resp, err := client.Do(req)

	if resp.Status != "200 OK" {
		panic(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
}
