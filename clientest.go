package main

import (
	"log"

	"github.com/hktalent/NatTrackerServer/lib"
)

func main() {
	// addrs := flag.String("addr", "127.0.0.1:51630", "test server addr, eg: 127.0.0.1:51630")
	// flag.Parse()

	// address := *addrs
	// s, err := lib.GetPublicIP(address)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Printf(s)
	// }

	// log.Println(lib.GetOutboundIP().String())
	// a := "//51pwn/P2P&E2E/" + lib.GenerateUUID() + "/0/" + lib.GetMacIPAddrLists2Line()
	// log.Printf(lib.Nat)
	// log.Printf(a)

	x1 := lib.NewNatTrackerProtocol()
	a, err := x1.GetPublicIP()
	if nil == err {
		for _, x := range a {
			log.Printf(x)
		}
	}

	dd := "kkk;"
	log.Printf(dd[:len(dd)-1])
}
