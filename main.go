package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

// https://stackoverflow.com/questions/21968266/handling-read-write-udp-connection-in-go
func main() {
	addrs := flag.String("addr", ":51630", "ListenUDP addr, eg: :51630")
	flag.Parse()
	// net.ListenPacket icmp
	// net.ListenUDP udp
	/*
		l := ipv4.NewPacketConn(listener)
		l.SetControlMessage(ipv4.FlagDst, true)

		n, cs, remoteAddr, err := l.ReadFrom(data)
		cs.Src = cs.Dst
		n, err = l.WriteTo([]byte("world"), cs, remoteAddr)
	*/
	port := *addrs
	protocol := "udp4"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		log.Println("Wrong Address")
		return
	}
	l, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()

	nC := make(chan *net.UDPAddr, 1000)
	for {

		select {
		case cc := <-nC:
			{
				go func(cc1 *net.UDPAddr) {
					s := fmt.Sprintf(cc1.String())
					log.Println(s)
					l.WriteToUDP([]byte(s), cc1)
				}(cc)
			}
		default:
			{
				message := make([]byte, 30)
				rlen, remote, err := l.ReadFromUDP(message[:])
				if err == nil && 3 <= rlen {
					nC <- remote
				}

			}
		}
	}
}
