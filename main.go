package main

import (
	"fmt"
	"net/http"
	"os"

	natsproxy "github.com/aliheydarabadii/nats-proxy"
	"github.com/nats-io/nats"
)

func main() {
	host := os.Getenv("host")
	proxyConn, err := nats.Connect(host)
	fmt.Println(err)
	proxy, err := natsproxy.NewNatsProxy(proxyConn)
	fmt.Println(err)
	defer proxyConn.Close()
	err = http.ListenAndServe("0.0.0.0:8083", proxy)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func consume(conn *nats.Conn) {
	conn.Subscribe("POST:.foo", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})
}
