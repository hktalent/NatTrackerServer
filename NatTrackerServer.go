package main

import (
	"flag"
	"log"
	"net"
	"runtime"
	"sync"

	ntp "github.com/hktalent/NatTrackerServer/lib"
)

var (
	bufferPool    sync.Pool
	mq            messageQueue
	nbWorkers     int
	address       string
	UDPPacketSize int = 1500
	maxQueueSize  int = 1000000
)

type messageQueue chan ntp.NatTrackerProtocol

// get message
func (mq messageQueue) enqueue(m ntp.NatTrackerProtocol) {
	mq <- m
}

// do message
func (mq messageQueue) dequeue() {
	for m := range mq {
		handleMessage(m)
		defer bufferPool.Put([]byte(m.Msg))
	}
}

// 监听
func listenAndReceive(maxWorkers int) error {
	c, err := net.ListenPacket("udp", address)
	if err != nil {
		return err
	}
	for i := 0; i < maxWorkers; i++ {
		go mq.dequeue()
		go receive(c)
	}
	return nil
}

// receive accepts incoming datagrams on c and calls handleMessage() for each message
func receive(c net.PacketConn) {
	for {
		message := bufferPool.Get().([]byte)
		rlen, remote, err := c.ReadFrom(message[0:])
		if err != nil {
			log.Printf("Error %s", err)
			continue
		}
		mq.enqueue(ntp.NatTrackerProtocol{Msg: string(message[0:rlen]), Remote: remote, Conn: c, MsgLen: rlen})
	}
}

// 处理协议
func handleMessage(m ntp.NatTrackerProtocol) {
	ntp.Dispacker(&m)
}

// https://stackoverflow.com/questions/21968266/handling-read-write-udp-connection-in-go
func main() {
	// init
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.StringVar(&address, "addr", ":51630", "Address of the UDP server to test")
	flag.IntVar(&nbWorkers, "concurrency", runtime.NumCPU(), "Number of workers to run in parallel")
	flag.IntVar(&UDPPacketSize, "UDPPacketSize", 1500, "Number of UDPPacketSize")
	flag.IntVar(&maxQueueSize, "maxQueueSize", 1000000, "Number of maxQueueSize")

	flag.Parse()

	bufferPool = sync.Pool{
		New: func() interface{} { return make([]byte, UDPPacketSize) },
	}
	mq = make(messageQueue, maxQueueSize)
	listenAndReceive(nbWorkers)
	// net.ListenPacket icmp
	// net.ListenUDP udp
	/*
		l := ipv4.NewPacketConn(listener)
		l.SetControlMessage(ipv4.FlagDst, true)

		n, cs, remoteAddr, err := l.ReadFrom(data)
		cs.Src = cs.Dst
		n, err = l.WriteTo([]byte("world"), cs, remoteAddr)
	*/
	// port := *ddrs
	// protocol := "udp4"

	// udpAddr, err := net.ResolveUDPAddr(protocol, port)
	// if err != nil {
	// 	log.Println("Wrong Address")
	// 	return
	// }
	// l, err := net.ListenUDP(protocol, udpAddr)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer l.Close()

	// nC := make(chan *ntp.NatTrackerProtocol, 10000)
	// for {

	// 	select {
	// 	case cc := <-nC:
	// 		{
	// 			go ntp.Dispacker(cc)
	// 		}
	// 	default:
	// 		{
	// 			message := make([]byte, 6666)
	// 			rlen, remote, err := l.ReadFromUDP(message[:])
	// 			if err == nil && 5 <= rlen {

	// 				nC <- &ntp.NatTrackerProtocol{Msg: string(message[0:rlen]), Remote: remote, Conn: l, MsgLen: rlen}
	// 			}

	// 		}
	// 	}
	// }
}
