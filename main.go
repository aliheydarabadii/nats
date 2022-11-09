package main

import (
	"fmt"
	"net/http"

	natsproxy "github.com/aliheydarabadii/nats-proxy"
	"github.com/nats-io/nats"
)

func main() {
	proxyConn, err := nats.Connect(nats.DefaultURL)
	fmt.Println(err)
	proxy, err := natsproxy.NewNatsProxy(proxyConn)
	fmt.Println(err)
	defer proxyConn.Close()
	err = http.ListenAndServe("localhost:8083", proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
}
