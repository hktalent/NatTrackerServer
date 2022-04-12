package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	addrs := flag.String("addr", "127.0.0.1:51630", "test server addr, eg: 127.0.0.1:51630")
	flag.Parse()
	address := *addrs
	conn, err := net.Dial("udp4", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("nat"))
	if nil != err {
		log.Println("WriteTo", err)
	} else {
		p := make([]byte, 2048)

		n, err := conn.Read(p)
		if err == nil {
			fmt.Printf("%s %v\n", string(p[0:n]), n)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
	}
}
