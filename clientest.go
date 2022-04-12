package main

import (
	"flag"
	"log"

	"github.com/hktalent/NatTrackerServer/lib"
)

func main() {
	addrs := flag.String("addr", "127.0.0.1:51630", "test server addr, eg: 127.0.0.1:51630")
	flag.Parse()

	address := *addrs
	s, err := lib.GetPublicIP(address)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf(s)
	}
}
